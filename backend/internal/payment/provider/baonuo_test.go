//go:build unit

package provider

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/payment"
	"github.com/stretchr/testify/require"
)

func TestNewBaonuoValidatesConfig(t *testing.T) {
	t.Parallel()

	_, err := NewBaonuo("1", map[string]string{
		"merchantId":  "10086",
		"apiKey":      "secret",
		"channelType": "123",
		"apiBase":     "http://baonuo.example.com",
	})
	require.ErrorContains(t, err, "HTTPS URL")

	prov, err := NewBaonuo("1", map[string]string{
		"merchantId":  "10086",
		"apiKey":      "secret",
		"channelType": "123",
	})
	require.NoError(t, err)
	require.Equal(t, payment.TypeBaonuo, prov.ProviderKey())
	require.Equal(t, []payment.PaymentType{payment.TypeBaonuo}, prov.SupportedTypes())
	require.Equal(t, baonuoDefaultAPIBase, prov.config["apiBase"])
	require.Equal(t, payment.DefaultPaymentCurrency, prov.config["currency"])
	require.Equal(t, "10086", prov.MerchantIdentityMetadata()["merchant_id"])
}

func TestBaonuoSignMatchesDocumentRule(t *testing.T) {
	t.Parallel()

	params := map[string]string{
		"merchantId":  "10086",
		"orderId":     "sub220260526abc",
		"orderAmount": "12.30",
		"channelType": "88",
		"notifyUrl":   "https://merchant.example.com/notify",
		"returnUrl":   "",
		"isForm":      "2",
		"payer_id":    "0",
		"sign":        "ignored",
	}
	wantRaw := "channelType=88&isForm=2&merchantId=10086&notifyUrl=https://merchant.example.com/notify&orderAmount=12.30&orderId=sub220260526abc&key=secret"
	want := fmt.Sprintf("%x", md5.Sum([]byte(wantRaw)))
	require.Equal(t, want, baonuoSign(params, "secret"))
}

func TestBaonuoCreatePaymentPostsSignedForm(t *testing.T) {
	t.Parallel()

	var gotForm url.Values
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/newOrder", r.URL.Path)
		require.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		gotForm, err = url.ParseQuery(string(body))
		require.NoError(t, err)
		require.True(t, baonuoVerifySign(flattenValues(gotForm), "secret", gotForm.Get("sign")))
		_, _ = w.Write([]byte(`{"code":200,"msg":"下单成功!","data":{"payUrl":"https://baonuo.example/pay/123"}}`))
	}))
	defer server.Close()

	prov := mustTestBaonuoProvider(t, server)
	resp, err := prov.CreatePayment(context.Background(), payment.CreatePaymentRequest{
		OrderID:   "sub2_order_123",
		Amount:    "12.34",
		Subject:   "Sub2API 12.34 CNY",
		NotifyURL: "https://merchant.example.com/api/v1/payment/webhook/baonuo",
		ReturnURL: "https://merchant.example.com/payment/result",
		ClientIP:  "203.0.113.9",
	})
	require.NoError(t, err)
	require.Equal(t, "sub2order123", resp.TradeNo)
	require.Equal(t, "https://baonuo.example/pay/123", resp.PayURL)
	require.Equal(t, "CNY", resp.Currency)
	require.Equal(t, "10086", gotForm.Get("merchantId"))
	require.Equal(t, "sub2order123", gotForm.Get("orderId"))
	require.Empty(t, gotForm.Get("payer_id"))
	require.Equal(t, "12.34", gotForm.Get("orderAmount"))
	require.Equal(t, "88", gotForm.Get("channelType"))
	require.Equal(t, "2", gotForm.Get("isForm"))
}

func TestBaonuoMerchantOrderIDMatchesDocumentConstraint(t *testing.T) {
	t.Parallel()

	got, err := baonuoMerchantOrderID("sub2_20260526AbCd1234")
	require.NoError(t, err)
	require.Equal(t, "sub220260526AbCd1234", got)
	require.Len(t, got, 20)

	got, err = baonuoMerchantOrderID("****")
	require.NoError(t, err)
	require.Regexp(t, `^[A-Za-z0-9]{10,50}$`, got)
}

