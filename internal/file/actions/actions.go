package actions

import (
	"fmt"
	"kli/internal/command"
	"kli/constants"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

type Action struct {
}

func NewAction() command.Command {
	return &Action{}
}

func (a *Action) Execute(ctx *cli.Context) error {
	switch ctx.Command.Name {
	case constants.FILE_COMMAND_MAKE_NAME:
		return a.create(ctx)
	case constants.FILE_COMMAND_RENAME_NAME:
		return a.rename(ctx)
	case constants.FILE_COMMAND_DELETE_NAME:
		return a.delete(ctx)
	}
	return nil
}

func (a *Action) Undo(ctx *cli.Context) error {
	// switch ctx.Command.Name {
	// case constants.FILE_COMMAND_MAKE_NAME:
	// 	return a.delete(ctx)
	// case constants.FILE_COMMAND_RENAME_NAME:
	// 	return a.rename(ctx)
	// case constants.FILE_COMMAND_DELETE_NAME:
	// 	return a.create(ctx)
	// }
	return nil
}

func (a *Action) create(ctx *cli.Context) error {
	filename := ctx.String(constants.FILE_COMMAND_FILE_FLAG_NAME)
	fmt.Println("Creating file:", filename)

	dir, _ := filepath.Split(filename)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	if err := os.WriteFile(filename, []byte{}, 0644); err != nil {
		return err
	}
	return nil
}

func (a *Action) rename(ctx *cli.Context) error {
	oldname := ctx.String(constants.FILE_COMMAND_SRC_FLAG_NAME)
	newname := ctx.String(constants.FILE_COMMAND_DST_FLAG_NAME)
	fmt.Println("Renaming file:", oldname, "to", newname)

	old, err := os.ReadFile(oldname)
	if err != nil {
		return err
	}
	if err := os.WriteFile(newname, old, 0644); err != nil {
		return err
	}
	if err := os.Remove(oldname); err != nil {
		return err
	}
	return nil
}

func (a *Action) delete(ctx *cli.Context) error {
	filename := ctx.String(constants.FILE_COMMAND_FILE_FLAG_NAME)
	fmt.Println("Deleting file:", filename)

	if err := os.Remove(filename); err != nil {
		return err
	}
	return nil
}
