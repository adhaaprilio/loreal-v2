CREATE SEQUENCE user_id_seq;

CREATE TABLE IF NOT EXISTS users(
   user_id  VARCHAR(255) PRIMARY KEY DEFAULT nextval('user_id_seq'),
   username VARCHAR (15) UNIQUE NOT NULL,
   name VARCHAR (50) NOT NULL,
   password VARCHAR (15) NOT NULL
);