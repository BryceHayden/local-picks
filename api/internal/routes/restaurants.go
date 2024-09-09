package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	// "net/http"

	"github.com/brycehayden/resume/internal/models"
	"github.com/labstack/echo/v4"
)

func (this Router) searchRestaurant(c echo.Context) error {
	json_data := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&json_data)
	if err != nil { return err }

	day := json_data["day"]
	time := json_data["time"]


	fmt.Printf("\nSearch Params: %v %v\n", day, time)
	var restaurants []models.Restaurant
	if day == "" && time == "" {
		restaurants, err = this.db.DB.FetchRestaurants()
		if err != nil { return err }
	} else if day != "" && time == "" {
		restaurants, err = this.db.DB.FetchRestaurantsByDay(day)
		if err != nil { return err }
	} else { 	
		restaurants, err = this.db.DB.FetchRestaurantsByDateTime(day, time)
		if err != nil { return err }
	}
	if err != nil { return err }

	return c.JSON(http.StatusOK, echo.Map{
		"restaurants": restaurants,
	})
}

func (this Router) getRestaurants(c echo.Context) error {
	restaurants, err := this.db.DB.FetchRestaurants()
	if err != nil { return err }
	
	return c.JSON(http.StatusOK, echo.Map{
		"restaurants": restaurants,
	})
}


func (this Router) getRestaurantDetails(c echo.Context) error {
	id := c.Param("id")
	restaurant, err := this.db.DB.FetchRestaurantDetails(id)
	
	if err != nil { return err }
	
	return c.JSON(http.StatusOK, echo.Map{
		"restaurant": restaurant,
	})
}

func adjust (hour string, minutes string, indicator string) string {
	adj := 0				
	if indicator == "pm" { adj = 12 }
	intHour, _ := strconv.Atoi(hour)
	return strconv.Itoa(intHour  + adj) + ":" + minutes

}

func militaryTime (times string) (string, string) {
	var time string
	remaining := "" 
	i := 0 
	out: 
	for ; i < len(times); i++ {
		switch current := times[i]; current {
			case ':':
				time = adjust(times[:i], string(times[i+1:i + 3]), string(times[i+4:i + 6]))
				break out;
				
			case ' ':
				time = adjust(times[:i], "00", string(times[i+1:i + 3]))
				break out;
				
		}	
	}

	for j := i; j < len(times); j++ {
		if times[j] == '-' {
			remaining = times[j + 2:]
		}
	}

	return time, remaining
}

func parseTime (times string) (string, string){	
	opening, remaining :=  militaryTime(times)
	closing, _  :=  militaryTime(remaining)

	return opening, closing
}

func parseDays (hours string, days []string) ([]string, string) {
	if hours[3] == '-' {
		order := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun" }
		first := hours[:3]
		last := hours[4:7]
		addDay := false; 

		for i := 0; i < len(order); i++ {
			if order[i] == first {
				addDay = true
			}

			if addDay {
				days = append(days, order[i])
			}

			if order[i] == last {
				break;
			}
		}

		if hours[7] == ',' {
			return parseDays(hours[9:], days)
		}

		return days, hours[8:]
	} 
		
	days = append(days, hours[0:3])	
	if hours[3] == ',' {
		return parseDays(hours[5:], days)
	}

	return days, hours[4:]
}

func (this Router) createRestaurants(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil { return err }

	src, err := file.Open()
	if err != nil { return err }
	defer src.Close()
	
	byteValue, err := io.ReadAll(src)
	if err != nil { return err }

	var data []models.RestaurantData
	err = json.Unmarshal(byteValue, &data)
	if err != nil { return err }

	restaurants := []models.UnsavedRestaurantDetails{}
	for _, info := range data {
		restaurant := models.UnsavedRestaurantDetails{}
		restaurant.Name = info.Name

		for _, hours := range info.Times {
			days, times := parseDays(hours, []string{})
			opening, closing := parseTime(times)
			
			for _, day := range days {
				temp := models.RestaurantHours{ Day: day, Opening: opening, Closing: closing}	
				restaurant.Hours = append(restaurant.Hours, temp)
			}
		}

		restaurants = append(restaurants, restaurant)
	}

	this.db.DB.CreateRestaurants(restaurants)
	return c.JSON(http.StatusOK,  echo.Map{ "create": 12 })
}