package http

import (
	"github.com/AnkushJadhav/kamaji-root/pkg/server/http/handlers"
)

func attachExternalRoutes(srv *Server) {
	srv.app.Post("/v1/api/users", handlers.HandleCreateUser(srv.config.StorageDriver))
}

func attachInternalRoutes(srv *Server) {

}
