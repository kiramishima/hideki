ALTER TABLE users ADD CONSTRAINT email_validation CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$');
