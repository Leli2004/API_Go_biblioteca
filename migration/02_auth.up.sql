
SET search_path TO biblioteca;

ALTER TABLE biblioteca.users
    ADD COLUMN username VARCHAR(50);

-- Gera um username temporário e único para usuários já existentes
UPDATE biblioteca.users
SET username = 'user_' || id
WHERE username IS NULL OR BTRIM(username) = '';

ALTER TABLE biblioteca.users
    ALTER COLUMN username SET NOT NULL;

ALTER TABLE biblioteca.users
    ADD CONSTRAINT uq_users_username UNIQUE (username);
