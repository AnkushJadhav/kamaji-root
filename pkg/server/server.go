package server

// Server is the interface to be fulfilled by any web server running on this application
type Server interface {
	Bootstrap() error
	Start(string, int) error
	Stop() error
}