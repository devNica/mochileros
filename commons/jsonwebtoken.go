package commons

import (
	"strconv"
	"time"

	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/exceptions"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(UserId string, profiles []map[string]interface{}, config configurations.Config) string {

	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MIN"))
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
