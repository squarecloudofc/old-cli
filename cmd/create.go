package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createAppCmd = &cobra.Command{
	Use: "app",
	Run: func(cmd *cobra.Command, args []string) {
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
				Name: "type",
				Prompt: &survey.Select{
					Message: "Which is your application type?",
					Options: []string{"bot", "site"},
					Default: "bot",
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
					Message: "How much memory do you want to use?",
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

		err := survey.Ask(questions, &answers)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		fmt.Println(answers)
	},
}

func init() {
	//createCmd.AddCommand(createAppCmd)
	//rootCmd.AddCommand(createCmd)
}
