CREATE TABLE h_notify_message
(
    id         VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    name       VARCHAR(64),
    status     VARCHAR(20),
    message    TEXT,
    created_at DECIMAL     NOT NULL,
    created_by VARCHAR(64),
    updated_at DECIMAL     NOT NULL,
    updated_by VARCHAR(64),
    deleted_at DECIMAL,
    deleted_by VARCHAR(64)
);