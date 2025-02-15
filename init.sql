-- Create the postgres role (if it doesn't exist)
DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'postgres') THEN
        CREATE ROLE postgres WITH SUPERUSER CREATEDB CREATEROLE LOGIN PASSWORD 'postgres';
    END IF;
END
$$;

-- Create the database (if it doesn't exist)
CREATE DATABASE booksdb;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE booksdb TO postgres;
