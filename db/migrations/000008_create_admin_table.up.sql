CREATE TABLE admin (
                       id BIGSERIAL PRIMARY KEY,
                        CNPJ varchar(20) UNIQUE NOT NULL ,
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                        password varchar (50) NOT NULL ,
                       created_at TIMESTAMP DEFAULT now()
);