package commons

import (
	"os"
	"strconv"
	"time"

	"github.com/devNica/mochileros/exceptions"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(UserId string, profiles []map[string]interface{}) string {

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	exceptions.PanicLogging(err)

	claims := jwt.MapClaims{
		"UserId":   UserId,
		"profiles": profiles,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exceptions.PanicLogging(err)

	return tokenSigned

}
