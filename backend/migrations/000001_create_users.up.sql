-- migrations/000001_create_users.up.sql
CREATE TABLE users (
    id             UUID         PRIMARY KEY DEFAULT uuidv7(),
    first_name     VARCHAR(100) NOT NULL,
    last_name      VARCHAR(100) NOT NULL,
    cpf            VARCHAR(11)  NOT NULL UNIQUE,
    email          VARCHAR(255) NOT NULL UNIQUE,
    phone          VARCHAR(15),
    password_hash  VARCHAR(255) NOT NULL,
    role           VARCHAR(50)  NOT NULL DEFAULT 'owner',
    is_active      BOOLEAN      NOT NULL DEFAULT TRUE,
    email_verified BOOLEAN      NOT NULL DEFAULT FALSE,
    last_login_at  TIMESTAMPTZ,
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMPTZ
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_cpf ON users(cpf);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);