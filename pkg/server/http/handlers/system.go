package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/modules/system"

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

// HandleRootTokenValidityCheck checks whther the root token entered is valid
func HandleRootTokenValidityCheck(str store.Driver) func(*fiber.Ctx) {
	type RequestBody struct {
		RootToken string `json:"roottoken"`
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

		if tokenValid {
			c.Status(http.StatusOK).Send()
		} else {
			c.Status(http.StatusUnauthorized).Send()
		}
		return
	}
}
