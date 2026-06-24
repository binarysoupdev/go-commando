package test

import (
	"github.com/binarysoupdev/go-commando/command"
	"github.com/stretchr/testify/suite"
)

// CommandSuite is a testify test suite for commands.
type CommandSuite[T command.Command] struct {
	suite.Suite
	Cmd    T
	result error
}

// Create a new NewCommandSuite for the given command interface.
func NewCommandSuite[T command.Command](cmd T) CommandSuite[T] {
	return CommandSuite[T]{
		Cmd: cmd,
	}
}

// Run the command with the provided arguments and cache the result.
// Note: if initialize fails, that result is cached instead.
func (s *CommandSuite[T]) RunCommand(args ...string) {
	err := s.Cmd.Initialize()
	if err != nil {
		s.result = err
	} else {
		s.result = s.Cmd.Run(args)
	}
}

// Assert that the command returned no error.
func (s *CommandSuite[T]) RequireResultPass() {
	s.Require().NoError(s.result)
}

// Assert the command returned an error that contains the provided message.
func (s *CommandSuite[T]) RequireResultFail(msg string) {
	s.Require().Error(s.result)
	s.Assert().ErrorContains(s.result, msg)
}
