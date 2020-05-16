package main

import (
	"os"

	"github.com/AnkushJadhav/kamaji-root/cmd"
	"github.com/AnkushJadhav/kamaji-root/logger"
)

const (
	exitFail = 1
)

func main() {
	logger.InitLogger()
	if err := cmd.RootCmd.Execute(); err != nil {
		logger.Errorf("Exiting application due to error : %v", err.Error())
		os.Exit(exitFail)
	}
}
