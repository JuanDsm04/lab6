-- Start an anonymous block in PL/pgSQL
DO $$ 
BEGIN
    -- Check if the database 'mydatabase' does not exist
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'mydatabase') THEN
        -- If it doesn't exist, create the 'mydatabase' database
        CREATE DATABASE mydatabase;
    END IF;
END $$;

-- Connect to the 'mydatabase' database
\c mydatabase;

-- Start another anonymous block in PL/pgSQL
DO $$ 
BEGIN
    -- Check if the role 'myuser' does not exist
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'myuser') THEN
        -- If it doesn't exist, create the 'myuser' role with the provided password
        CREATE USER myuser WITH PASSWORD 'mypassword';
    END IF;
END $$;

-- Grant all privileges on the 'mydatabase' database to the 'myuser' role
GRANT ALL PRIVILEGES ON DATABASE mydatabase TO myuser;

-- Create the 'series' table if it does not exist
CREATE TABLE IF NOT EXISTS series (
    id SERIAL PRIMARY KEY,          -- Unique identifier for each series
    title VARCHAR(150),             -- Title of the series
    status VARCHAR(50),             -- Current status of the series
    last_episode_watched INT,       -- Number of the last episode watched
    total_episodes INT,             -- Total number of episodes in the series
    ranking INT                     -- Ranking assigned by the user to the series
);