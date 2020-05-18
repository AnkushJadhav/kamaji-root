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
	type RequestBody struct {
		Email  string `json:"email"`
		RoleID int    `json:"roleID"`
	}
	type ResponseBody struct {
		Data interface{} `json:"data"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		request := RequestBody{}
		if err := c.BodyParser(&request); err != nil {
			c.Status(http.StatusBadRequest).Send()
		}

		user, err := users.CreateUser(ctx, str, request.Email, request.RoleID)
		if err != nil {
			c.Status(http.StatusInternalServerError).Send("Oops! Something went wrong!")
		}

		response := ResponseBody{
			Data: user,
		}
		c.Status(http.StatusOK).JSON(response)
	}
}
