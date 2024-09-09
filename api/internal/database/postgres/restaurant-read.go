package postgres

import (
	"encoding/json"

	_ "github.com/lib/pq"

	"github.com/brycehayden/resume/internal/models"
)

func (psql PSQL) FetchRestaurants() ([]models.Restaurant, error){
	restaurants := []models.Restaurant{};

	rows, err := psql.DB.Queryx(`select * from restaurant;`)
	if err != nil { return nil, err }

	for rows.Next() {
		restaurant := &models.Restaurant{}
		err = rows.StructScan(&restaurant)
		if err != nil { return nil, err }
		
		restaurants = append(restaurants, *restaurant)
	}	
	
	return restaurants, nil	
}

func (psql PSQL) FetchRestaurantsByDay(day string) ([]models.Restaurant, error){
	restaurants := []models.Restaurant{};

	rows, err := psql.DB.Queryx(`
		select  r.restaurant_id, r.name
		from restaurant r
		left join restaurant_hours rh on rh.restaurant_id = r.restaurant_id
		where rh.day_of_week = cast($1 as day)
		GROUP BY r.restaurant_id, r.name;
	`, 
	day)
	if err != nil { return nil, err }
	
	for rows.Next() {
		restaurant := &models.Restaurant{}
		err = rows.StructScan(&restaurant)
		if err != nil { return nil, err }
		
		restaurants = append(restaurants, *restaurant)
	}	
	
	return restaurants, nil	
}

func (psql PSQL) FetchRestaurantsByDateTime(day string, time string) ([]models.Restaurant, error){
	restaurants := []models.Restaurant{};

	rows, err := psql.DB.Queryx(`
		select  r.restaurant_id, r.name
		from restaurant r
		left join restaurant_hours rh on rh.restaurant_id = r.restaurant_id
		where rh.day_of_week = cast($1 as day)
		and rh.opening_time <= $2
		and rh.closing_time >= $2
		GROUP BY r.restaurant_id, r.name;
	`, 
	day,
	time)
	if err != nil { return nil, err }
	
	for rows.Next() {
		restaurant := &models.Restaurant{}
		err = rows.StructScan(&restaurant)
		if err != nil { return nil, err }
		
		restaurants = append(restaurants, *restaurant)
	}	
	
	return restaurants, nil	
}

func (psql PSQL) FetchRestaurantDetails(desired_id string) (*models.RestaurantDetails, error){
	row, err := psql.DB.Query(`
		select 
    		r.restaurant_id, r.name,
			json_agg(
				distinct
				jsonb_build_object(
					'day', rh.day_of_week,
					'opening', rh.opening_time,
					'closing', rh.closing_time
				)
			) as hours			
		from restaurant r
		left join restaurant_hours rh on rh.restaurant_id = r.restaurant_id
		where r.restaurant_id = $1
		GROUP BY r.restaurant_id, r.name;
	`, 
	desired_id)
	if err != nil { return nil, err }
	
	var id *string
	var name *string
	
	//Couldn't find the restaurant by the given id
	if !row.Next(){ return nil, err }
	
	//Prepare the results to be saved 
	var hoursRaw []uint8
	err = row.Scan(&id, &name, &hoursRaw)
	if err != nil { return nil, err }
	
	var hours []models.RestaurantHours
	err = json.Unmarshal(hoursRaw, &hours)
	if err != nil { return nil, err }

	details := &models.RestaurantDetails{
		ID: *id,
		Name: *name,
		Hours: hours,
	}
	
	return details, nil	
}

func (psql PSQL) CreateRestaurants(restaurants []models.UnsavedRestaurantDetails) error {

	for _, restaurant := range restaurants {
		_, err := psql.DB.Exec(`
				INSERT INTO restaurant(name)
				Values ($1)
				on conflict ("name") do nothing;
			`,
			restaurant.Name,
		)
		if err != nil { return err }

		for _, details := range restaurant.Hours {
			_, err := psql.DB.Exec(`
					with data(name, day, opening, closing) as (
						Values ($1, $2, $3, $4)
					), rest as (
						Select r.restaurant_id from restaurant r, data d where r.name = d.name
					)
					INSERT INTO restaurant_hours(restaurant_id, day_of_week, opening_time, closing_time)
					SELECT r.restaurant_id, cast(d.day as day), d.opening::TIME, d.closing::TIME from rest r, data d;
				`,
				restaurant.Name,
				details.Day,
				details.Opening,
				details.Closing,
			)
			if err != nil { return err }
		}		
	}

	return nil
}
