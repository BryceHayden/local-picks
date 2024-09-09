insert into 
roles (role_id, name) values
    ('d29aa376-d6fd-48be-82a3-f78bfff3ed9b', 'admin'),
    ('9a2cfc25-9548-4ec6-a237-1d3e378637ac', 'manager')
on conflict do nothing;
