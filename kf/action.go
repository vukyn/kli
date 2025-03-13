package main

import (
	"fmt"
	"kli/constant"
	"kli/core/command"
	"kli/core/file"
	"strings"

	"errors"

	"github.com/urfave/cli/v2"
)

type Action struct {
}

func NewAction() command.Command {
	return &Action{}
}

func (a *Action) Execute(ctx *cli.Context) error {
	switch ctx.Command.Name {
	case constant.FILE_COMMAND_DEFAULT:
		switch ctx.Args().Len() {
		case 0:
			fmt.Println("No arguments provided")
			return nil
		case 1:
			return a.create(ctx)
		case 2:
			return a.rename(ctx)
		default:
			return fmt.Errorf("invalid number of arguments (1: create, 2: rename)")
		}
	case constant.FILE_COMMAND_CREATE_NAME:
		return a.create(ctx)
	case constant.FILE_COMMAND_RENAME_NAME:
		return a.rename(ctx)
	case constant.FILE_COMMAND_REMOVE_NAME:
		return a.remove(ctx)
	}
	return nil
}

func (a *Action) Undo(ctx *cli.Context) error {
	return errors.New("Undo not implemented")
}

func (a *Action) create(ctx *cli.Context) error {
	fileNames := make([]string, 0, ctx.Args().Len())
	for i := range ctx.Args().Len() {
		fileNames = append(fileNames, ctx.Args().Get(i))
	}
	fmt.Printf("Creating file(s): %v\n", strings.Join(fileNames, ", "))

	err := file.Create(fileNames...)
	if err != nil {
		return err
	}
	return nil
}

func (a *Action) rename(ctx *cli.Context) error {
	oldName, newName := ctx.Args().Get(0), ctx.Args().Get(1)
	fmt.Printf("Renaming file: %s to %s\n", oldName, newName)

	err := file.Rename(oldName, newName)
	if err != nil {
		return err
	}
	return nil
}

func (a *Action) remove(ctx *cli.Context) error {
	fileNames := make([]string, 0, ctx.Args().Len())
	for i := range ctx.Args().Len() {
		fileNames = append(fileNames, ctx.Args().Get(i))
	}
	fmt.Printf("Removing file: %v\n", strings.Join(fileNames, ", "))

	err := file.Remove(fileNames...)
	if err != nil {
		return err
	}
	return nil
}
