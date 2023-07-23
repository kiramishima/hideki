CREATE TABLE IF NOT EXISTS roles(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(125) NOT NULL ,
    guard_name VARCHAR(125) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT roles_name_guard_name_unique
        UNIQUE (name, guard_name)
);