CREATE TABLE cadastro (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          cpf VARCHAR(100) UNIQUE DEFAULT '',
                          cnpj VARCHAR(20) UNIQUE DEFAULT '',
                          email VARCHAR(100) NOT NULL UNIQUE,
                          celular VARCHAR(20) NOT NULL,
                          password VARCHAR(100) NOT NULL,
                          status VARCHAR(20) NOT NULL DEFAULT 'pendente',
                          activation_code VARCHAR,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
