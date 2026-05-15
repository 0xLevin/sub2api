package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAccountCheckUsagePercentSchedulability(t *testing.T) {
	now := time.Date(2026, 5, 13, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		account *Account
		want    WindowCostSchedulability
	}{
		{
			name: "openai fresh 5h usage reaches limit",
			account: &Account{
				Platform: PlatformOpenAI,
				Type:     AccountTypeOAuth,
				Extra: map[string]any{
					"usage_percent_limit_5h": 80,
					"codex_usage_updated_at": now.Add(-time.Hour).Format(time.RFC3339),
					"codex_5h_used_percent":  80,
					"codex_5h_reset_at":      now.Add(time.Hour).Format(time.RFC3339),
				},
			},
			want: WindowCostNotSchedulable,
		},
		{
			name: "openai stale usage fails open",
			account: &Account{
				Platform: PlatformOpenAI,
				Type:     AccountTypeOAuth,
				Extra: map[string]any{
					"usage_percent_limit_5h": 80,
					"codex_usage_updated_at": now.Add(-(usagePercentLimitFreshTTL + time.Minute)).Format(time.RFC3339),
					"codex_5h_used_percent":  99,
					"codex_5h_reset_at":      now.Add(time.Hour).Format(time.RFC3339),
				},
			},
			want: WindowCostSchedulable,
		},
		{
			name: "openai expired reset fails open for that window",
			account: &Account{
				Platform: PlatformOpenAI,
				Type:     AccountTypeOAuth,
				Extra: map[string]any{
					"usage_percent_limit_5h": 80,
					"codex_usage_updated_at": now.Add(-time.Hour).Format(time.RFC3339),
					"codex_5h_used_percent":  99,
					"codex_5h_reset_at":      now.Add(-time.Minute).Format(time.RFC3339),
				},
			},
			want: WindowCostSchedulable,
		},
		{
			name: "anthropic fresh 5h utilization reaches limit",
			account: &Account{
				Platform: PlatformAnthropic,
				Type:     AccountTypeOAuth,
				Extra: map[string]any{
					"usage_percent_limit_5h":     80,
					"passive_usage_sampled_at":   now.Add(-time.Hour).Format(time.RFC3339),
					"session_window_utilization": 0.8,
				},
			},
			want: WindowCostNotSchedulable,
		},
		{
			name: "anthropic expired 5h window fails open for that window",
			account: &Account{
				Platform:         PlatformAnthropic,
				Type:             AccountTypeOAuth,
				SessionWindowEnd: usagePercentTestTimePtr(now.Add(-time.Minute)),
				Extra: map[string]any{
					"usage_percent_limit_5h":     80,
					"passive_usage_sampled_at":   now.Add(-time.Hour).Format(time.RFC3339),
					"session_window_utilization": 0.99,
				},
			},
			want: WindowCostSchedulable,
		},
		{
			name: "anthropic fresh 7d utilization reaches limit",
			account: &Account{
				Platform: PlatformAnthropic,
				Type:     AccountTypeSetupToken,
				Extra: map[string]any{
					"usage_percent_limit_7d":       70,
					"passive_usage_sampled_at":     now.Add(-time.Hour).Format(time.RFC3339),
					"passive_usage_7d_utilization": 0.7,
					"passive_usage_7d_reset":       now.Add(24 * time.Hour).Unix(),
				},
			},
			want: WindowCostNotSchedulable,
		},
		{
			name: "anthropic expired 7d reset fails open for that window",
			account: &Account{
				Platform: PlatformAnthropic,
				Type:     AccountTypeSetupToken,
				Extra: map[string]any{
					"usage_percent_limit_7d":       70,
					"passive_usage_sampled_at":     now.Add(-time.Hour).Format(time.RFC3339),
					"passive_usage_7d_utilization": 0.99,
					"passive_usage_7d_reset":       now.Add(-time.Minute).Unix(),
				},
			},
			want: WindowCostSchedulable,
		},
		{
			name: "unsupported account type ignores limits",
			account: &Account{
				Platform: PlatformOpenAI,
				Type:     AccountTypeAPIKey,
				Extra: map[string]any{
					"usage_percent_limit_5h": 1,
					"codex_usage_updated_at": now.Add(-time.Hour).Format(time.RFC3339),
					"codex_5h_used_percent":  100,
				},
			},
			want: WindowCostSchedulable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.account.CheckUsagePercentSchedulability(now))
		})
	}
}

func usagePercentTestTimePtr(t time.Time) *time.Time {
	return &t
}
