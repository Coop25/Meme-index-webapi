CREATE TABLE IF NOT EXISTS files (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    contenttype VARCHAR(255) NOT NULL,
    url VARCHAR(255),
    description TEXT
);