create table if not exists 
member
(
    member_id uuid primary key default gen_random_uuid(),
    username text not null,
    first_name text not null,
  	middle_name text,
  	surname text not null,
    email text not null unique,
    joined_date timestamptz not null default now()
);
