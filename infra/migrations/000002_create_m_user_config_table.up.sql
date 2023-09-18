CREATE TABLE m_user_config(
   id VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
   profile_id VARCHAR(64),
   config_name VARCHAR(64),
   config_value JSONB,
   status VARCHAR(5),
   created_at DECIMAL NOT NULL,
   created_by VARCHAR(64),
   updated_at DECIMAL NOT NULL,
   updated_by VARCHAR(64),
   deleted_at DECIMAL,
   deleted_by VARCHAR(64),
   CONSTRAINT fk_m_profile
      FOREIGN KEY (profile_id)
    REFERENCES m_profiles(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);