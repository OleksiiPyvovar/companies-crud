-- Create initial database structure.

CREATE TABLE companies (
        id SERIAL PRIMARY KEY,
        name VARCHAR (64),
	    code    VARCHAR (64),
	    country VARCHAR (64),
	    website VARCHAR (64),
	    phone   VARCHAR (64),
);