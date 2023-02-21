package cmd

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
	"squarecloud/internal/api"
	"squarecloud/internal/appconfig"
	"squarecloud/pkg/properties"
	"strconv"
)

var createAppCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize squarecloud config",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := api.GetSelfUser()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		answers := struct {
			AppName     string
			Description string
			Type        string
			MainFile    string
			Memory      string
			Version     string
		}{}

		questions := []*survey.Question{
			{
				Name: "appname",
				Prompt: &survey.Input{
					Message: "Which is your application name?",
				},
			},
			{
				Name: "description",
				Prompt: &survey.Input{
					Message: "Which is your application description?",
				},
			},
			{
				Name: "mainfile",
				Prompt: &survey.Input{
					Message: "Which is your application main file?",
				},
			},
			{
				Name: "memory",
				Prompt: &survey.Input{
					Message: fmt.Sprintf("How much memory do you want to use? (%dmb available)", resp.User.Plan.Memory.Available),
				},
				Validate: func(ans interface{}) error {
					num, err := strconv.Atoi(ans.(string))
					if err != nil {
						return errors.New(ans.(string) + " is not a valid number")
					}

					if num > resp.User.Plan.Memory.Available {
						return errors.New(fmt.Sprintf("%dmb exceeds the limit you have available", num))
					}

					return nil
				},
			},
			{
				Name: "version",
				Prompt: &survey.Select{
					Message: "Which version do want to use in your application?",
					Options: []string{"recommended", "latest"},
					Default: "recommended",
				},
			},
		}

		err = survey.Ask(questions, &answers)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		file, err := os.Create("squarecloud.app")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		config := appconfig.AppConfig{
			DisplayName: answers.AppName,
			Description: answers.Description,
			Main:        answers.MainFile,
			Memory:      answers.Memory,
			Version:     answers.Version,
		}

		toWrite, err := properties.Marshal(config)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		_, err = file.Write(toWrite)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("Configuration file created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(createAppCmd)
}
