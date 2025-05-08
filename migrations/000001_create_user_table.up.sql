CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL
);

INSERT INTO users (name, email, created_at) VALUES ('Alice', 'alice@example.com', now());
