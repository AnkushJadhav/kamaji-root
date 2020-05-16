package http

import (
	"strconv"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/server"

	"github.com/AnkushJadhav/kamaji-root/store"
	"github.com/gofiber/fiber"
)

// Server is the default HTTP server for the kamaji-root application
type Server struct {
	app      *fiber.App
	settings *fiber.Settings
	store    *store.Store
	config   *server.Config
}

// Bootstrap initialises the http server without starting it
func (srv *Server) Bootstrap(conf *server.Config) error {
	srv.initServerSettings()
	srv.prepopulatePool(conf.PopulatePool)

	srv.initServer()
	attachExternalRoutes(srv.app)
	attachInternalRoutes(srv.app)

	srv.config = conf

	return nil
}

// Start runs the default HTTP server
func (srv *Server) Start() error {
	srv.app.Listen(srv.config.BindIP + ":" + strconv.Itoa(srv.config.Port))
	return nil
}

// Stop stops the default HTTP server
func (*Server) Stop() error {
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
