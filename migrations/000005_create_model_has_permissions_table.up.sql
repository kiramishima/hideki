CREATE TABLE IF NOT EXISTS model_has_permissions(
    permission_id BIGINT NOT NULL,
    model_type VARCHAR(255) NOT NULL ,
    model_id BIGINT NOT NULL,
    PRIMARY KEY (permission_id, model_id, model_type),
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
);