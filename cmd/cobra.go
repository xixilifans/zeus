package cmd

import (
	"os"
	"zeus/cmd/api"
	"zeus/cmd/auths"
	"zeus/cmd/migrate"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "zeus",
	Short:             "zeus API server",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `Start zeus API server`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.MigrateCmd)
	rootCmd.AddCommand(auths.AuthCmd)
}

//Execute : run commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
