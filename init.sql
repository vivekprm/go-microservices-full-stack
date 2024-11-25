CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR (50) NOT NULL,
  last_name VARCHAR (50) NOT NULL,
  email VARCHAR (255) UNIQUE NOT NULL,
  password VARCHAR (255) NOT NULL
);
\c todostore
CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  description VARCHAR (255),
  status VARCHAR (50) UNIQUE NOT NULL,
  created_by VARCHAR (255) NOT NULL,
  created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- INSERT INTO users (first_name, last_name, email, password) VALUES ('vivek', 'mishra', 'vivek@example.com', 'test1234')
