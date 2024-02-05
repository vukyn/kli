package main

import (
	"fmt"
	"kli/internal/file"
	"kli/internal/folder"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = KLI_VERSION
	app.Commands = make([]*cli.Command, 0)

	// Add file commands
	app.Commands = append(app.Commands, file.NewCommands()...)
	
	// Add folder commands
	app.Commands = append(app.Commands, folder.NewCommands()...)

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
