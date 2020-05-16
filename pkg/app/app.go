package app

import (
	"os"
	"strings"

	"github.com/AnkushJadhav/kamaji-root/store"

	"github.com/AnkushJadhav/kamaji-root/store/mongo"

	"github.com/AnkushJadhav/kamaji-root/logger"
	"github.com/AnkushJadhav/kamaji-root/pkg/server"
	"github.com/AnkushJadhav/kamaji-root/pkg/server/http"
)

// Start starts the kamaji-root server
func Start(cfgFile string) error {
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
	db, err = db.Connect()
	if err != nil {
		return err
	}

	logger.Infoln("starting http server")
	httpServer := &http.Server{}
	serverConfig := &server.Config{
		IsProd:        false,
		EnableTLS:     false,
		PopulatePool:  true,
		BindIP:        conf.Server.BindIP,
		Port:          conf.Server.Port,
		StorageDriver: db,
	}
	startServer(httpServer, serverConfig)

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

func getStorageDriver(dst string) (store.Store, error) {
	drv, err := mongo.NewMongoDriver(dst)
	if err != nil {
		return mongo.Mongo{}, err
	}

	return drv, nil
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
