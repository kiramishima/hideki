CREATE TABLE IF NOT EXISTS permissions(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(125),
    guard_name VARCHAR(125),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT permissions_name_guard_name_unique
        UNIQUE (name, guard_name)
);