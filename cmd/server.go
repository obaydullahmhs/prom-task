package cmd

import (
	"fmt"
	"github.com/obaydullahmhs/prom-task/api"

	"github.com/spf13/cobra"
)

var (
	port = 8090
)
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		api.StartServer(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8090, "Default Port For HTTP Server")

	rootCmd.AddCommand(serverCmd)

}
