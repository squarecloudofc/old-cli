package main

import (
	"squarecloud/cmd"
	_ "squarecloud/internal/config"
)

func main() {
	cmd.Execute()
}
