package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/goforbroke1006/go-msvc-tpl/internal/config"
)

var (
	apiCmd       = &cobra.Command{Use: "api"}
	daemonCmd    = &cobra.Command{Use: "daemon"}
	messagingCmd = &cobra.Command{Use: "messaging"}
	utilCmd      = &cobra.Command{Use: "util"}
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "msvc-tpl",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Defaults()
			if err := config.LoadFromEnvFile(); err != nil {
				return err
			}

			if err := config.LoadFromVault(); err != nil {
				return err
			}

			return nil
		},
	}
	rootCmd.Version = "v0.0.0-unreleased"

	rootCmd.AddCommand(apiCmd, daemonCmd, messagingCmd, utilCmd)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
