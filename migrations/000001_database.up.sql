CREATE TABLE users 
(
    id    BIGSERIAL    NOT NULL UNIQUE,
    name  VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    image TEXT
);

CREATE TABLE job_types
(
    id BIGSERIAL NOT NULL UNIQUE,
    job_type_name VARCHAR(255)
);