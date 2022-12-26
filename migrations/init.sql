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
    file_type varchar(100) not null,
    file_id varchar(255) unique not null,
    count bigint not null check ( count >= 0 ),
    owner_email text not null,
    download_url text not null,
    preview_url text not null
    );

create table if not exists file_relation (
    id uuid primary key default uuid_generate_v4(),
    original_file_id varchar(255) not null references file(file_id),
    copied_file_id varchar(255) not null references file(file_id)
    );


create or replace function get_less_usages_link(fileId varchar(255)) returns table (download_url text, file_id varchar(255)) as $$
WITH all_copies AS (select copied_file_id, count, download_url from file_relation
     join file on file_relation.copied_file_id = file.file_id
where original_file_id = (select original_file_id from file_relation where copied_file_id = fileId))
SELECT download_url, copied_file_id from all_copies where count = (select min(count) from all_copies)
    $$ language sql;

insert into service_account(id, email, refresh_token, storage_limit, storage_usage) values('832f91a1-2bfd-47a5-b1b2-fd1f07276e44', 'ivannaviivan159@gmail.com', '1//0ckiJHdcZ4FJICgYIARAAGAwSNgF-L9IrKmWL4LolSMoY4W2d-SdTXKjAvxN8g4JoL3_2ZCeWioDam5SNEdV5yWcUYJ8CyHiTxg', 0 , 0);