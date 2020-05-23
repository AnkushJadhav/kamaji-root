package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/AnkushJadhav/kamaji-root/logger"
	"github.com/AnkushJadhav/kamaji-root/pkg/modules/nodes"

	"github.com/AnkushJadhav/kamaji-root/pkg/store"
	"github.com/gofiber/fiber"
)

// HandleRegisterNode handles the registration of a created node
func HandleRegisterNode(str store.Driver) func(*fiber.Ctx) {
	type RequestBody struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		HostData struct {
			OS            string `json:"os"`
			DockerVersion string `json:"dockerversion"`
		} `json:"hostdata"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if c.Params("user_id") == "" {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		request := RequestBody{}
		if err := c.BodyParser(&request); err != nil {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		if err := nodes.RegisterNode(ctx, str, c.Params("user_id"), request.Name, request.Version, request.HostData.OS, request.HostData.DockerVersion); err != nil {
			logger.Errorln(err)
			c.Status(http.StatusInternalServerError).Send("Oops! Something went wrong!")
			return
		}

		c.Status(http.StatusOK).Send()
		return
	}
}
