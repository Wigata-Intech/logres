-- migrate: up
-- create admin_refresh_tokens

CREATE TABLE IF NOT EXISTS admin_refresh_tokens (
    id            TEXT    NOT NULL PRIMARY KEY,
    admin_user_id TEXT    NOT NULL,
    token_hash    TEXT    NOT NULL,
    family_id     TEXT    NOT NULL,
    replaced_by   TEXT,
    expires_at    DATETIME NOT NULL,
    revoked_at    DATETIME,
    created_at    DATETIME NOT NULL,
    -- metadata: JSON envelope, currently {device: {user_agent, ip_address, device_label}}; future per-token keys (location/login/risk) land as sibling top-level keys with no schema change. device.ip_address is PII (GDPR/UU PDP). Justification: detect admin account takeover; operator sees own sessions. Deletion: purged with the row on logout/rotation-expiry and via admin_user erasure (keyed on admin_user_id). Never logged.
    metadata      TEXT NOT NULL DEFAULT '{}',
    last_used_at  DATETIME NOT NULL
);

-- No foreign keys; admin_user_id/family_id are indexed columns, not constraints.
CREATE UNIQUE INDEX IF NOT EXISTS idx_admin_refresh_tokens_token_hash
    ON admin_refresh_tokens (token_hash);

CREATE INDEX IF NOT EXISTS idx_admin_refresh_tokens_admin_user_id
    ON admin_refresh_tokens (admin_user_id);

CREATE INDEX IF NOT EXISTS idx_admin_refresh_tokens_family_id
    ON admin_refresh_tokens (family_id);
