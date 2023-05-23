package main

import (
	micro "go_web_server/app"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go_web_server/migrations"
)

func main() {
	//connect with db
	dsn := "host=localhost user=andrew password=TiYx9a395%k^ dbname=web_server_go port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//open config
	dbConfig, err := db.DB()
	if err != nil {
		panic(err)
	}

	defer dbConfig.Close()

	//config pool connection
	dbConfig.SetMaxIdleConns(3)
	dbConfig.SetMaxOpenConns(3)

	//run function that automigrate models
	migrations.RunMigrations(db)

	//new app (fiber object)
	app := fiber.New()

	api := app.Group("/api")

	//setup all router that are in app
	micro.SetupRouters(api, db)

	app.Listen(":1024")
}
