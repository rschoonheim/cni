package ip

import (
	"errors"
	"os/exec"
	"regexp"
)

var allowed_command_patterns = []string{}

// CommandNotAllowedError - error for when a command is not allowed.
type CommandNotAllowedError struct {
	Command string
}

// CommandVerify - verifies that a command can be executed
var CommandVerify = func(command string, args ...string) *CommandNotAllowedError {
	// Check if command passes any regex patterns.
	//
	for _, pattern := range allowed_command_patterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			continue
		}

		if regex.MatchString(command) {
			return nil
		}
	}

	// Return an error if the command is not in the white list.
	//
	return &CommandNotAllowedError{
		Command: command,
	}
}

// CommandExecute - executes a command.
var CommandExecute = func(command string, args ...string) ([]byte, error) {
	// Verify that the command can be executed.
	//
	canbeExecuted := CommandVerify(command, args...)
	if canbeExecuted != nil {
		return nil, errors.New("COMMAND_NOT_ALLOWED")
	}

	// Execute the command.
	//
	cmd, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	// Return the output of the command.
	//
	return cmd, nil
}
