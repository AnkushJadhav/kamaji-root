package app

import (
	"os"
	"strings"

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

	logger.Infoln("starting http server")
	httpServer := &http.Server{}
	startServer(httpServer, conf.Server.BindIP, conf.Server.Port)

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

func startServer(srv server.Server, bindIP string, port int) error {
	if err := srv.Bootstrap(); err != nil {
		return err
	}

	if err := srv.Start(bindIP, port); err != nil {
		return err
	}

	return nil
}
