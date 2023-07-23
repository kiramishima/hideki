CREATE TABLE IF NOT EXISTS model_has_roles(
    role_id BIGINT NOT NULL,
    model_type VARCHAR(255) NOT NULL ,
    model_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, model_id, model_type),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);