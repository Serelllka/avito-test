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


CREATE VIEW account_balance AS
(
    SELECT ua.id,
    (SELECT sum(innerTr.amount) FROM transactions AS innerTr
        WHERE innerTr.consumer_id = ua.id) as income,
    (SELECT sum(innerTr.amount) FROM transactions AS innerTr
        WHERE innerTr.producer_id = ua.id) as outcome
    FROM users_account as ua
);
