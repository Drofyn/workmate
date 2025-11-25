package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/things-go/go-socks5"
)

var listenAddr string

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "start proxy",
	Run: func(cmd *cobra.Command, args []string) {
		server := socks5.NewServer()

		fmt.Println("socks5 listening ", listenAddr)
		if err := server.ListenAndServe("tcp", listenAddr); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	proxyCmd.Flags().StringVarP(&listenAddr, "listen", "l", ":1080", "listen address")
	Cmd.AddCommand(proxyCmd)
}
