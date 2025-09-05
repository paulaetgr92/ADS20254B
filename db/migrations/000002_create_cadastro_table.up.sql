CREATE TABLE cadastro (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          cnpj VARCHAR(20) UNIQUE,
                          email VARCHAR(100) NOT NULL,
                          celular VARCHAR(20) NOT NULL,
                          password VARCHAR(100) NOT NULL,
                          status VARCHAR(20) NOT NULL DEFAULT 'inativo',
                          activation_code VARCHAR(4),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
