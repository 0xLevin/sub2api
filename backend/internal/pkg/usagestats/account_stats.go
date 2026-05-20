package usagestats

// AccountStats 账号使用统计
//
// cost: 账号口径费用（使用 total_cost * account_rate_multiplier）
// standard_cost: 标准费用（使用 total_cost，不含倍率）
// user_cost: 用户/API Key 口径费用（使用 actual_cost，受分组倍率影响）
type AccountStats struct {
	Requests     int64   `json:"requests"`
	Tokens       int64   `json:"tokens"`
	Cost         float64 `json:"cost"`
	StandardCost float64 `json:"standard_cost"`
	UserCost     float64 `json:"user_cost"`
}

type OpenAICodexUsageCycle struct {
	StartTime            string   `json:"start_time"`
	EndTime              string   `json:"end_time"`
	ResetAt              *string  `json:"reset_at,omitempty"`
	DurationDays         float64  `json:"duration_days"`
	Requests             int64    `json:"requests"`
	Tokens               int64    `json:"tokens"`
	Cost                 float64  `json:"cost"`
	StandardCost         float64  `json:"standard_cost"`
	UserCost             float64  `json:"user_cost"`
	Equivalent7dTokens   int64    `json:"equivalent_7d_tokens"`
	Equivalent7dCost     float64  `json:"equivalent_7d_cost"`
	Equivalent7dUserCost float64  `json:"equivalent_7d_user_cost"`
	CostPer1MTokens      *float64 `json:"cost_per_1m_tokens,omitempty"`
	MaxUsed7dPercent     *float64 `json:"max_used_7d_percent,omitempty"`
	SampleCount          int64    `json:"sample_count"`
	Complete             bool     `json:"complete"`
}

type OpenAICodexCapacitySummary struct {
	CycleCount            int     `json:"cycle_count"`
	CompleteCycleCount    int     `json:"complete_cycle_count"`
	Median7dTokens        int64   `json:"median_7d_tokens"`
	Median7dCost          float64 `json:"median_7d_cost"`
	Median7dUserCost      float64 `json:"median_7d_user_cost"`
	P257dTokens           int64   `json:"p25_7d_tokens"`
	P257dCost             float64 `json:"p25_7d_cost"`
	P257dUserCost         float64 `json:"p25_7d_user_cost"`
	MedianCostPer1MTokens float64 `json:"median_cost_per_1m_tokens"`
}

type OpenAICodexCapacityStats struct {
	Summary *OpenAICodexCapacitySummary `json:"summary,omitempty"`
	Cycles  []OpenAICodexUsageCycle     `json:"cycles"`
}
