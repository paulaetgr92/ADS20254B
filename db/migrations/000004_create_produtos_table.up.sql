CREATE TABLE PRODUTO (
    id_roupa BIGSERIAL PRIMARY KEY UNIQUE NOT NULL,
    categoria varchar (100) not null ,
    tamanho varchar(5) not null ,
    cores varchar,
    tempo_valor double precision,
    status varchar (100) ,
    localizacao varchar (100)
);





