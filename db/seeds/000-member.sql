insert into 
member (member_id, username, first_name, middle_name, surname, email, joined_date) values
    ('037b4897-8a2a-46b6-8ed7-47a555bb40f2', 'NormUser', 'Bruce', 'C', 'Wayne', 'user@notreal.com', now()),
    ('b9add00b-1956-44ed-9565-6026260419aa', 'AdminUser', 'Mr.', '', 'Batman', 'admin@notreal.com', now())
on conflict do nothing;
