/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tuxdotrs/trok/internal/client"
)

// clientCmd represents the local command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Initiates a local proxy to the remote server",
	Long:  "Initiates a local proxy to the remote server",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, err := cmd.Flags().GetString("serverAddr")
		if err != nil {
			panic(err)
		}

		localAddr, err := cmd.Flags().GetString("localAddr")
		if err != nil {
			panic(err)
		}

		client.Start(serverAddr, localAddr)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringP("serverAddr", "s", "trok.tux.rs:1337", "Remote server address")
	clientCmd.Flags().StringP("localAddr", "a", "0.0.0.0:80", "Local addr to expose")
}
