-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    age INT
);

-- Insert a sample user record
INSERT INTO users (name, email, age)
VALUES ('John Doe', 'john.doe@example.com', 30)
ON CONFLICT (email) DO NOTHING;  -- Avoid errors if the record already exists
