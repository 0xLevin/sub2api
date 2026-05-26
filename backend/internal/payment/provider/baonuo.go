package provider

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/payment"
)

const (
	baonuoDefaultAPIBase = "https://baonuo.roseland.life"
	baonuoHTTPTimeout    = 10 * time.Second
	baonuoMaxBodySize    = 1 << 20
	baonuoMaxSummaryLen  = 512

	baonuoSuccessCode  = 200
	baonuoStatusPaid   = "paid"
	baonuoStatusUnpaid = "unpaid"
	baonuoNotifyOK     = "ok"
	baonuoIsFormJSON   = "2"
)

// Baonuo implements payment.Provider for BaoNuo's form-encoded gateway API.
type Baonuo struct {
	instanceID string
	config     map[string]string
	httpClient *http.Client
}

func NewBaonuo(instanceID string, config map[string]string) (*Baonuo, error) {
	for _, k := range []string{"merchantId", "apiKey", "channelType"} {
		if strings.TrimSpace(config[k]) == "" {
			return nil, fmt.Errorf("baonuo config missing required key: %s", k)
		}
	}
	cfg := cloneStringMap(config)
	if strings.TrimSpace(cfg["apiBase"]) == "" {
		cfg["apiBase"] = baonuoDefaultAPIBase
	}
	apiBase, err := normalizeBaonuoAPIBase(cfg["apiBase"])
	if err != nil {
		return nil, err
	}
	cfg["apiBase"] = apiBase
	currency, err := payment.NormalizePaymentCurrency(cfg["currency"])
	if err != nil {
		return nil, fmt.Errorf("baonuo config currency: %w", err)
	}
	cfg["currency"] = currency
	return &Baonuo{
		instanceID: instanceID,
		config:     cfg,
		httpClient: &http.Client{Timeout: baonuoHTTPTimeout},
	}, nil
}

func normalizeBaonuoAPIBase(raw string) (string, error) {
	base := strings.TrimSpace(raw)
	if base == "" {
		base = baonuoDefaultAPIBase
	}
	parsed, err := url.Parse(base)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" {
		return "", fmt.Errorf("baonuo apiBase must be an HTTPS URL")
	}
	parsed.RawQuery = ""
	parsed.Fragment = ""
	parsed.RawPath = ""
	parsed.Path = trimBaonuoEndpointPath(parsed.Path)
	return strings.TrimRight(parsed.String(), "/"), nil
}

func trimBaonuoEndpointPath(path string) string {
	path = strings.TrimRight(strings.TrimSpace(path), "/")
	lower := strings.ToLower(path)
	for _, endpoint := range []string{"/api/neworder", "/api/queryorderv2", "/api/queryorder", "/api/querybalance"} {
		if strings.HasSuffix(lower, endpoint) {
			return strings.TrimRight(path[:len(path)-len(endpoint)], "/")
		}
	}
	return path
}

func (b *Baonuo) Name() string        { return "BaoNuo" }
func (b *Baonuo) ProviderKey() string { return payment.TypeBaonuo }
func (b *Baonuo) SupportedTypes() []payment.PaymentType {
	return []payment.PaymentType{payment.TypeBaonuo}
}

func (b *Baonuo) MerchantIdentityMetadata() map[string]string {
	if b == nil {
		return nil
	}
	metadata := map[string]string{"currency": b.currency()}
	if merchantID := strings.TrimSpace(b.config["merchantId"]); merchantID != "" {
		metadata["merchant_id"] = merchantID
	}
	if channelType := strings.TrimSpace(b.config["channelType"]); channelType != "" {
		metadata["channel_type"] = channelType
	}
	return metadata
}

func (b *Baonuo) currency() string {
	if b == nil {
		return payment.DefaultPaymentCurrency
	}
	currency, err := payment.NormalizePaymentCurrency(b.config["currency"])
	if err != nil {
		return payment.DefaultPaymentCurrency
	}
	return currency
}

