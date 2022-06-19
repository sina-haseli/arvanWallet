CREATE TABLE user_transactions
(
    id         serial    not null,
    user_id         VARCHAR       NOT NULL,
    amount          INT       NOT NULL,
    description     VARCHAR,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);