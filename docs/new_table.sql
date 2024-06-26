CREATE TABLE memberships (
    membership_id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL,
    created_date TIMESTAMP NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    updated_date TIMESTAMP,
    updated_by VARCHAR(255)
);

CREATE TYPE contact_type AS ENUM ('email', 'phone');

CREATE TABLE contacts (
    contact_id SERIAL PRIMARY KEY NOT NULL,
    membership_id INTEGER REFERENCES memberships(membership_id) NOT NULL,
    contact_type contact_type NOT NULL,
    contact_value VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL,
    created_date TIMESTAMP NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    updated_date TIMESTAMP,
    updated_by VARCHAR(255)
);