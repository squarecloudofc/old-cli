package cmd

import (
	"os/exec"
	"runtime"
	"squarecloud/internal/api"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [appid]",
	Short: "get your application logs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		full, _ := cmd.Flags().GetBool("full")
		noopen, _ := cmd.Flags().GetBool("no-open")
		// Todo: handle error
		logs, err := api.ApplicationLogs(args[0], full)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		if full {
			var err error
			var url = logs.Logs
			cmd.Println("Opening your application's logs in your preferred browser, if it does not open, copy the link below and place it in your browser")
			cmd.Println(url)

			if !noopen {
				switch runtime.GOOS {
				case "linux":
					err = exec.Command("xdg-open", url).Start()
				case "windows":
					err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				case "darwin":
					err = exec.Command("open", url).Start()
				}

				if err != nil {
					cmd.PrintErrln(err)
				}
			}
		} else {
			cmd.Println(logs.Logs)
		}
	},
}

func init() {
	logsCmd.Flags().Bool("full", false, "see all your application logs")
	logsCmd.Flags().Bool("no-open", false, "do not open the logs in the browser when running the command")

	rootCmd.AddCommand(logsCmd)
}
