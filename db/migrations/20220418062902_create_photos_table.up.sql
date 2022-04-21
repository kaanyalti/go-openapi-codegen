CREATE TABLE photos (
    id serial,
    uuid uuid NOT NULL,
    name VARCHAR NOT NULL,
    size INT NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL,
    deleted_at DATE,
    PRIMARY KEY (uuid)
);