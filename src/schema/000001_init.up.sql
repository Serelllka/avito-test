CREATE DOMAIN UINT4 AS BIGINT
CHECK(VALUE > 0);

CREATE TABLE IF NOT EXISTS users_account
(
    id              SERIAL          NOT NULL UNIQUE,
    name            VARCHAR(255)    NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions
(
    id                  SERIAL                              NOT NULL UNIQUE,
    producer_id         INT REFERENCES users_account(id),
    consumer_id         INT REFERENCES users_account(id),
    amount              UINT4                               NOT NULL,
    transaction_type    INT                                 NOT NULL,
    description         VARCHAR(255)
);

