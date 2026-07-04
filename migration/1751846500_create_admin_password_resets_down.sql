-- migrate: down
-- drop admin_password_resets

DROP INDEX IF EXISTS idx_admin_password_resets_expires_at;
DROP INDEX IF EXISTS idx_admin_password_resets_admin_user_id;
DROP TABLE IF EXISTS admin_password_resets;
