package repositories

import (
	"net/http"

	"github.com/MarselBissengaliyev/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
)

// Create book handler
func (r *Repository) CreateBook(context *fiber.Ctx) (err error) {
	book := models.Book{}

	if err = context.BodyParser(&book); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "request failed",
		})
	}

	if err = r.DB.Create(&book).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "create failed",
		})
	}

	context.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"data":    book,
		"message": "book created",
	})

	return nil
}
