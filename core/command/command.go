package command

import "github.com/urfave/cli/v2"

type Command interface {
	Execute(ctx *cli.Context) error
	Undo(ctx *cli.Context) error
}
