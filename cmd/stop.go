package cmd

import (
	"squarecloud/internal/rest"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [appid]",
	Short: "Stop running your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		success, err := rest.ApplicationStop(args[0])
		if !success && err != nil {
			cmd.PrintErrln(err)
			return
		}
		cmd.Println("Your application has been stopped")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
