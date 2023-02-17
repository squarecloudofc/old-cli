package cmd

import (
	"squarecloud/internal/api"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [appid]",
	Short: "Start running your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		success, err := api.ApplicationStart(args[0])
		if !success && err != nil {
			cmd.PrintErrln(err)
			return
		}
		cmd.Println("Your application will be started, this can take a few minutes")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
