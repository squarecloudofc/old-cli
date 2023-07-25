package command

import (
	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/squarecloudofc/cli/cli/command/commands"
	"github.com/squarecloudofc/cli/cli/command/commands/app"
	"github.com/urfave/cli/v2"
)

func NewRootCommand(scli *squarecli.SquareCloudCli) *cli.App {
	app := &cli.App{
		Name:     "squarecloud",
		Commands: GetCommands(scli),
		Action: func(ctx *cli.Context) error {
			return cli.ShowAppHelp(ctx)
		},
	}

	return app
}

func GetCommands(scli *squarecli.SquareCloudCli) []*cli.Command {

	return []*cli.Command{
		commands.NewLoginCommand(scli),
		commands.NewAppsCommand(scli),
		app.NewAppCommand(scli),

		commands.NewWhoamiCommand(scli),
	}

}
