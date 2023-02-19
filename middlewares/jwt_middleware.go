package middlewares

import (
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/models"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticateJWT(role string, config configurations.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			profiles := claims["profiles"].([]interface{})

			for _, profileInterface := range profiles {
				roleMap := profileInterface.(map[string]interface{})
				if roleMap["Role"] == role {
					return ctx.Next()
				}
			}

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(models.GeneralResponseModel{
					Code:    401,
					Message: "Unathorized",
					Data:    "Invalid Role",
				})
		},

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(models.GeneralResponseModel{
						Code:    400,
						Message: "Bad Requets",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusForbidden).
					JSON(models.GeneralResponseModel{
						Code:    403,
						Message: "Forbidden",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})
}
