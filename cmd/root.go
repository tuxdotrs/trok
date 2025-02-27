/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tuxdotrs/trok/internal/config"
)

var rootCmd = &cobra.Command{
	Use:    "trok",
	Hidden: true,
	Long:   "Simple TCP tunnel in Go that exposes local ports to internet, bypassing NAT firewalls.",
}

func init() {
	config.InitLogger()
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
