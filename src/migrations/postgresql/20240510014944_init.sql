-- +goose Up
create extension if not exists "uuid-ossp";


create table if not exists address
(
    id      uuid primary key unique,
    country varchar(255) not null,
    city    varchar(255) not null,
    street  varchar(255) not null
);

create table if not exists client
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_name       varchar(255) not null,
    client_surname    varchar(255) not null,
    birthday          timestamp    not null,
    gender            varchar(255) not null,
    registration_date timestamp    not null,
    address_id        uuid         not null,

    constraint fk_address_id_address
        foreign key (address_id) references address (id)
);

create table if not exists images
(
    id    uuid primary key unique,
    image bytea
);

create table if not exists supplier
(
    id           uuid primary key unique,
    name         varchar(255) not null,
    address_id   uuid         not null,
    phone_number varchar(255) not null,

    constraint fk_address_id_address
        foreign key (address_id) references address (id),
    constraint ch_phone_number
        check (phone_number ~* '^\+7\d{10}$'
            )
);

create table if not exists product
(
    id               uuid primary key unique,
    name             varchar(255) not null,
    category         varchar(255) not null,
    price            money,
    available_stock  integer,
    last_update_date timestamp default now(),
    supplier_id      uuid         not null,
    image_id         uuid         not null,


    constraint fk_supplier_id_supplier
        foreign key (supplier_id) references supplier (id),
    constraint fk_image_id_images
        foreign key (image_id) references address (id)
);


-- +goose Down
drop table if exists client cascade;
drop table if exists product cascade;
drop table if exists supplier cascade;
drop table if exists address cascade;
drop table if exists images cascade;

