CREATE TABLE bars (
    id serial PRIMARY KEY,
    name text NOT NULL,
    LOCATION point NOT NULL,
    CONSTRAINT unq_name UNIQUE (name)
);

