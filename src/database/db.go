package database

import (
	"github.com/aleksbgs/ambassador/src/models"
	"github.com/aleksbgs/ambassador/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	host := utils.ViperEnvVariable("DBDOCKERHOST")
	dbuser := utils.ViperEnvVariable("DBUSER")
	password := utils.ViperEnvVariable("DBPASSWORD")
	dbname := utils.ViperEnvVariable("DBNAME")
    //microservice setup
	 dsn := "host=" + host + " " + "user=" + dbuser + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=5432 sslmode=disable"
    //local setup
	//dsn := "host=" + "localhost" + " " + "user=" + "root" + " " + "password=" + "root" + " " + "dbname=" + "ambassador" + " " + "port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.Product{}, models.Link{}, models.Order{}, models.OrderItem{})
}
