package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	defaultAPIBase     = "https://baonuo.roseland.life"
	defaultMerchantID  = "10084"
	defaultChannelType = "8003"
	defaultAmount      = "1.00"
	defaultIsForm      = "2"
	httpTimeout        = 15 * time.Second
	maxBodySize        = 1 << 20
)

type createOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PayURL string `json:"payUrl"`
	} `json:"data"`
}

func main() {
	var (
		apiBase     = flag.String("api-base", envDefault("BAONUO_API_BASE", defaultAPIBase), "BaoNuo API base URL")
		merchantID  = flag.String("merchant-id", envDefault("BAONUO_MERCHANT_ID", defaultMerchantID), "BaoNuo merchantId")
		apiKey      = flag.String("api-key", os.Getenv("BAONUO_API_KEY"), "BaoNuo signing key; prefer BAONUO_API_KEY")
		channelType = flag.String("channel-type", envDefault("BAONUO_CHANNEL_TYPE", defaultChannelType), "BaoNuo channelType")
		amount      = flag.String("amount", envDefault("BAONUO_AMOUNT", defaultAmount), "orderAmount, for example 1.00")
		orderID     = flag.String("order-id", os.Getenv("BAONUO_ORDER_ID"), "merchant order id; generated if empty")
		notifyURL   = flag.String("notify-url", os.Getenv("BAONUO_NOTIFY_URL"), "notifyUrl, required by BaoNuo")
		returnURL   = flag.String("return-url", os.Getenv("BAONUO_RETURN_URL"), "returnUrl, optional")
		clientIP    = flag.String("client-ip", os.Getenv("BAONUO_CLIENT_IP"), "payer_ip, optional")
		title       = flag.String("title", envDefault("BAONUO_TITLE", "Sub2API BaoNuo test order"), "order_title/order_body")
	)
	flag.Parse()

	if err := run(context.Background(), createOrderInput{
		APIBase:     *apiBase,
		MerchantID:  *merchantID,
		APIKey:      *apiKey,
		ChannelType: *channelType,
		Amount:      *amount,
		OrderID:     *orderID,
		NotifyURL:   *notifyURL,
		ReturnURL:   *returnURL,
		ClientIP:    *clientIP,
		Title:       *title,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}

type createOrderInput struct {
	APIBase     string
	MerchantID  string
	APIKey      string
	ChannelType string
	Amount      string
	OrderID     string
	NotifyURL   string
	ReturnURL   string
	ClientIP    string
	Title       string
}

func run(ctx context.Context, in createOrderInput) error {
	apiBase, err := normalizeAPIBase(in.APIBase)
	if err != nil {
		return err
	}
	if strings.TrimSpace(in.MerchantID) == "" {
		return errors.New("merchant-id is required")
	}
	if strings.TrimSpace(in.APIKey) == "" {
		return errors.New("api-key is required; set BAONUO_API_KEY")
	}
	if strings.TrimSpace(in.ChannelType) == "" {
		return errors.New("channel-type is required")
	}
	if strings.TrimSpace(in.NotifyURL) == "" {
		return errors.New("notify-url is required; set BAONUO_NOTIFY_URL")
	}
	if _, err := strconv.ParseFloat(strings.TrimSpace(in.Amount), 64); err != nil {
		return fmt.Errorf("amount must be numeric: %w", err)
	}

	orderID := strings.TrimSpace(in.OrderID)
	if orderID == "" {
		orderID = "S2" + time.Now().UTC().Format("20060102150405") + randomishSuffix()
	}
	merchantOrderID, err := merchantOrderID(orderID)
	if err != nil {
		return err
	}

	params := map[string]string{
		"merchantId":  strings.TrimSpace(in.MerchantID),
		"orderId":     merchantOrderID,
		"orderAmount": strings.TrimSpace(in.Amount),
		"channelType": strings.TrimSpace(in.ChannelType),
		"notifyUrl":   strings.TrimSpace(in.NotifyURL),
		"returnUrl":   strings.TrimSpace(in.ReturnURL),
		"isForm":      defaultIsForm,
		"payer_ip":    strings.TrimSpace(in.ClientIP),
		"order_title": strings.TrimSpace(in.Title),
		"order_body":  strings.TrimSpace(in.Title),
	}
	params["sign"] = sign(params, in.APIKey)

	form := url.Values{}
	for k, v := range params {
		if strings.TrimSpace(v) == "" {
			continue
		}
		form.Set(k, v)
	}

	fmt.Println("POST", apiBase+"/api/newOrder")
	fmt.Println("orderId:", merchantOrderID)
	fmt.Println("request form:")
	for _, k := range sortedValueKeys(form) {
		fmt.Printf("  %s=%s\n", k, form.Get(k))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiBase+"/api/newOrder", bytes.NewBufferString(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := (&http.Client{Timeout: httpTimeout}).Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBodySize))
	if err != nil {
		return err
	}
	fmt.Println("http status:", resp.Status)
	fmt.Println("raw response:", string(body))

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("non-2xx response: %s", resp.Status)
	}
	var parsed createOrderResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return fmt.Errorf("parse response json: %w", err)
	}
	if parsed.Code != 200 {
		return fmt.Errorf("baonuo create failed: code=%d msg=%s", parsed.Code, parsed.Msg)
	}
	if strings.TrimSpace(parsed.Data.PayURL) == "" {
		return errors.New("baonuo create succeeded but payUrl is empty")
	}
	fmt.Println("payUrl:", parsed.Data.PayURL)
	return nil
}

func normalizeAPIBase(raw string) (string, error) {
	base := strings.TrimSpace(raw)
	if base == "" {
		base = defaultAPIBase
	}
	parsed, err := url.Parse(base)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" {
		return "", errors.New("api-base must be an HTTPS URL")
	}
	parsed.RawQuery = ""
	parsed.Fragment = ""
	parsed.RawPath = ""
	parsed.Path = trimEndpointPath(parsed.Path)
	return strings.TrimRight(parsed.String(), "/"), nil
}

func trimEndpointPath(path string) string {
	path = strings.TrimRight(strings.TrimSpace(path), "/")
	lower := strings.ToLower(path)
	for _, endpoint := range []string{"/api/neworder", "/api/queryorderv2", "/api/queryorder", "/api/querybalance"} {
		if strings.HasSuffix(lower, endpoint) {
			return strings.TrimRight(path[:len(path)-len(endpoint)], "/")
		}
	}
	return path
}

func sign(params map[string]string, apiKey string) string {
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || isEmptySignValue(v) {
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

func isEmptySignValue(value string) bool {
	v := strings.TrimSpace(value)
	return v == "" || v == "0"
}

func merchantOrderID(orderID string) (string, error) {
	orderID = strings.TrimSpace(orderID)
	var b strings.Builder
	b.Grow(len(orderID))
	for _, ch := range orderID {
		if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			_, _ = b.WriteRune(ch)
		}
	}
	id := b.String()
	if len(id) >= 10 && len(id) <= 50 {
		return id, nil
	}
	sum := md5.Sum([]byte(orderID))
	id = "S2" + hex.EncodeToString(sum[:])
	if len(id) > 50 {
		id = id[:50]
	}
	if len(id) < 10 {
		return "", fmt.Errorf("invalid order id %q", orderID)
	}
	return id, nil
}

func sortedValueKeys(values url.Values) []string {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func envDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func randomishSuffix() string {
	sum := md5.Sum([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return strings.ToUpper(hex.EncodeToString(sum[:])[:8])
}
