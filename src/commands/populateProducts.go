package main

import (
	"github.com/aleksbgs/ambassador/src/database"
	"github.com/aleksbgs/ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&product)
	}
}
