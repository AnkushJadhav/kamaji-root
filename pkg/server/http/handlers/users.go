package handlers

import (
	"net/http"

	"github.com/AnkushJadhav/kamaji-root/logger"
	"github.com/AnkushJadhav/kamaji-root/pkg/users"
	"github.com/AnkushJadhav/kamaji-root/store"
	"github.com/gofiber/fiber"
)

// HandleGetAllUsers handles the request to get all users
func HandleGetAllUsers(str store.Store) func(*fiber.Ctx) {
	type ResponseGetAllUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Role  int    `json:"role"`
	}
	return func(ctx *fiber.Ctx) {
		out, err := users.GetAllUsers(str)
		if err != nil {
			logger.Errorln(err)
			ctx.Status(http.StatusInternalServerError).Send("Oops! Something went wrong.")
		}

		resp := make([]ResponseGetAllUser, 0)
		for _, u := range out {
			resp = append(resp, ResponseGetAllUser{
				ID:    u.ID,
				Email: u.Email,
				Role:  u.Role,
			})
		}
		ctx.Status(http.StatusOK).JSON(resp)
	}
}

// HandleCreateUser handles the request to craete a user
func HandleCreateUser(str store.Store) func(*fiber.Ctx) {
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
		body := new(RequestCreateUser)
		if err := ctx.BodyParser(body); err != nil {
			logger.Errorln(err)
			ctx.Status(http.StatusInternalServerError).Send("Oops! Something went wrong.")
		}

		inp := users.CreateUserInput{
			Email: body.Email,
			Role:  body.Role,
		}
		out, err := users.CreateUser(str, inp)
		if err != nil {
			logger.Errorln(err)
			ctx.Status(http.StatusInternalServerError).Send("Oops! Something went wrong.")
		}

		resp := ResponseCreateUser{
			ID:    out.ID,
			Email: out.Email,
			Role:  out.Role,
		}
		ctx.Status(http.StatusOK).JSON(resp)
	}
}
