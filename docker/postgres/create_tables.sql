CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;

CREATE TABLE markets(
  id serial not null primary key,
  long varchar(10) not null,
  lat varchar(10) not null,
  setcens varchar(15) not null,
  areap varchar(13) not null,
  coddist int not null,
  distrito citext not null,
  codsubpref int not null,
  subprefe varchar(25) not null,
  regiao5 citext not null,
  regiao8 varchar(7) not null,
  nome_feira citext not null,
  registro varchar(6) not null UNIQUE,
  logradouro varchar(34) not null,
  numero varchar(5),
  bairro citext,
  referencia varchar(60)
);

CREATE INDEX idx_market_cod_dist ON markets(registro);
CREATE INDEX idx_market_distrito ON markets(distrito);
CREATE INDEX idx_market_regiao5 ON markets(regiao5);
CREATE INDEX idx_market_nome_feira ON markets(nome_feira);
CREATE INDEX idx_market_bairro ON markets(bairro);


