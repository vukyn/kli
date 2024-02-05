package file

import (
	"kli/constants"
	"kli/internal/file/actions"
	"kli/internal/file/flags"

	"github.com/urfave/cli/v2"
)

func NewCommands() []*cli.Command {
	action := actions.NewAction()
	commands := []*cli.Command{
		{
			Name:    constants.FILE_COMMAND_NAME,
			Aliases: []string{constants.FILE_COMMAND_ALIAS},
			Usage:   constants.FILE_COMMAND_USAGE,
			Subcommands: []*cli.Command{
				{
					Name:    constants.FILE_COMMAND_MAKE_NAME,
					Aliases: []string{constants.FILE_COMMAND_MAKE_ALIAS},
					Usage:   constants.FILE_COMMAND_MAKE_USAGE,
					Action:  action.Execute,
					Flags:   flags.CreateFlags(),
				},
				{
					Name:    constants.FILE_COMMAND_RENAME_NAME,
					Aliases: []string{constants.FILE_COMMAND_RENAME_ALIAS},
					Usage:   constants.FILE_COMMAND_RENAME_USAGE,
					Action:  action.Execute,
					Flags:   flags.RenameFlags(),
				},
				{
					Name:    constants.FILE_COMMAND_DELETE_NAME,
					Aliases: []string{constants.FILE_COMMAND_DELETE_ALIAS},
					Usage:   constants.FILE_COMMAND_DELETE_USAGE,
					Action:  action.Execute,
					Flags:   flags.DeleteFlags(),
				},
			},
		},
	}
	return commands
}
