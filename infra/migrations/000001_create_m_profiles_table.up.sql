CREATE TABLE m_profiles (
   id VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
   user_id VARCHAR(64),
   quotes VARCHAR(128),
   profesi VARCHAR(60),
   created_at DECIMAL NOT NULL,
   created_by VARCHAR(64),
   updated_at DECIMAL NOT NULL,
   updated_by VARCHAR(64),
   deleted_at DECIMAL,
   deleted_by VARCHAR(64)
);