func TestBaonuoQueryOrderMapsPaidStatus(t *testing.T) {
	t.Parallel()

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/queryOrderV2", r.URL.Path)
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		values, err := url.ParseQuery(string(body))
		require.NoError(t, err)
		require.Equal(t, "sub2_order", values.Get("orderId"))
		responseData := map[string]string{
			"merchantId": "10086",
			"orderId":    "sub2_order",
			"status":     "paid",
			"amount":     "12.34",
			"msg":        "已支付",
		}
		responseData["sign"] = baonuoSign(responseData, "secret")
		_, _ = fmt.Fprintf(w, `{"code":200,"msg":"查询成功","data":{"merchantId":"%s","orderId":"%s","status":"%s","amount":"%s","msg":"%s","sign":"%s"}}`,
			responseData["merchantId"], responseData["orderId"], responseData["status"], responseData["amount"], responseData["msg"], responseData["sign"])
	}))
	defer server.Close()

	prov := mustTestBaonuoProvider(t, server)
	resp, err := prov.QueryOrder(context.Background(), "sub2_order")
	require.NoError(t, err)
	require.Equal(t, payment.ProviderStatusPaid, resp.Status)
	require.InDelta(t, 12.34, resp.Amount, 0.0001)
	require.Equal(t, "10086", resp.Metadata["merchant_id"])
	require.Equal(t, "paid", resp.Metadata["status"])
}

func TestBaonuoVerifyNotification(t *testing.T) {
	t.Parallel()

	prov, err := NewBaonuo("1", map[string]string{
		"merchantId":  "10086",
		"apiKey":      "secret",
		"channelType": "88",
		"notifyUrl":   "https://merchant.example.com/api/v1/payment/webhook/baonuo",
	})
	require.NoError(t, err)
	params := map[string]string{
		"merchantId": "10086",
		"orderId":    "sub2_order",
		"amount":     "12.34",
		"status":     "ok",
	}
	params["sign"] = baonuoSign(params, "secret")
	raw := encodeForm(params)

	notification, err := prov.VerifyNotification(context.Background(), raw, nil)
	require.NoError(t, err)
	require.Equal(t, "sub2_order", notification.TradeNo)
	require.Equal(t, "sub2_order", notification.OrderID)
	require.Equal(t, payment.NotificationStatusSuccess, notification.Status)
	require.InDelta(t, 12.34, notification.Amount, 0.0001)
	require.Equal(t, "10086", notification.Metadata["merchant_id"])
	require.Equal(t, "ok", notification.Metadata["status"])

	badParams := cloneStringMap(params)
	badParams["sign"] = "bad"
	_, err = prov.VerifyNotification(context.Background(), encodeForm(badParams), nil)
	require.ErrorContains(t, err, "invalid signature")
}

func TestBaonuoRefundUnsupported(t *testing.T) {
	t.Parallel()

	prov, err := NewBaonuo("1", map[string]string{
		"merchantId":  "10086",
		"apiKey":      "secret",
		"channelType": "88",
	})
	require.NoError(t, err)
	_, err = prov.Refund(context.Background(), payment.RefundRequest{OrderID: "sub2_order", Amount: "1.00"})
	require.ErrorContains(t, err, "not supported")
}

func mustTestBaonuoProvider(t *testing.T, server *httptest.Server) *Baonuo {
	t.Helper()
	prov, err := NewBaonuo("1", map[string]string{
		"merchantId":  "10086",
		"apiKey":      "secret",
		"channelType": "88",
		"apiBase":     baonuoDefaultAPIBase,
		"notifyUrl":   "https://merchant.example.com/api/v1/payment/webhook/baonuo",
	})
	require.NoError(t, err)
	prov.config["apiBase"] = server.URL
	prov.httpClient = server.Client()
	return prov
}

func flattenValues(values url.Values) map[string]string {
	out := make(map[string]string, len(values))
	for key := range values {
		out[key] = values.Get(key)
	}
	return out
}

func encodeForm(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	return strings.ReplaceAll(values.Encode(), "+", "%20")
}
