CREATE TABLE IF NOT EXISTS openai_codex_usage_snapshots (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    sampled_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    used_5h_percent NUMERIC(10,4),
    reset_5h_at TIMESTAMPTZ,
    window_5h_minutes INTEGER,
    used_7d_percent NUMERIC(10,4),
    reset_7d_at TIMESTAMPTZ,
    window_7d_minutes INTEGER,
    raw JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE INDEX IF NOT EXISTS idx_openai_codex_usage_snapshots_account_sampled_at
    ON openai_codex_usage_snapshots(account_id, sampled_at DESC);

CREATE INDEX IF NOT EXISTS idx_openai_codex_usage_snapshots_account_reset_7d
    ON openai_codex_usage_snapshots(account_id, reset_7d_at, sampled_at);
