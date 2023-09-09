package routes

import (
	"github.com/MarselBissengaliyev/go-fiber-postgres/repositories"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookStoreRoutes(rg fiber.Router) {
	r := repositories.Repository{}

	routes := rg.Group("/books")
	routes.Get("/", r.GetBooks)
	routes.Get("/:id", r.GetBookById)
	routes.Post("/", r.CreateBook)
	routes.Delete("/:id", r.DeleteBookById)
}
