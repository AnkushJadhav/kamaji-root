package cmd

import (
	"fmt"
	"os"

	"github.com/AnkushJadhav/kamaji-root/logger"

	"github.com/AnkushJadhav/kamaji-root/pkg/app"
	"github.com/spf13/cobra"
)

var (
	// ConfigFile is the default config file path
	ConfigFile = ""

	// LogFile is the default log file path
	LogFile = ""
)

func genStartCmd() *cobra.Command {
	var cfgFile string

	startCmd := &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateFlagPaths(cfgFile); err != nil {
				logger.Errorf("Error while reading command line files : %v", err.Error())
				return err
			}

			if err := app.Start(cfgFile); err != nil {
				logger.Errorf("Error while starting app : %v", err.Error())
				return err
			}

			return nil
		},
		SilenceUsage: true,
	}

	startCmd.Flags().StringVarP(&cfgFile, "configfile", "c", ConfigFile, fmt.Sprintf("path of the config file for kamaji-root (defaults to %s)", ConfigFile))
	startCmd.MarkFlagRequired("configfile")

	return startCmd
}

func validateFlagPaths(cfgFile string) error {
	f, err := os.OpenFile(cfgFile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}
