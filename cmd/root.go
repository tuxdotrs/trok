/*
Copyright © 2024 tux <0xtux@pm.me>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:    "trok",
	Hidden: true,
	Long:   "Simple TCP tunnel in Go that exposes local ports to internet, bypassing NAT firewalls.",
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
