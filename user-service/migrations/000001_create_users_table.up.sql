CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at VARCHAR(255),
    updated_at VARCHAR(255),
    deleted_at INT DEFAULT 0
);
