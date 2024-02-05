package actions

import (
	"fmt"
	"kli/constants"
	"kli/internal/command"
	"os"

	"github.com/urfave/cli/v2"
)

type Action struct {
}

func NewAction() command.Command {
	return &Action{}
}

func (a *Action) Execute(ctx *cli.Context) error {
	switch ctx.Command.Name {
	case constants.FOLDER_COMMAND_MAKE_NAME:
		return a.create(ctx)
	case constants.FOLDER_COMMAND_RENAME_NAME:
		return a.rename(ctx)
	case constants.FOLDER_COMMAND_DELETE_NAME:
		return a.delete(ctx)
	}
	return nil
}

func (a *Action) Undo(ctx *cli.Context) error {
	// switch ctx.Command.Name {
	// case constants.FOLDER_COMMAND_MAKE_NAME:
	// 	return a.delete(ctx)
	// case constants.FOLDER_COMMAND_RENAME_NAME:
	// 	return a.rename(ctx)
	// case constants.FOLDER_COMMAND_DELETE_NAME:
	// 	return a.create(ctx)
	// }
	return nil
}

func (a *Action) create(ctx *cli.Context) error {
	foldername := ctx.String(constants.FOLDER_COMMAND_FOLDER_FLAG_NAME)
	fmt.Println("Creating folder:", foldername)

	if _, err := os.Stat(foldername); err != nil {
		if err := os.MkdirAll(foldername, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (a *Action) rename(ctx *cli.Context) error {
	oldname := ctx.String(constants.FOLDER_COMMAND_SRC_FLAG_NAME)
	newname := ctx.String(constants.FOLDER_COMMAND_DST_FLAG_NAME)
	fmt.Println("Renaming folder:", oldname, "to", newname)

	if err := os.Rename(oldname, newname); err != nil {
		return err
	}
	return nil
}

func (a *Action) delete(ctx *cli.Context) error {
	foldername := ctx.String(constants.FOLDER_COMMAND_FOLDER_FLAG_NAME)
	fmt.Println("Deleting folder:", foldername)

	if err := os.Remove(foldername); err != nil {
		return err
	}
	return nil
}
