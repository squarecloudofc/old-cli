package cmd

import (
	"github.com/richaardev/squarecloud-cli/internal/api"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [appid]",
	Short: "delete your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		confirm := false
		survey.AskOne(&survey.Confirm{
			Message: "Are you sure you want to delete your application? (This is an IRREVERSIBLE action)",
		}, &confirm)

		if confirm {
			resp, err := api.ApplicationDelete(args[0])
			if !resp && err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Println("Your application has been deleted.")
		} else {
			cmd.Println("Action canceled! Glad you chose to continue using Square!")
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
