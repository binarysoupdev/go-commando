package command

import (
	"fmt"

	"github.com/binarysoupdev/got-style/style"
)

type Runner struct {
	cmds map[string]Command
	ids  []string
}

// Creates a new runner using the provided commands.
func NewRunner(commands ...Command) Runner {
	cmds := map[string]Command{}
	ids := make([]string, len(commands))

	for i, cmd := range commands {
		cmds[cmd.GetID()] = cmd
		ids[i] = cmd.GetID()
	}

	return Runner{
		cmds: cmds,
		ids:  ids,
	}
}

// Run the command that matches the provided id with the provided arguments.
// Returns any run errors, or an error if the command was not found.
func (r Runner) RunCommand(id string, args []string) error {
	cmd, ok := r.cmds[id]
	if !ok {
		return fmt.Errorf("unknown command \"%s\"", id)
	}

	cmd.Initialize()
	return cmd.Run(args)
}

// Print the usage for all commands to the console.
func (r Runner) ListCommands() {
	for _, id := range r.ids {
		fmt.Printf("%s\t| %s\n", style.BoldInfo.Sprint(id), r.cmds[id].GetUsage())
	}
}
