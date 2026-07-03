-- migrate: down
-- drop admin_users

DROP INDEX IF EXISTS idx_admin_users_email_active;
DROP TABLE IF EXISTS admin_users;
