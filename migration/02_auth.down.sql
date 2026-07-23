
SET search_path TO biblioteca;

ALTER TABLE biblioteca.users
    DROP CONSTRAINT IF EXISTS uq_users_username;

ALTER TABLE biblioteca.users
    DROP COLUMN IF EXISTS username;
