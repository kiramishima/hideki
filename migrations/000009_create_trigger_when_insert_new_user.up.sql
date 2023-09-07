CREATE FUNCTION after_user_insert() RETURNS TRIGGER AS $$
BEGIN
INSERT INTO user_profile(user_id, name, first_lastname, second_lastname, bio)
VALUES(NEW.id, '', '', '', 'Default bio');

INSERT INTO model_has_roles(role_id, model_type, model_id) VALUES (1, 'User', NEW.id);

RETURN NULL;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER trig_after_user_insert
    AFTER INSERT ON users
    FOR EACH ROW EXECUTE PROCEDURE after_user_insert();