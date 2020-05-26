package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"github.com/AnkushJadhav/kamaji-root/pkg/modules/system"
	"github.com/AnkushJadhav/kamaji-root/pkg/modules/users"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"

	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"
)

// HandleGetBootupState handles the request to get the bootup state
func HandleGetBootupState(str store.Driver) func(*fiber.Ctx) {
	type ResponseBody struct {
		Data interface{} `json:"data"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		state, err := system.GetBootupState(ctx, str)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}

		response := ResponseBody{
			Data: state,
		}
		c.Status(http.StatusOK).JSON(response)
		return
	}
}

// HandleCreateRootUser checks whther the root token entered is valid
func HandleCreateRootUser(str store.Driver) func(*fiber.Ctx) {
	type RequestBody struct {
		RootToken string `json:"roottoken"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		request := RequestBody{}
		if err := c.BodyParser(&request); err != nil {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		tokenValid, err := system.IsRootTokenValid(ctx, str, request.RootToken)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}

		if !tokenValid {
			c.Status(http.StatusUnauthorized).Send()
		} else {
			user, err := users.CreateUser(ctx, str, request.Email, models.RoleAdmin)
			if err != nil {
				c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
				return
			}
			if err := users.RegisterUser(ctx, str, user.ID, request.Username, request.Password); err != nil {
				c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
				return
			}
			// mark bootup as completed
			str.SetBootupState(ctx, models.BootupStateCompleted)
			c.Status(http.StatusOK).Send()
		}
		return
	}
}
