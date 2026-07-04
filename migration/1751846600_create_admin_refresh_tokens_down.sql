-- migrate: down
-- drop admin_refresh_tokens

DROP INDEX IF EXISTS idx_admin_refresh_tokens_family_id;
DROP INDEX IF EXISTS idx_admin_refresh_tokens_admin_user_id;
DROP INDEX IF EXISTS idx_admin_refresh_tokens_token_hash;
DROP TABLE IF EXISTS admin_refresh_tokens;
