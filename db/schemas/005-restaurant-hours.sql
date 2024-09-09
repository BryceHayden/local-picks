CREATE TYPE day AS ENUM ('Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat');

-- Note this is a simplistic usecase, and doesn't account for holidays or special hours.
create table if not exists 
restaurant_hours
(
    restaurant_id integer references restaurant(restaurant_id),
    day_of_week day not null,
    opening_time time not null,
    closing_time time not null,
    primary key (restaurant_id, day_of_week, opening_time)
);
