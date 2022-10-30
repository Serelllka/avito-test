CREATE TABLE users_account
(
    id              SERIAL          NOT NULL UNIQUE,
    name            VARCHAR(255)    NOT NULL
);

CREATE TABLE transactions
(
    id                  SERIAL                              NOT NULL UNIQUE,
    producer_id         INT REFERENCES users_account(id),
    consumer_id         INT REFERENCES users_account(id),
    transaction_type    INT                                 NOT NULL,
    description         VARCHAR(255)
);

