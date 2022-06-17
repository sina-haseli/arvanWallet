CREATE TABLE user_transactions
(
    id         serial    not null,
    user_id         INT       NOT NULL,
    amount          INT       NOT NULL,
    current_balance INT       NOT NULL,
    description     VARCHAR(150) CHARACTER SET utf8 COLLATE utf8_general_ci,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);