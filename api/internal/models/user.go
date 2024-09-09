package models

type NameOfUser struct {
	First		string 		`db:"first_name"		json:"first"`
	Middle		string 		`db:"middle_name"		json:"middle"`
	Surname		string 		`db:"surname"			json:"surname"`
}

// Age			string 		`db:"age"	 			json:"age"`
type User struct {
	ID   		string 		`db:"member_id" 		json:"id"`
	Username	string		`db:"username" 			json:"username"`
	NameOfUser
	Email		string 		`db:"email" 			json:"email"`
	Joined		string 		`db:"joined_date"		json:"joined_date"`
}

type UserRole struct {
	User
	Role		string 		`db:"role"		 		json:"role"`
}


type UserRolePasshash struct {
	ID   		string 		`db:"member_id" 		json:"id"`
	Username	string		`db:"username" 			json:"username"`
	First		string 		`db:"first_name"		json:"first"`
	Middle		string 		`db:"middle_name"		json:"middle"`
	Surname		string 		`db:"surname"			json:"surname"`
	Email		string 		`db:"email" 			json:"email"`
	Joined		string 		`db:"joined_date"		json:"joined_date"`
	Role		string 		`db:"role"		 		json:"role"`	
	Passhash	string 		`db:"passhash"	 		json:"passhash"`
}