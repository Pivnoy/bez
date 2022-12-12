CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists "users", service_account, file, file_relation cascade;

create table if not exists "users" (
    id uuid primary key default uuid_generate_v4(),
    email text unique not null,
    refresh_token text unique not null
    );

create table if not exists service_account (
    id uuid primary key default uuid_generate_v4(),
    email text unique not null,
    refresh_token text unique not null,
    storage_limit bigint check ( storage_limit >= 0 ),
    storage_usage bigint check ( storage_usage >= 0 AND storage_limit >= storage_usage )
    );

create table if not exists file (
    id uuid primary key default uuid_generate_v4(),
    file_name text not null,
    file_type varchar(10) not null,
    file_id varchar(255) unique not null,
    count bigint not null check ( count >= 0 ),
    owner_email text not null
    );

create table if not exists file_relation (
    id uuid primary key default uuid_generate_v4(),
    original_file_id varchar(255) not null references file(file_id),
    copied_file_id varchar(255) not null references file(file_id)
    );

insert into service_account(id, email, refresh_token, storage_limit, storage_usage) values('832f91a1-2bfd-47a5-b1b2-fd1f07276e44', 'ivannaviivan159@gmail.com', '1//0cbyi8J5FgnIyCgYIARAAGAwSNgF-L9IrvhIyH60DzP2c0N-MMO8wVzQuhUO9MiZpNjg1GGPWNwOekGcL2X4iDqHYp6Mrmg9mCQ', 0 , 0);