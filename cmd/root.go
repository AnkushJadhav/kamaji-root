package cmd

import (
	"github.com/spf13/cobra"
)

const (
	//VERSION is the version of the cli
	VERSION = "dev"
)

// RootCmd is the base command for kamaji-root cli
var RootCmd = &cobra.Command{
	Use: "kamaji-root",
	Version: VERSION,
}

func init() {
	RootCmd.AddCommand(genStartCmd())
}