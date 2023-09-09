package repositories

import (
	"errors"
	"net/http"

	"github.com/MarselBissengaliyev/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
)

// Get books handler
func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := []models.Book{}

	if err := r.DB.Find(&bookModels).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "could not get books",
		})

		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"data":    bookModels,
		"message": "books fetched successfully",
	})

	return nil
}

// Create book handler
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := models.Book{}

	if err := context.BodyParser(&book); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "request failed",
		})

		return err
	}

	if err := r.DB.Create(&book).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "create failed",
		})

		return err
	}

	context.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"data":    book,
		"message": "book created",
	})

	return nil
}

// Delete book handler
func (r *Repository) DeleteBookById(context *fiber.Ctx) error {
	bookModel := models.Book{}

	id := context.Params("id")
	if id == "" {
		err := errors.New("id cannot be empty")

		context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "request failed",
		})

		return err
	}

	if err := r.DB.Delete(bookModel, id).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "delete failed",
		})

		return err
	}

	context.Status(http.StatusNoContent).JSON(fiber.Map{
		"status": "success",
		"message": "book deleted succefully",
	})

	return nil
}
