CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(25)  NOT NULL,
    last_name VARCHAR(25)  NOT NULL,
    email VARCHAR(75)  NOT NULL UNIQUE,
    hire_date DATE NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);