func (b *Baonuo) CreatePayment(ctx context.Context, req payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	if strings.TrimSpace(req.OrderID) == "" {
		return nil, fmt.Errorf("baonuo create payment: missing order id")
	}
	amount := strings.TrimSpace(req.Amount)
	if _, err := strconv.ParseFloat(amount, 64); err != nil || amount == "" {
		return nil, fmt.Errorf("baonuo create payment: invalid amount %s", req.Amount)
	}
	notifyURL, returnURL := b.resolveURLs(req)
	if notifyURL == "" {
		return nil, fmt.Errorf("baonuo create payment: notifyUrl is required")
	}

	params := map[string]string{
		"merchantId":  b.config["merchantId"],
		"orderId":     req.OrderID,
		"orderAmount": amount,
		"channelType": b.config["channelType"],
		"notifyUrl":   notifyURL,
		"returnUrl":   returnURL,
		"isForm":      baonuoIsFormJSON,
		"payer_ip":    strings.TrimSpace(req.ClientIP),
		"payer_id":    strings.TrimSpace(req.OrderID),
		"order_title": strings.TrimSpace(req.Subject),
		"order_body":  strings.TrimSpace(req.Subject),
	}
	params["sign"] = baonuoSign(params, b.config["apiKey"])

	body, err := b.post(ctx, "/api/newOrder", params)
	if err != nil {
		return nil, fmt.Errorf("baonuo create payment: %w", err)
	}
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			PayURL string `json:"payUrl"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("baonuo parse create response: %w", err)
	}
	if resp.Code != baonuoSuccessCode {
		return nil, fmt.Errorf("baonuo create failed: %s", strings.TrimSpace(resp.Msg))
	}
	payURL := strings.TrimSpace(resp.Data.PayURL)
	if payURL == "" {
		return nil, fmt.Errorf("baonuo create payment: missing payUrl")
	}
	return &payment.CreatePaymentResponse{
		TradeNo:  req.OrderID,
		PayURL:   payURL,
		Currency: b.currency(),
	}, nil
}

func (b *Baonuo) resolveURLs(req payment.CreatePaymentRequest) (string, string) {
	notifyURL := strings.TrimSpace(req.NotifyURL)
	if notifyURL == "" {
		notifyURL = strings.TrimSpace(b.config["notifyUrl"])
	}
	returnURL := strings.TrimSpace(req.ReturnURL)
	if returnURL == "" {
		returnURL = strings.TrimSpace(b.config["returnUrl"])
	}
	return notifyURL, returnURL
}

func (b *Baonuo) QueryOrder(ctx context.Context, tradeNo string) (*payment.QueryOrderResponse, error) {
	orderID := strings.TrimSpace(tradeNo)
	if orderID == "" {
		return nil, fmt.Errorf("baonuo query order: missing order id")
	}
	params := map[string]string{
		"merchantId": b.config["merchantId"],
		"orderId":    orderID,
	}
	params["sign"] = baonuoSign(params, b.config["apiKey"])

	body, err := b.post(ctx, "/api/queryOrderV2", params)
	if err != nil {
		return nil, fmt.Errorf("baonuo query order: %w", err)
	}
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			MerchantID string `json:"merchantId"`
			OrderID    string `json:"orderId"`
			Status     string `json:"status"`
			Amount     string `json:"amount"`
			Msg        string `json:"msg"`
			Sign       string `json:"sign"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("baonuo parse query response: %w", err)
	}
	if resp.Code != baonuoSuccessCode {
		return nil, fmt.Errorf("baonuo query failed: %s", strings.TrimSpace(resp.Msg))
	}
	if resp.Data.Sign != "" && !baonuoVerifySign(map[string]string{
		"merchantId": resp.Data.MerchantID,
		"orderId":    resp.Data.OrderID,
		"status":     resp.Data.Status,
		"amount":     resp.Data.Amount,
		"msg":        resp.Data.Msg,
		"sign":       resp.Data.Sign,
	}, b.config["apiKey"], resp.Data.Sign) {
		return nil, fmt.Errorf("baonuo query invalid signature")
	}

	status := payment.ProviderStatusPending
	switch strings.ToLower(strings.TrimSpace(resp.Data.Status)) {
	case baonuoStatusPaid:
		status = payment.ProviderStatusPaid
	case baonuoStatusUnpaid, "":
		status = payment.ProviderStatusPending
	default:
		status = payment.ProviderStatusFailed
	}
	amount, _ := strconv.ParseFloat(strings.TrimSpace(resp.Data.Amount), 64)
	return &payment.QueryOrderResponse{
		TradeNo:  orderID,
		Status:   status,
		Amount:   amount,
		Metadata: b.notificationMetadata(resp.Data.MerchantID, resp.Data.Status),
	}, nil
}

