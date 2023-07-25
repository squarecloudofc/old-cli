package main

import (
	"fmt"
	"os"

	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/squarecloudofc/cli/cli/command"
)

func runSquareCloud(scli *squarecli.SquareCloudCli) error {
	app := command.NewRootCommand(scli)
	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}

func main() {
	scli := squarecli.NewSquareCloudCli()

	if err := runSquareCloud(scli); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
