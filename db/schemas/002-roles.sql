CREATE TYPE role_type AS ENUM ('base', 'admin', 'manager');

create table if not exists 
roles
(
    role_id uuid primary key default gen_random_uuid(),
    name role_type not null,
    mod_date timestamptz not null default now()
);
