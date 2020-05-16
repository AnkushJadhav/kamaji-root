package http

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

// Server is the default HTTP server for the kamaji-root application
type Server struct {
	app      *fiber.App
	settings *fiber.Settings
}

// Bootstrap initialises the http server without starting it
func (srv *Server) Bootstrap() error {
	srv.initServerSettings()
	srv.prepopulatePool(false)

	srv.initServer()

	attachExternalRoutes(srv.app)
	attachInternalRoutes(srv.app)

	return nil
}

// Start runs the default HTTP server
func (srv *Server) Start(bindIP string, port int) error {
	srv.app.Listen(bindIP + ":" + strconv.Itoa(port))
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
