package cmd

import (
	"github.com/richaardev/squarecloud-cli/internal/api"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [appid]",
	Short: "restart your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		success, err := api.ApplicationRestart(args[0])
		if !success && err != nil {
			cmd.PrintErrln(err)
			return
		}
		cmd.Println("Your application will be restarted, this can take a few minutes")
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
