package sample

import (
	"fmt"

	"github.com/binarysoupdev/go-commando/command"
)

// Sample hello command for printing "Hello" to the console.
type HelloCommand struct {
	command.FlagCommandBase
}

// Create a new HelloCommand.
func NewHelloCommand() *HelloCommand {
	return &HelloCommand{
		FlagCommandBase: command.NewFlagCommandBase("hello", "prints \"Hello {name}\" to the console"),
	}
}

// Run the command. See usage string for details.
func (cmd HelloCommand) Run(args []string) error {
	name := cmd.Flags.String("name", "World", "name to use when saying hello")
	cmd.Flags.Parse(args)

	if *name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	fmt.Printf("Hello %s!\n", *name)
	return nil
}
