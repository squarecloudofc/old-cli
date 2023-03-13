package cmd

import (
	"fmt"
	"os"
	"squarecloud/internal/api"
	"squarecloud/internal/appconfig"
	"squarecloud/internal/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [appid]",
	Short: "commit a new version of your application",
	Run: func(cmd *cobra.Command, args []string) {
		appId := getApplicationId(cmd, args)
		if appId == "" {
			cmd.Println("No application ID found, please specify one")
			return
		}

		tempDir, _ := util.GetTempDir()
		file, err := os.CreateTemp(tempDir, "commitArchive")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())

		saveIdInSettings(appId)

		workDir, _ := os.Getwd()
		err = util.ZipFolder(workDir, file)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		commit, err := api.ApplicationCommit(args[0], file)
		if !commit && err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("Commit successful, I hope you like your new application version :)")
	},
}

func getApplicationId(cmd *cobra.Command, args []string) string {
	var appId string
	if len(args) == 0 {
		cfg, err := appconfig.Get()
		if err != nil {
			return ""
		}
		appId = cfg.ID
	} else {
		appId = args[0]
	}

	return appId
}

func saveIdInSettings(id string) {
	config, err := appconfig.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.ID == id {
		return
	}

	confirm := false
	question := []*survey.Question{
		{
			Name: "",
			Prompt: &survey.Confirm{
				Message: "Do you want to save your application ID?",
			},
		},
	}
	survey.Ask(question, &confirm)

	if confirm {
		err = appconfig.Save(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
