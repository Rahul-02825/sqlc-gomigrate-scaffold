create table users(
    id int primary key,
    name text not null,
    mail text not null unique
);