package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnkushJadhav/kamaji-root/pkg/store"

	"github.com/AnkushJadhav/kamaji-root/pkg/store/drivers/mongo"

	"github.com/AnkushJadhav/kamaji-root/logger"
	"github.com/AnkushJadhav/kamaji-root/pkg/server"
	"github.com/AnkushJadhav/kamaji-root/pkg/server/http"
)

type app struct {
	isProd     bool
	httpServer *http.Server
}

var mainApp *app

// Start starts the kamaji-root server
func Start(cfgFile string) error {
	if mainApp != nil {
		return fmt.Errorf("kamaji-root application has already started")
	}

	logger.Infoln("staring applciation")
	conf, err := getConfig(cfgFile)
	if err != nil {
		return err
	}
	logger.Infoln("config loaded successfully")

	var logFile = conf.Server.LogFile
	if fileLoggingEnabled(logFile) {
		startFileLogging(logFile)
	}

	logger.Infoln("initialising db driver")
	db, err := getStorageDriver(conf.Mongo.ConnString)
	if err != nil {
		return err
	}

	if err = db.Connect(); err != nil {
		return err
	}

	logger.Infoln("starting http server")
	httpServer := &http.Server{}
	serverConfig := &server.Config{
		EnableTLS:     false,
		PopulatePool:  true,
		BindIP:        conf.Server.BindIP,
		Port:          conf.Server.Port,
		StorageDriver: db,
	}
	mainApp = &app{
		httpServer: httpServer,
	}

	startServer(httpServer, serverConfig)

	return nil
}

// Stop gracefully stops the kamaji-root application
func Stop() error {
	if err := mainApp.httpServer.Stop(); err != nil {
		return err
	}
	return nil
}

func fileLoggingEnabled(logFile string) bool {
	return strings.TrimSpace(logFile) != ""
}

func startFileLogging(logFile string) error {
	f, err := os.Open(logFile)
	if err != nil {
		return err
	}
	defer f.Close()

	logger.SetOutput(f)
	return nil
}

func getStorageDriver(conn string) (store.Driver, error) {
	drvr, err := mongo.NewDriver(conn)
	if err != nil {
		return nil, err
	}

	return drvr, nil
}

func startServer(srv server.Server, conf *server.Config) error {
	if err := srv.Bootstrap(conf); err != nil {
		return err
	}

	if err := srv.Start(); err != nil {
		return err
	}

	return nil
}
