package command

import (
	"flag"
	"fmt"

	"github.com/binarysoupdev/got-style/style"
)

// A base type for commands that use a flagset to parse arguments.
type FlagCommandBase struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
}

// Create a new FlagCommandBase from a command name and description.
func NewFlagCommandBase(name, description string) FlagCommandBase {
	return FlagCommandBase{
		Name:        name,
		Description: description,
	}
}

// Return the command's id (ie. its name).
func (cmd FlagCommandBase) GetID() string {
	return cmd.Name
}

// Return the command's usage string (ie. its description).
func (cmd FlagCommandBase) GetUsage() string {
	return cmd.Description
}

// Create a new flagset and override its usage function.
func (cmd *FlagCommandBase) Initialize() {
	cmd.Flags = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.Flags.Usage = func() {
		fmt.Println(cmd.Description)
		style.New(style.MAGENTA).Println("Options:")
		cmd.Flags.PrintDefaults()
	}
}
