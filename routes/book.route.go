package routes

import "github.com/gofiber/fiber/v2"

func RegisterBookStoreRoutes(rg fiber.Router) {
	routes := rg.Group("/books")
	routes.Get("/")
}
