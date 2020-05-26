package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/modules/users"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"

	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"
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
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}

		response := ResponseBody{
			Data: users,
		}
		c.Status(http.StatusOK).JSON(response)
		return
	}
}

// HandleGetUserByID handles the request to get a user by id
func HandleGetUserByID(str store.Driver) func(*fiber.Ctx) {
	type ResponseBody struct {
		Data interface{} `json:"data"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if c.Params("id") == "" {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		user, err := users.GetUserByID(ctx, str, c.Params("id"))
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}
		if user == nil {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		response := ResponseBody{
			Data: user,
		}
		c.Status(http.StatusOK).JSON(response)
		return
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
			return
		}

		user, err := users.CreateUser(ctx, str, request.Email, request.RoleID)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}

		response := ResponseBody{
			Data: user,
		}
		c.Status(http.StatusOK).JSON(response)
		return
	}
}

// HandleDeleteUser handles the request to delete a user
func HandleDeleteUser(str store.Driver) func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if c.Params("id") == "" {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		//TODO: Validate if user with id exists

		if err := users.DeleteUser(ctx, str, c.Params("id")); err != nil {
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}

		c.Status(http.StatusOK).Send()
		return
	}
}

// HandleRegisterUser handles the registration of a created user
func HandleRegisterUser(str store.Driver) func(*fiber.Ctx) {
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if c.Params("id") == "" {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		request := RequestBody{}
		if err := c.BodyParser(&request); err != nil {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		if err := users.RegisterUser(ctx, str, c.Params("id"), request.Username, request.Password); err != nil {
			if _, isPPError := err.(*users.PasswordDoesNotMatchPolicy); isPPError {
				c.Status(http.StatusUnprocessableEntity).Send(err.Error())
				return
			}
			c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			return
		}

		c.Status(http.StatusOK).Send()
		return
	}
}

// HandleLoginUser handles the registration of a created user
func HandleLoginUser(str store.Driver) func(*fiber.Ctx) {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *fiber.Ctx) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		request := RequestBody{}
		if err := c.BodyParser(&request); err != nil {
			c.Status(http.StatusBadRequest).Send()
			return
		}

		if err := users.AuthenticateUser(ctx, str, request.Email, request.Password); err != nil {
			_, invUser := err.(*users.UserDoesNotExist)
			_, invCred := err.(*users.AuthCredentialMismatch)
			if invUser || invCred {
				c.Status(http.StatusUnauthorized).Send("Invalid Email or Password")
			} else {
				c.Status(http.StatusInternalServerError).JSON(Handle500InternalServerError(requestid.Get(c), err))
			}
		}
		return
	}
}
