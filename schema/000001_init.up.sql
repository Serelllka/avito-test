CREATE TABLE users
(
    id              SERIAL          NOT NULL UNIQUE,
    name            VARCHAR(255)    NOT NULL,
    username        VARCHAR(255)    NOT NULL UNIQUE,
    password_hash   VARCHAR(255)    NOT NULL
);

CREATE TABLE transactions
(
    id              SERIAL                      NOT NULL UNIQUE,
    producer_id     INT REFERENCES users(id)    NOT NULL,
    consumer_id     INT REFERENCES users(id)    NOT NULL
);
