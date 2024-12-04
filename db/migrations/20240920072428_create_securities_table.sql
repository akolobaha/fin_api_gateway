-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS securities
(
    ticker    VARCHAR(255) PRIMARY KEY,
    shortname VARCHAR(255),
    secname   VARCHAR(255)
);
COMMENT ON TABLE securities IS 'Эмитенты';
COMMENT ON COLUMN securities.ticker IS 'Тикер';
COMMENT ON COLUMN securities.shortname IS 'Краткое наименование';
COMMENT ON COLUMN securities.secname IS 'Полное наименование';

CREATE INDEX idx_security_ticker ON securities (ticker);

CREATE TYPE currency AS ENUM ('RUB', 'USD', 'EUR', 'CYN');
COMMENT ON TYPE currency IS 'Валюты';

CREATE TABLE "users"
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    email    VARCHAR(255) NOT NULL UNIQUE,
    telegram VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
COMMENT ON COLUMN "users".name IS 'Имя';
COMMENT ON COLUMN "users".email IS 'Email';
COMMENT ON COLUMN "users".telegram IS 'Телеграмм';

CREATE TABLE "user_tokens"
(
    id SERIAL PRIMARY KEY,
    token VARCHAR(255) NOT NULL UNIQUE,
    user_id int REFERENCES users(id),
    expiration_time TIMESTAMP
);

CREATE TYPE notification_method AS ENUM ('telegram', 'email', 'sms');
COMMENT ON TYPE notification_method IS 'Механизм уведомления';
CREATE TYPE financial_report AS ENUM ('msfo', 'rsbu');
COMMENT ON TYPE financial_report IS 'Финансовая отчетность';
CREATE type valuation_ratio AS ENUM ('pe', 'pbv', 'price', 'ps');
COMMENT ON TYPE valuation_ratio IS 'Коэффицент оценки';

CREATE TABLE user_targets
(
    id               SERIAL PRIMARY KEY,
    ticker           VARCHAR(255) REFERENCES securities (ticker) ON DELETE CASCADE,
    user_id          int REFERENCES users(id),
    valuation_ratio valuation_ratio,
    value DECIMAL(10, 2),
    financial_report financial_report default 'rsbu',
    achieved boolean default false,
    notification_method notification_method DEFAULT 'email'
);
COMMENT ON TABLE user_targets IS 'Цели пользователей по эмитентам';
COMMENT ON COLUMN user_targets.ticker IS 'Тикер';
COMMENT ON COLUMN user_targets.user_id IS 'ID пользователя';
COMMENT ON COLUMN user_targets.value IS 'Цель';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS user_targets;
DROP TABLE IF EXISTS user_tokens;
DROP TABLE IF EXISTS securities;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS currency;
DROP TYPE IF EXISTS notification_method;
DROP TYPE IF EXISTS financial_report;
DROP TYPE IF EXISTS valuation_ratio;
-- +goose StatementEnd
