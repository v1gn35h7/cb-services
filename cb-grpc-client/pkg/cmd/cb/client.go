package cb

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v1gn35h7/cb-grpc-client/internal/config"
	"github.com/v1gn35h7/cb-grpc-client/pkg/cmd/cli"
	"github.com/v1gn35h7/cb-grpc-client/pkg/logging"
)

var (
	configPath string
	verbose    bool
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "cb-grpc-client",
		Short: "Cloudbees client service",
		Long:  "Starts Cloudbees grpc client command line tool",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Init read config
			fmt.Println("Config path set to: ", configPath)
			config.ReadConfig(configPath, logging.Logger())
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "conf", "", "config file path")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", true, "verbose mode")

	// Add sub commands
	rootCmd.AddCommand(cli.NewTestCommand())
	rootCmd.AddCommand(cli.NewBookCommand())
	rootCmd.AddCommand(cli.NewModifyCommand())
	rootCmd.AddCommand(cli.NewRemoveCommand())
	rootCmd.AddCommand(cli.NewReceiptCommand())
	rootCmd.AddCommand(cli.NewViewCommand())

	return rootCmd
}
