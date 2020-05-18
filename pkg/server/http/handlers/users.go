package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/modules/users"

	"github.com/AnkushJadhav/kamaji-root/pkg/store"

	"github.com/gofiber/fiber"
)

// HandleGetAllUsers handles the request to get all users
func HandleGetAllUsers(str store.Driver) func(*fiber.Ctx) {
	type ResponseBody struct {
		Data interface{} `json:"data"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		users, err := users.GetAllUsers(ctx, str)
		if err != nil {
			c.Status(http.StatusInternalServerError).Send("Oops! Something went wrong!")
		}

		response := ResponseBody{
			Data: users,
		}
		c.Status(http.StatusOK).JSON(response)
	}
}

// HandleCreateUser handles the request to craete a user
func HandleCreateUser(str store.Driver) func(*fiber.Ctx) {
	type RequestCreateUser struct {
		Email string `json:"email"`
		Role  int    `json:"role"`
	}
	type ResponseCreateUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Role  int    `json:"role"`
	}
	return func(ctx *fiber.Ctx) {
		ctx.Status(http.StatusOK)
	}
}
