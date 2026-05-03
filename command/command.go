// Provides several types for managing multiple commands within a single application.
package command

// The interface definition of a command.
type Command interface {
	// Returns the command's name (essentially it's identifier).
	GetName() string

	// Prints usage information.
	PrintUsage()

	// Initializes the command before it's run.
	Initialize()

	// Runs the command using the provided args and returns any error.
	Run(args []string) error
}
