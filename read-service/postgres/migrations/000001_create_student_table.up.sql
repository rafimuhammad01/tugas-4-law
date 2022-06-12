CREATE TABLE IF NOT EXISTS student (
    "id" serial primary key,
    "npm" varchar(64) not null unique,
    "name" varchar(128) not null
);