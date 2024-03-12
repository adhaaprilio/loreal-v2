CREATE TABLE IF NOT EXISTS users(
   user_id serial PRIMARY KEY,
   username VARCHAR (15) UNIQUE NOT NULL,
   name VARCHAR (50) NOT NULL,
   password VARCHAR (15) NOT NULL
);