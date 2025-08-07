CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS "tag"
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS "transactional_account"
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID,
    account_no VARCHAR,
    UNIQUE (user_id, account_no)
);

CREATE TABLE IF NOT EXISTS "recipient"
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_no VARCHAR UNIQUE,
    name       VARCHAR
);

CREATE TABLE IF NOT EXISTS "recipient_tag"
(
    id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tag_id UUID,
    CONSTRAINT fk_tag_id FOREIGN KEY (tag_id) REFERENCES tag (id),
    recipient_id UUID,
    CONSTRAINT fk_recipient_id FOREIGN KEY (recipient_id) REFERENCES recipient(id)
);

CREATE TABLE IF NOT EXISTS "transaction"
(
    id                       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transactional_account_id UUID,
    CONSTRAINT fk_transactional_account_id FOREIGN KEY (transactional_account_id) REFERENCES transactional_account (id),
    original_transaction_id  VARCHAR UNIQUE,
    recipient_id             UUID,
    amount                   FLOAT,
    valuta                   CHAR(3),
    transaction_date_time    TIMESTAMPTZ,
    created_at               TIMESTAMPTZ      DEFAULT NOW()
);

CREATE INDEX ix_transactional_account_id ON transaction (transactional_account_id);
CREATE UNIQUE INDEX uix_original_transaction_id ON transaction (original_transaction_id);
CREATE INDEX ux_transaction_date_time ON transaction (transaction_date_time)