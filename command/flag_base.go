package command

import (
	"flag"
	"fmt"

	"github.com/binarysoupdev/got-style/style"
)

type FlagCommandBase struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
}

func NewFlagCommandBase(name, description string) FlagCommandBase {
	return FlagCommandBase{
		Name:        name,
		Description: description,
	}
}

// Command interface implementation. Returns the Name field.
func (cmd FlagCommandBase) GetName() string {
	return cmd.Name
}

// Command interface implementation. Prints usage using the command's Name and Description fields.
func (cmd FlagCommandBase) PrintUsage() {
	fmt.Printf("%s\t| %s\n", style.BoldInfo.Sprint(cmd.Name), cmd.Description)
}

func (cmd *FlagCommandBase) Initialize() {
	cmd.Flags = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.Flags.Usage = func() {
		cmd.PrintUsage()
		style.New(style.MAGENTA).Println("Options:")
		cmd.Flags.PrintDefaults()
	}
}
