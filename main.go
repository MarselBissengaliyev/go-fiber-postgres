package main

import (
	"log"

	"github.com/MarselBissengaliyev/go-fiber-postgres/config"
	"github.com/MarselBissengaliyev/go-fiber-postgres/models"
	"github.com/MarselBissengaliyev/go-fiber-postgres/routes"
	"github.com/MarselBissengaliyev/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(&config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	if err = models.MigrateBooks(db); err != nil {
		log.Fatalf("could not migrate db. %s", err.Error())
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
