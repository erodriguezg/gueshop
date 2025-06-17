package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, service *UserService) {
	group := app.Group("/api/users")

	group.Post("/", func(c *fiber.Ctx) error {
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := c.BodyParser(&input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := service.CreateUser(context.Background(), input.Name, input.Email); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	group.Get("/", func(c *fiber.Ctx) error {
		users, err := service.GetAllUsers(context.Background())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(users)
	})
}
