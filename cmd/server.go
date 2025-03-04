/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tuxdotrs/trok/internal/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Initiates the remote proxy server",
	Long:  "Initiates the remote proxy server",
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			panic(err)
		}
		server.Start(addr)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("addr", "a", "0.0.0.0:1337", "Addr for the server to listen on")
}
