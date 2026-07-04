-- migrate: up
-- create admin_password_resets

CREATE TABLE IF NOT EXISTS admin_password_resets (
    id            TEXT    NOT NULL PRIMARY KEY,
    admin_user_id TEXT    NOT NULL,
    otp_hash      TEXT    NOT NULL,
    purpose       TEXT    NOT NULL,
    attempt_count INTEGER NOT NULL DEFAULT 0,
    expires_at    DATETIME NOT NULL,
    used_at       DATETIME,
    created_at    DATETIME NOT NULL
);

-- No foreign keys; admin_user_id is an indexed column, not a constraint.
CREATE INDEX IF NOT EXISTS idx_admin_password_resets_admin_user_id
    ON admin_password_resets (admin_user_id);

CREATE INDEX IF NOT EXISTS idx_admin_password_resets_expires_at
    ON admin_password_resets (expires_at);
