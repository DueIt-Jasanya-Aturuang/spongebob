CREATE TABLE t_income_history
(
    id                         VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    income_id                  VARCHAR(64),
    income_amount              DECIMAL,
    description                VARCHAR(55),
    time_spending_history      DATE,
    show_time_spending_history VARCHAR(64),
    created_at                 DECIMAL     NOT NULL,
    created_by                 VARCHAR(64),
    updated_at                 DECIMAL     NOT NULL,
    updated_by                 VARCHAR(64),
    deleted_at                 DECIMAL,
    deleted_by                 VARCHAR(64),
    constraint fk_m_income_type
        foreign key (income_id)
            references m_income_type (id)
            on delete cascade
            on update cascade
);