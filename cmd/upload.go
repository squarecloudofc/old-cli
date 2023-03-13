package cmd

import (
	"os"
	"squarecloud/internal/api"
	"squarecloud/internal/util"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"up"},
	Short:   "upload your new application to Square Cloud",
	Run: func(cmd *cobra.Command, args []string) {

		workDir, _ := os.Getwd()

		tempDir, _ := util.GetTempDir()
		file, err := os.CreateTemp(tempDir, "archive*.zip")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
		defer os.Remove(file.Name())
		defer file.Close()

		err = util.ZipFolder(workDir, file)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		success, err := api.ApplicationUpload(file)
		if !success && err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("Upload successful, Thank you for choosing Square Cloud!")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
