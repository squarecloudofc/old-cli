package cmd

import (
	"squarecloud/internal/api"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup [appid]",
	Short: "Make a backup of your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		backup, err := api.ApplicationBackup(args[0])
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("Your backup is ready!")
		cmd.Println(backup.DownloadURL)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
