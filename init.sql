CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR (50) NOT NULL,
  last_name VARCHAR (50) NOT NULL,
  email VARCHAR (255) UNIQUE NOT NULL,
  password VARCHAR (255) NOT NULL
);

INSERT INTO users (first_name, last_name, email, password) VALUES ('vivek', 'mishra', 'vivek@example.com', 'test1234')