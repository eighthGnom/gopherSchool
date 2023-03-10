create table users (
    id bigserial primary key ,
    email varchar unique not null ,
    enscripted_password varchar not null
)