package command

import "fmt"

// Runner stores multiple commands and provides methods for running by name.
type Runner struct {
	cmds  map[string]Command
	names []string
}

// Creates a new runner using the provided commands.
func NewRunner(commands ...Command) Runner {
	cmds := map[string]Command{}
	names := make([]string, len(commands))

	for i, cmd := range commands {
		cmds[cmd.GetName()] = cmd
		names[i] = cmd.GetName()
	}

	return Runner{
		cmds:  cmds,
		names: names,
	}
}

// Run the command that matches the provided name (case sensitive) with the provided arguments.
// Returns any run errors, or an error if the command was not found.
func (r Runner) RunCommand(name string, args []string) error {
	cmd, ok := r.cmds[name]
	if !ok {
		return fmt.Errorf("unknown command \"%s\"", name)
	}

	cmd.Initialize()
	return cmd.Run(args)
}

// Print the usage for all commands to the console.
func (r Runner) ListCommands() {
	for _, name := range r.names {
		r.cmds[name].PrintUsage()
	}
}
