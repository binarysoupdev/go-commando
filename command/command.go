// Provides several types for managing multiple commands within a single application.
package command

type Command interface {
	// Returns the command's identifier.
	GetID() string

	// Get the command's usage string.
	GetUsage() string

	// Initializes the command before it's run.
	Initialize()

	// Runs the command using the provided args and return any errors.
	Run(args []string) error
}
