package handlers

import (
	"github.com/AnkushJadhav/kamaji-root/logger"
	"github.com/AnkushJadhav/kamaji-root/pkg/users"
	"github.com/AnkushJadhav/kamaji-root/store"
	"github.com/gofiber/fiber"
)

// HandleCreateUser handles the request to craete a user
func HandleCreateUser(str store.Store) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		i := users.CreateUserInput{
			Email:    "ankush",
			Password: "test",
			Role:     1,
			Username: "blahblah",
		}

		o, err := users.CreateUser(str, i)
		if err != nil {
			logger.Errorln(err)
		}
		logger.Infoln(o)

		ctx.Send("FUCKING SHIT AWESOME!")
	}
}
