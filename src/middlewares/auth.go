package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/aleksbgs/ambassador/src/models"
	"github.com/aleksbgs/ambassador/src/services"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func IsAuthenticated(c *fiber.Ctx) error {

	isAmbassador := strings.Contains(c.Path(), "/api/ambassador")

	var scope string

	if isAmbassador {
		scope = "ambassador"
	} else {
		scope = "admin"
	}

	response, err := services.UserService.Get(fmt.Sprintf("user/%s", scope), c.Cookies("jwt", ""))

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

