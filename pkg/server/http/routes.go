package http

import (
	"github.com/AnkushJadhav/kamaji-root/pkg/server/http/handlers"
)

func attachExternalRoutes(srv *Server) {
	srv.app.Get("/v1/api/users", handlers.HandleGetAllUsers(srv.config.StorageDriver))
	srv.app.Get("/v1/api/users/:id", handlers.HandleGetUserByID(srv.config.StorageDriver))
	srv.app.Post("/v1/api/users", handlers.HandleCreateUser(srv.config.StorageDriver))
	srv.app.Delete("/v1/api/users/:id", handlers.HandleDeleteUser(srv.config.StorageDriver))
	srv.app.Post("/v1/api/users/register/:id", handlers.HandleRegisterUser(srv.config.StorageDriver))
}

func attachInternalRoutes(srv *Server) {

}
