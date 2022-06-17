CREATE TABLE user_transactions
(
    id         serial    not null,
    user_id         INT       NOT NULL,
    amount          INT       NOT NULL,
    description     VARCHAR(150),
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);