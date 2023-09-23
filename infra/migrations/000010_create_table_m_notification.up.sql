CREATE TABLE m_notificatio
(
    id             VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id     VARCHAR(64),
    user_config_id VARCHAR(64),
    message        TEXT,
    title          VARCHAR(64),
    icon           VARCHAR(255),
    created_at     DECIMAL     NOT NULL,
    created_by     VARCHAR(64),
    updated_at     DECIMAL     NOT NULL,
    updated_by     VARCHAR(64),
    deleted_at     DECIMAL,
    deleted_by     VARCHAR(64),
    constraint fk_m_profile
        foreign key (profile_id)
            references m_profiles (id)
            on delete cascade
            on update cascade,
    constraint fk_m_user_config
        foreign key (profile_id)
            references m_user_config (id)
            on delete cascade
            on update cascade
);