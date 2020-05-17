package http

import (
	"github.com/AnkushJadhav/kamaji-root/pkg/server/http/handlers"
)

func attachExternalRoutes(srv *Server) {
	srv.app.Get("/v1/api/users", handlers.HandleGetAllUsers(srv.config.StorageDriver))
	srv.app.Post("/v1/api/users", handlers.HandleCreateUser(srv.config.StorageDriver))
}

func attachInternalRoutes(srv *Server) {

}
