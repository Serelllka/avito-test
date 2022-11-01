CREATE DOMAIN UINT4 AS BIGINT
CHECK(VALUE > 0);

CREATE TABLE IF NOT EXISTS users_account
(
    id              SERIAL          NOT NULL UNIQUE,
    name            VARCHAR(255)    NOT NULL
);

CREATE TABLE IF NOT EXISTS services
(
    id                  SERIAL NOT NULL UNIQUE,
    title               VARCHAR(255) NOT NULL,
    description         VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS transactions
(
    id                  SERIAL                              NOT NULL UNIQUE,
    producer_id         INT REFERENCES users_account(id),
    consumer_id         INT REFERENCES users_account(id),
    service_id          INT REFERENCES services(id),
    order_id            INT,
    amount              UINT4                               NOT NULL,
    transaction_type    INT                                 NOT NULL,
    description         VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS reservations
(
    producer_id         INT REFERENCES users_account(id),
    service_id          INT REFERENCES services(id),
    order_id            INT,
    amount              UINT4                               NOT NULL,
    PRIMARY KEY (producer_id, service_id, order_id)
);

CREATE VIEW account_balance AS
(
    SELECT ua.id,
    (SELECT COALESCE(sum(innerTr.amount), 0) FROM transactions AS innerTr
        WHERE innerTr.consumer_id = ua.id) AS income,
    (SELECT COALESCE(sum(innerTr.amount), 0) FROM transactions AS innerTr
        WHERE innerTr.producer_id = ua.id) AS outcome,
    (SELECT COALESCE(sum(res.amount), 0) FROM reservations AS res
        WHERE res.producer_id = ua.id) AS reserved
    FROM users_account AS ua
);
