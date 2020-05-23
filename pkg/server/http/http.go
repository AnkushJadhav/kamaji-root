package http

import (
	"strconv"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/server"

	"github.com/gofiber/fiber"
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

	loadUnrestrictedRootRoutes(srv)
	loadUnrestrictedNodeRoutes(srv)

	loadRestrictedRootRoutes(srv)
	loadRestrictedNodeRoutes(srv)

	return nil
}

// Start runs the default HTTP server
func (srv *Server) Start() error {
	srv.app.Listen(srv.config.BindIP + ":" + strconv.Itoa(srv.config.Port))
	return nil
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

func (srv *Server) initServerSettings() {
	srv.settings = &fiber.Settings{}
}

func (srv *Server) prepopulatePool(b bool) {
	srv.settings.Prefork = b
}

func (srv *Server) setReadTimeout(t time.Duration) {
	srv.settings.ReadTimeout = t
}

func (srv *Server) setWriteTimeout(t time.Duration) {
	srv.settings.WriteTimeout = t
}

func (srv *Server) setMaxConcurrency(c int) {
	srv.settings.Concurrency = c
}

func (srv *Server) initServer() {
	srv.app = fiber.New(srv.settings)
}
