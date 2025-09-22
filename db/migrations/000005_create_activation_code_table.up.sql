CREATE TABLE activation_codes (
                                  id SERIAL PRIMARY KEY,
                                  cadastro_id BIGINT NOT NULL,
                                  code VARCHAR(6) NOT NULL,
                                  expires_at TIMESTAMP NOT NULL,
                                  created_at TIMESTAMP DEFAULT NOW(),
                                  CONSTRAINT fk_cadastro FOREIGN KEY (cadastro_id) REFERENCES cadastro(id) ON DELETE CASCADE
);
