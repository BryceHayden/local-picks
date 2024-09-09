create table if not exists 
passhash
(
    member_id uuid primary key references member(member_id),
    passhash text not null,
    mod_date timestamptz not null default now()
);
