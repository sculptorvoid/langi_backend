CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE dictionaries
(
    id    serial       not null unique,
    title varchar(255) not null
);

CREATE TABLE words
(
    id          serial       not null unique,
    title       varchar(255) not null,
    translation varchar(255)
);

CREATE TABLE users_dictionaries
(
    id      serial                                             not null unique,
    user_id int references users (id) on delete cascade        not null,
    dict_id int references dictionaries (id) on delete cascade not null
);

CREATE TABLE dictionaries_words
(
    id      serial                                             not null unique,
    word_id int references words (id) on delete cascade        not null,
    dict_id int references dictionaries (id) on delete cascade not null
);
