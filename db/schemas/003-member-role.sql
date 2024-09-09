create table if not exists 
member_role
(
    member_id uuid references member(member_id),
    role_id uuid references roles(role_id),
    primary key (member_id, role_id)
);
