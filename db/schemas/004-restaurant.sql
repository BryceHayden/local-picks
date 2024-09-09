-- Note we don't need to worry about someone stealing the restaurant_id as it isn't meant to be private
create table if not exists 
restaurant
(
    restaurant_id serial primary key,
    name text unique not null    
);
