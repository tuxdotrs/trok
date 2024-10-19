/*
Copyright © 2024 tux <0xtux@pm.me>
*/
package cmd

import (
	"github.com/0xtux/trok/internal/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the local command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Initiates a local proxy to the remote server",
	Long:  "Initiates a local proxy to the remote server",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			panic(err)
		}
		client.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().Uint16P("port", "p", 0, "Local port to expose")
	clientCmd.MarkFlagRequired("port")
}
