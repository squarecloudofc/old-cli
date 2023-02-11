package cmd

import (
	"os"
	"squarecloud/internal/rest"
	"squarecloud/internal/util"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [appid]",
	Short: "Update your application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tempDir, _ := util.GetTempDir()
		file, err := os.CreateTemp(tempDir, "commitArchive")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())

		workDir, _ := os.Getwd()

		err = util.ZipFolder(workDir, file)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		commit, err := rest.ApplicationCommit(args[0], file)
		if !commit && err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("Commit successful, I hope you like your new application version :)")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
