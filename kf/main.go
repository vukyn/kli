package main

import (
	"fmt"
	"kli/constant"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	APP      *cli.App       // APP is the main application
	COMMANDS []*cli.Command // COMMANDS is a list of all commands
)

const USAGE_TEXT = `kf [command] [arguments]
kf new_file.txt --> creates a new file
kf new_file.txt renamed_file.txt --> renames new_file.txt to renamed_file.txt
`

func init() {
	action := NewAction()
	APP = &cli.App{
		Name:      constant.FILE_COMMAND_DEFAULT,
		Usage:     constant.FILE_COMMAND_USAGE,
		Version:   constant.KF_VERSION,
		Action:    action.Execute,
		Args:      true,
		UsageText: USAGE_TEXT,
	}
	COMMANDS = []*cli.Command{
		{
			Name:    constant.FILE_COMMAND_CREATE_NAME,
			Aliases: []string{constant.FILE_COMMAND_CREATE_ALIAS},
			Usage:   constant.FILE_COMMAND_CREATE_USAGE,
			Action:  action.Execute,
		},
		{
			Name:    constant.FILE_COMMAND_RENAME_NAME,
			Aliases: []string{constant.FILE_COMMAND_RENAME_ALIAS},
			Usage:   constant.FILE_COMMAND_RENAME_USAGE,
			Action:  action.Execute,
		},
		{
			Name:    constant.FILE_COMMAND_REMOVE_NAME,
			Aliases: []string{constant.FILE_COMMAND_REMOVE_ALIAS},
			Usage:   constant.FILE_COMMAND_REMOVE_USAGE,
			Action:  action.Execute,
		},
	}
}

func main() {
	app := APP
	app.Commands = COMMANDS

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
