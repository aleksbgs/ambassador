package main

import (
	"context"
	"encoding/json"
	"github.com/aleksbgs/ambassador/src/database"
	"github.com/aleksbgs/ambassador/src/models"
	"github.com/aleksbgs/ambassador/src/services"
	"github.com/go-redis/redis/v8"
)

func main() {
	database.Connect()
	database.SetupRedis()

	ctx := context.Background()

	response, err := services.UserService.Get("users", "")
	if err != nil {
		panic(err)
	}

	var users []models.User

	json.NewDecoder(response.Body).Decode(&users)

	for _, user := range users {
		if user.IsAmbassador {
			ambassador := models.Ambassador(user)
			ambassador.CalculateRevenue(database.DB)

			database.Cache.ZAdd(ctx, "rankings", &redis.Z{
				Score:  *ambassador.Revenue,
				Member: user.Name(),
			})
		}

	}
}
