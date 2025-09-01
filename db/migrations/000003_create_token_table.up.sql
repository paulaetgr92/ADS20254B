CREATE TABLE token (
                       id BIGSERIAL PRIMARY KEY,
                       user_id BIGINT NOT NULL,
                       token TEXT NOT NULL,
                       expires_at TIMESTAMP NOT NULL,
                       active BOOLEAN NOT NULL DEFAULT TRUE,
                       created_at TIMESTAMP NOT NULL DEFAULT now(),
                       access_id BIGINT NOT NULL,
                       CONSTRAINT fk_user FOREIGN KEY (user_id)
                           REFERENCES cadastro (id)
);
