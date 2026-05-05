package command

type Command interface {
	// Get the command's identifier.
	GetID() string

	// Get the command's usage string.
	GetUsage() string

	// Initializes the command before it's run.
	Initialize()

	// Run the command using the provided arguments and return any errors.
	Run(args []string) error
}
