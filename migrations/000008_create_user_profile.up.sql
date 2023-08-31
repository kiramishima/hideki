CREATE TABLE IF NOT EXISTS user_profile(
    user_id BIGINT NOT NULL,
    name VARCHAR(25),
    first_lastname VARCHAR(25),
    second_lastname VARCHAR(25),
    bio VARCHAR(250),
    picture TEXT DEFAULT 'default_user.png',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);