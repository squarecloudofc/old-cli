package commands

import (
	"fmt"

	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/squarecloudofc/cli/cli/auth"
	"github.com/squarecloudofc/cli/cli/config"
	"github.com/urfave/cli/v2"
)

func NewLoginCommand(scli *squarecli.SquareCloudCli) *cli.Command {

	cmd := &cli.Command{
		Name:  "login",
		Usage: "Login to Square Cloud",
		Action: func(ctx *cli.Context) error {
			return runLoginCommand(scli, ctx)
		},
	}

	return cmd
}

func runLoginCommand(scli *squarecli.SquareCloudCli, ctx *cli.Context) error {
	token, err := auth.AuthFlow(auth.TokenType)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	config.Update(&config.Config{
		Identity: config.Identity{
			IdentityToken: token,
		},
	})

	return nil
}
