CREATE TABLE sales (
                       id SERIAL PRIMARY KEY,
                       produto_id BIGINT NOT NULL REFERENCES produto(id_roupa),
                       quantidade INT NOT NULL,
                       tempo_valor DECIMAL(10,2) NOT NULL,
                       total DECIMAL(10,2) GENERATED ALWAYS AS (quantidade * tempo_valor ) STORED,
                       data_venda TIMESTAMP DEFAULT NOW()

);
