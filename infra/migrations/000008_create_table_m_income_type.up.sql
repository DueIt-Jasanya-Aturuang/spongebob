create type income_type_enum as enum('utama', 'sampingan', 'lainnya');
create type period_enum as enum('satu kali', 'harian', 'bulanan');

CREATE TABLE m_income_type
(
    id           VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id   VARCHAR(64),
    name         VARCHAR(64),
    description  VARCHAR(255) NULL,
    icon         VARCHAR(255),
    income_type  income_type_enum default 'lainnya',
    fixed_income BOOL NULL,
    periode      period_enum      default NULL,
    amount       DECIMAL NULL,
    created_at   DECIMAL     NOT NULL,
    created_by   VARCHAR(64),
    updated_at   DECIMAL     NOT NULL,
    updated_by   VARCHAR(64),
    deleted_at   DECIMAL,
    deleted_by   VARCHAR(64),
    constraint fk_m_profile
        foreign key (profile_id)
            references m_profiles (id)
            on delete cascade
            on update cascade
);