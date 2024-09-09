insert into 
passhash (member_id, passhash, mod_date) values
    ('037b4897-8a2a-46b6-8ed7-47a555bb40f2', '$2y$10$wzXLen3Yfswk8m7fmYJ7r.Yf7ax3Nvusm0ecqnDl4Hubig.NabZm.',  now()),
    ('b9add00b-1956-44ed-9565-6026260419aa', '$2y$10$wzXLen3Yfswk8m7fmYJ7r.Yf7ax3Nvusm0ecqnDl4Hubig.NabZm.',  now())
on conflict do nothing;
