/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tuxdotrs/trok/internal/client"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp [port]",
	Short: "Start TCP proxy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port := args[0]

		serverAddr, err := cmd.Flags().GetString("server")
		if err != nil {
			panic(err)
		}

		client.Start(serverAddr, fmt.Sprintf(":%s", port))
	},
}

func init() {
	rootCmd.AddCommand(tcpCmd)
	tcpCmd.Flags().StringP("server", "s", "trok.cloud:1337", "Remote server address")
}
