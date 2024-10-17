create table users
(
    id   serial primary key,
    name text not null,
    age  int  not null
);

create table items
(
    id   serial primary key,
    name text not null,
    price int not null
);
