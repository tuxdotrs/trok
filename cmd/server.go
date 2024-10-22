/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package cmd

import (
	"github.com/0xtux/trok/internal/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Initiates the remote proxy server",
	Long:  "Initiates the remote proxy server",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			panic(err)
		}
		server.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Uint16P("port", "p", 1421, "Port for the server to listen on")
}
