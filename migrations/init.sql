create extension if not exists "uuid-ossp";

create table if not exists "user" (
    id uuid unique default uuid_generate_v4(),
    email varchar(255) not null,
    "password" text not null
);