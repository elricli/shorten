BEGIN;

CREATE TABLE shorten_url (
    key text NOT NULL,
    short_url text NOT NULL,
    long_url text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (key)
);

END;
