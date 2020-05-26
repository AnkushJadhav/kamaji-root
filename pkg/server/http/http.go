package http

import (
	"strconv"

	"github.com/AnkushJadhav/kamaji-root/pkg/server"

	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"
)

// Server is the default HTTP server for the kamaji-root application
type Server struct {
	app      *fiber.App
	settings *fiber.Settings
	config   *server.Config
}

// Bootstrap initialises the http server without starting it
func (srv *Server) Bootstrap(conf *server.Config) error {
	srv.config = conf
	srv.initServerSettings()
	srv.prepopulatePool(conf.PopulatePool)

	srv.initServer()
	srv.app.Use(requestid.New())

	loadUnrestrictedRootRoutes(srv)
	loadUnrestrictedNodeRoutes(srv)

	if err := srv.EnableJWTAuthentication(); err != nil {
		return err
	}

	loadRestrictedRootRoutes(srv)
	loadRestrictedNodeRoutes(srv)

	return nil
}

// Start runs the default HTTP server
func (srv *Server) Start() error {
	return srv.app.Listen(srv.config.BindIP + ":" + strconv.Itoa(srv.config.Port))
}

// Stop stops the default HTTP server
func (srv *Server) Stop() error {
	if err := srv.app.Shutdown(); err != nil {
		return err
	}
	if err := srv.config.StorageDriver.Disconnect(); err != nil {
		return err
	}
	return nil
}
