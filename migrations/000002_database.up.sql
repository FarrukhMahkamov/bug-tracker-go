CREATE TABLE users 
(
    id    BIGSERIAL    NOT NULL UNIQUE,
    name  VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    image TEXT DEFAULT NULL,
    role VARCHAR(50) DEFAULT 'engineer'
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

CREATE TABLE projects
(
    id BIGSERIAL NOT NULL UNIQUE,
    project_name VARCHAR(255)
);

CREATE TABLE bugs
(
    id BIGSERIAL NOT NULL UNIQUE,
    bug_title VARCHAR(255) NOT NULL,
    bug_description TEXT,
    is_completed BOOLEAN DEFAULT FALSE,
    status_id INT,
    category_id INT,
    project_id INT,
    FOREIGN KEY(project_id) REFERENCES projects(id) ON DELETE SET NULL,
    FOREIGN KEY(status_id) REFERENCES statuses(id) ON DELETE SET NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE bugs_tags
(
    id BIGSERIAL NOT NULL UNIQUE,
    bug_id INT NOT NULL REFERENCES bugs(id) ON DELETE CASCADE,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE
);

CREATE TABLE projects_users
(
    id BIGSERIAL NOT NULL UNIQUE,
    poroject_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE bugs_users
(
    id BIGSERIAL NOT NULL UNIQUE,
    bug_id INT NOT NULL REFERENCES bugs(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE teams_users
(
    id BIGSERIAL NOT NULL UNIQUE,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE projects_teams (
    id BIGSERIAL NOT NULL UNIQUE,
    project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE
);

CREATE TABLE bugs_teams (
    id BIGSERIAL NOT NULL UNIQUE,
    bug_id INT NOT NULL REFERENCES bugs(id) ON DELETE CASCADE,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE
);
