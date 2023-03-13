package cmd

import (
	"squarecloud/internal/api"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [appid]",
	Short: "stop your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		success, err := api.ApplicationStop(args[0])
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
