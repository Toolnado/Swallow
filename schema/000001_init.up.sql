CREATE TABLE users (
    id serial not null unique,
    username varchar(255) not null unique,
    password varchar not null,
    first_name varchar(255),
    last_name varchar(255),
    create_at timestamp not null unique
);

CREATE TABLE posts(
    id serial not null unique,
    post_title varchar(255) not null unique,
    post_description varchar(255) not null unique,
    user_id int references users(id) on delete cascade not null
);