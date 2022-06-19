CREATE TABLE user_balance
(
    id         serial    not null,
    user_id         VARCHAR       NOT NULL,
    current_balance INT       NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);