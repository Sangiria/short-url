package database

import (
	"short_url/enviroment"
	"short_url/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetDBConnection() *gorm.DB {
	if dbConnection == nil {
		connectDB()

	}
	return dbConnection
}

func connectDB() *gorm.DB {
	dburl := enviroment.GoDotEnvVariable("DATABASE_URL")


	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&models.Shorten{})

	//все переменные с _ игнорируются, тут переменная err нахуй не нужна так что игнорим

	dbConnection = db

	return dbConnection
}