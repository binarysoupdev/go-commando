// Package sample provides sample commands that demonstrate the command package.
//
// Boiler plate flag-based command for convenience:
//
//	type SampleCommand struct {
//		command.FlagCommandBase
//	}
//
//	func NewSampleCommand() *SampleCommand {
//		return &SampleCommand{
//			FlagCommandBase: command.NewFlagCommandBase("name", "description"),
//		}
//	}
//
//	func (cmd SampleCommand) Run(args []string) error {
//		return nil
//	}
package sample
