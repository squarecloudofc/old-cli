package main

import (
	"github.com/richaardev/squarecloud-cli/cmd"
	_ "github.com/richaardev/squarecloud-cli/internal/config"
)

func main() {
	cmd.Execute()
}
