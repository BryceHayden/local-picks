package models

// type Details []Detail

type Restaurant struct {
	ID   		string 		`db:"restaurant_id" 	json:"id"`
	Name		string 		`db:"name" 				json:"name"`	
}

type RestaurantHours struct {
	Day 		string 		`db:"day" 				json:"day"`	
	Opening 	string 		`db:"opening" 				json:"opening"`	
	Closing 	string 		`db:"closing" 				json:"closing"`	
}

type RestaurantDetails struct {
	ID   		string 					`db:"restaurant_id" 	json:"id"`
	Name		string 					`db:"name" 				json:"name"`	
	Hours		[]RestaurantHours 		`db:"hours"				json:"hours"`	
}

type UnsavedRestaurantDetails struct {
	Name		string 					`db:"name" 				json:"name"`	
	Hours		[]RestaurantHours 		`db:"hours"				json:"hours"`	
}

type RestaurantData struct {
    Name   string 	`json:"name"`
    Times  []string `json:"times"`
}
