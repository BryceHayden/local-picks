package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	// "net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}


func (this Router) Login(c echo.Context) error {
	json_data := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&json_data)
	if err != nil { return err }
	
	email := json_data["email"]
	user, err := this.db.DB.FetchUserRolePasshash(email)
	if err != nil { return err }

	password := json_data["password"]
	err = bcrypt.CompareHashAndPassword([]byte(user.Passhash), []byte(password))
	if err != nil { return err }

	expires := time.Now().Add(15 * time.Minute)
	claims := JWTClaims{
		user.Role,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(expires),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "local-picks",
		},
	}
	
	jwt_secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)	
	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil { return err }
	
	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
		"expires": expires,
		"role": user.Role,
	})
}
