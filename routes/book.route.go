package routes

import (
	"github.com/MarselBissengaliyev/go-fiber-postgres/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterBookStoreRoutes(rg fiber.Router, db *gorm.DB) {
	r := repositories.Repository{
		DB: db,
	}

	routes := rg.Group("/books")
	routes.Get("/", r.GetBooks)
	routes.Get("/:id", r.GetBookById)
	routes.Post("/", r.CreateBook)
	routes.Delete("/:id", r.DeleteBookById)
}
