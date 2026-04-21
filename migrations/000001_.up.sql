create schema todo;

create table todo.users(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL CHECK(char_length(full_name) BETWEEN 3 AND 100),
    phone_number VARCHAR(15) NOT NULL CHECK(
        char_length(phone_number) BETWEEN 10 AND 15
        AND
        phone_number ~ '^\+[0-9]+$'
    )
);

create table todo.tasks(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 3 AND 100),
    description VARCHAR(1000),
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    completed_at TIMESTAMPTZ,

    CHECK(
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at > created_at)
    ),

    user_id INT NOT NULL REFERENCES todo.users(id)
);