func (b *Baonuo) VerifyNotification(_ context.Context, rawBody string, _ map[string]string) (*payment.PaymentNotification, error) {
	values, err := url.ParseQuery(rawBody)
	if err != nil {
		return nil, fmt.Errorf("baonuo parse notification: %w", err)
	}
	params := make(map[string]string, len(values))
	for k := range values {
		params[k] = values.Get(k)
	}
	sign := strings.TrimSpace(params["sign"])
	if sign == "" {
		return nil, fmt.Errorf("baonuo notification missing sign")
	}
	if !baonuoVerifySign(params, b.config["apiKey"], sign) {
		return nil, fmt.Errorf("baonuo invalid signature")
	}
	status := payment.ProviderStatusFailed
	notifyStatus := strings.ToLower(strings.TrimSpace(params["status"]))
	if notifyStatus == baonuoNotifyOK {
		status = payment.NotificationStatusSuccess
	}
	amount, _ := strconv.ParseFloat(strings.TrimSpace(params["amount"]), 64)
	return &payment.PaymentNotification{
		TradeNo:  strings.TrimSpace(params["orderId"]),
		OrderID:  strings.TrimSpace(params["orderId"]),
		Amount:   amount,
		Status:   status,
		RawData:  rawBody,
		Metadata: b.notificationMetadata(params["merchantId"], params["status"]),
	}, nil
}

func (b *Baonuo) Refund(context.Context, payment.RefundRequest) (*payment.RefundResponse, error) {
	return nil, fmt.Errorf("baonuo refund is not supported")
}

func (b *Baonuo) notificationMetadata(merchantID, status string) map[string]string {
	metadata := b.MerchantIdentityMetadata()
	if metadata == nil {
		metadata = map[string]string{}
	}
	if merchantID = strings.TrimSpace(merchantID); merchantID != "" {
		metadata["merchant_id"] = merchantID
	}
	if status = strings.TrimSpace(status); status != "" {
		metadata["status"] = status
	}
	metadata["currency"] = b.currency()
	return metadata
}

func (b *Baonuo) post(ctx context.Context, path string, params map[string]string) ([]byte, error) {
	form := url.Values{}
	for k, v := range params {
		if strings.TrimSpace(v) == "" {
			continue
		}
		form.Set(k, v)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, b.config["apiBase"]+path, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := b.httpClient
	if client == nil {
		client = &http.Client{Timeout: baonuoHTTPTimeout}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(io.LimitReader(resp.Body, baonuoMaxBodySize))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, summarizeBaonuoResponse(body))
	}
	return body, nil
}

func summarizeBaonuoResponse(body []byte) string {
	summary := strings.Join(strings.Fields(string(body)), " ")
	if summary == "" {
		return "<empty>"
	}
	if len(summary) > baonuoMaxSummaryLen {
		return summary[:baonuoMaxSummaryLen] + "..."
	}
	return summary
}

func baonuoSign(params map[string]string, apiKey string) string {
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || baonuoIsEmptySignValue(v) {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, k+"="+params[k])
	}
	raw := strings.Join(parts, "&") + "&key=" + apiKey
	sum := md5.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func baonuoIsEmptySignValue(value string) bool {
	v := strings.TrimSpace(value)
	return v == "" || v == "0"
}

func baonuoVerifySign(params map[string]string, apiKey string, sign string) bool {
	expected := baonuoSign(params, apiKey)
	return hmac.Equal([]byte(expected), []byte(strings.ToLower(strings.TrimSpace(sign))))
}

var (
	_ payment.Provider                 = (*Baonuo)(nil)
	_ payment.MerchantIdentityProvider = (*Baonuo)(nil)
)
