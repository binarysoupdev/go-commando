package test

import (
	"github.com/binarysoupdev/go-commando/command"
	"github.com/stretchr/testify/suite"
)

// CommandSuite is a test suite for commands.
// Provides helpers for running the command with different assertions.
type CommandSuite[T command.Command] struct {
	suite.Suite
	Cmd T

	result error
}

// Creates a new NewCommandSuite for the given command interface.
func NewCommandSuite[T command.Command](cmd T) CommandSuite[T] {
	return CommandSuite[T]{
		Cmd: cmd,
	}
}

func (s *CommandSuite[T]) RunCommand(args ...string) {
	s.Cmd.Initialize()
	s.result = s.Cmd.Run(args)
}

func (s *CommandSuite[T]) RequireResultPass() {
	s.Require().NoError(s.result)
}

func (s *CommandSuite[T]) RequireResultFail(err string) {
	s.Require().Error(s.result)
	s.Require().ErrorContains(s.result, err)
}
