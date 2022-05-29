create table users
(
    id_user    serial primary key,
    username   varchar(16) unique  not null,
    email      varchar(255) unique not null,
    created_on timestamp           not null,
    last_login timestamp
);
