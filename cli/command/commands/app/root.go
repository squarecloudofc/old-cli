package app

import (
	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/urfave/cli/v2"
)

func subcommands(scli *squarecli.SquareCloudCli) []*cli.Command {
	return []*cli.Command{
		NewStatusCommand(scli),
	}
}

func NewAppCommand(scli *squarecli.SquareCloudCli) *cli.Command {
	cmd := &cli.Command{
		Name:        "app",
		Usage:       "Manage your Square Cloud applications",
		Subcommands: subcommands(scli),
		Action: func(ctx *cli.Context) error {
			return runAppCommand(scli, ctx)
		},
	}

	return cmd
}

func runAppCommand(scli *squarecli.SquareCloudCli, ctx *cli.Context) error {
	return cli.ShowSubcommandHelp(ctx)
}
