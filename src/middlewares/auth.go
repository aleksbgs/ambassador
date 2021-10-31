package middlewares

import (
	"encoding/json"
	"github.com/aleksbgs/ambassador/src/models"
	"github.com/aleksbgs/ambassador/src/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

const SecretKey = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func IsAuthenticated(c *fiber.Ctx) error {
	response, err := services.UserService.Get("user", c.Cookies("jwt", ""))

	if err != nil {
		return err
	}
	var user models.User

	json.NewDecoder(response.Body).Decode(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unautheticated",
		})

	}
	c.Context().SetUserValue("user", user)
	return c.Next()
}
//
//func GenerateJWT(id uint, scope string) (string, error) {
//
//}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}
