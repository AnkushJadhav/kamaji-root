package http

import (
	"github.com/AnkushJadhav/kamaji-root/pkg/server/http/handlers"
)

func loadUnrestrictedRootRoutes(srv *Server) {
	srv.app.Get("/v1/api/config/bootupstate", handlers.HandleGetAllUsers(srv.config.StorageDriver))
	srv.app.Post("/v1/api/config/rootuser", handlers.HandleCreateRootUser(srv.config.StorageDriver))
}

func loadRestrictedRootRoutes(srv *Server) {
	srv.app.Get("/v1/api/users", handlers.HandleGetAllUsers(srv.config.StorageDriver))
	srv.app.Get("/v1/api/users/:id", handlers.HandleGetUserByID(srv.config.StorageDriver))
	srv.app.Post("/v1/api/users", handlers.HandleCreateUser(srv.config.StorageDriver))
	srv.app.Delete("/v1/api/users/:id", handlers.HandleDeleteUser(srv.config.StorageDriver))
}

func loadUnrestrictedNodeRoutes(srv *Server) {
	srv.app.Post("/v1/api/users/register/:id", handlers.HandleRegisterUser(srv.config.StorageDriver))
}

func loadRestrictedNodeRoutes(srv *Server) {
	srv.app.Post("/v1/api/nodes/register/:user_id", handlers.HandleRegisterNode(srv.config.StorageDriver))
}
