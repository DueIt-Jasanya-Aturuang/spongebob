CREATE TABLE m_icon
(
    id         VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    title      VARCHAR(64),
    icon       VARCHAR(255),
    created_at DECIMAL     NOT NULL,
    created_by VARCHAR(64),
    updated_at DECIMAL     NOT NULL,
    updated_by VARCHAR(64),
    deleted_at DECIMAL,
    deleted_by VARCHAR(64)
);