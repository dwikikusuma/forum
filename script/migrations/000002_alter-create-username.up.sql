ALTER TABLE users ADD username VARCHAR(100) NOT NULL;

Alter TABLE users ADD CONSTRAINT UNIQUE username_unique (username);
