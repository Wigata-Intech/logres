-- migrate: up
-- create admin_users

CREATE TABLE IF NOT EXISTS admin_users (
    id            TEXT    NOT NULL PRIMARY KEY,
    full_name     TEXT    NOT NULL,
    email         TEXT    NOT NULL,
    password      TEXT    NOT NULL DEFAULT '',
    status        INTEGER  NOT NULL DEFAULT 0,
    last_login_at DATETIME,
    created_at    DATETIME NOT NULL,
    updated_at    DATETIME NOT NULL,
    deleted_at    DATETIME
);

-- One active admin per email; soft-deleted rows are excluded from the
-- uniqueness constraint. Relationships use indexed columns, never foreign keys.
CREATE UNIQUE INDEX IF NOT EXISTS idx_admin_users_email_active
    ON admin_users (email)
    WHERE deleted_at IS NULL;
