CREATE TABLE activation_code (
                                 id SERIAL PRIMARY KEY,
                                 activation_code VARCHAR NOT NULL,
                                 cadastro_id BIGINT NOT NULL,
                                 code VARCHAR(6) NOT NULL,
                                 status VARCHAR,
                                 expires_at TIMESTAMP NOT NULL,
                                 created_at TIMESTAMP DEFAULT NOW(),
                                 CONSTRAINT fk_cadastro FOREIGN KEY (cadastro_id) REFERENCES cadastro(id) ON DELETE CASCADE
);
