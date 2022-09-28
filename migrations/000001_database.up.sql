CREATE TABLE users 
(
    id    BIGSERIAL    NOT NULL UNIQUE,
    name  VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    image TEXT DEFAULT NULL
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
    tag_name VARCHAR(255)
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

CREATE TABLE bugs
(
    id BIGSERIAL NOT NULL UNIQUE,
    bug_title VARCHAR(255) NOT NULL,
    bug_description TEXT,
    is_completed BOOLEAN DEFAULT FALSE,
    status_id INT,
    category_id INT,
    FOREIGN KEY(status_id) REFERENCES statuses(id) ON DELETE SET NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE bugs_tags
(
    id BIGSERIAL NOT NULL UNIQUE,
    bug_id INT NOT NULL REFERENCES bugs(id) ON DELETE CASCADE,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE
);