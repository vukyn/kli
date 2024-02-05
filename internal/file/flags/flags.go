package flags

import (
	"kli/constants"

	"github.com/urfave/cli/v2"
)

func CreateFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    constants.FILE_COMMAND_FILE_FLAG_NAME,
			Aliases: []string{constants.FILE_COMMAND_FILE_FLAG_ALIAS},
			Value:   constants.FILE_COMMAND_FILE_FLAG_DEFAULT_VALUE,
			Usage:   constants.FILE_COMMAND_FILE_FLAG_USAGE,
		},
	}
}
func RenameFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    constants.FILE_COMMAND_SRC_FLAG_NAME,
			Aliases: []string{constants.FILE_COMMAND_SRC_FLAG_ALIAS},
			Value:   constants.FILE_COMMAND_SRC_FLAG_DEFAULT_VALUE,
			Usage:   constants.FILE_COMMAND_SRC_FLAG_USAGE,
		},
		&cli.StringFlag{
			Name:    constants.FILE_COMMAND_DST_FLAG_NAME,
			Aliases: []string{constants.FILE_COMMAND_DST_FLAG_ALIAS},
			Value:   constants.FILE_COMMAND_DST_FLAG_DEFAULT_VALUE,
			Usage:   constants.FILE_COMMAND_DST_FLAG_USAGE,
		},
	}
}
func DeleteFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    constants.FILE_COMMAND_FILE_FLAG_NAME,
			Aliases: []string{constants.FILE_COMMAND_FILE_FLAG_ALIAS},
			Value:   constants.FILE_COMMAND_FILE_FLAG_DEFAULT_VALUE,
			Usage:   constants.FILE_COMMAND_FILE_FLAG_USAGE,
		},
	}
}
