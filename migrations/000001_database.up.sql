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

CREATE TABLE teams
(
    id BIGSERIAL NOT NULL UNIQUE,
    team_name VARCHAR(255)
);

CREATE TABLE tags
(
    id BIGSERIAL NOT NULL UNIQUE,
    tags_name VARCHAR(255)
);

CREATE TABLE categories
(
    id BIGSERIAL NOT NULL UNIQUE,
    category_name VARCHAR(255)
);

CREATE TABLE statuses
(
    id BIGSERIAL NOT NULL UNIQUE,
    status_name VARCHAR(255)
);
