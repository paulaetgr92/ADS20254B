CREATE table cadastro (
                       id BIGSERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(100) NOT NULL,
    AcessID BIGINT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
