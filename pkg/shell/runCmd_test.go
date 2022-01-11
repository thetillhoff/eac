package shell

import (
	"errors"
	"testing"
)

func TestRunCmd(t *testing.T) {
	testCommands := []struct {
		cmd    string
		output string
		err    error
	}{
		{"echo hello", "hello", nil},
		{"echo hello world", "hello world", nil},
		{"echo $0", "/bin/sh", nil},
		{"exit 1", "", errors.New("exit status 1")},
	}
	// // testCommands with empty value are intended to fail!
	// testCommands map[string]string = map[string]string{
	// 	"echo hello":       "hello",
	// 	"echo hello world": "hello world",
	// 	"echo $0":          "/bin/sh",
	// 	"exit 1":           "",
	// }

	for _, testCommand := range testCommands {
		got, err := RunCmd(testCommand.cmd)

		// If value is not as expected
		if got != testCommand.output {
			t.Error("RunCmd(" + testCommand.cmd + ") = \"" + got + "\"; expected \"" + testCommand.output + "\"")
		}

		// If testCommand should have failed but didn't
		if testCommand.err != nil && err == nil {
			t.Error("RunCmd(" + testCommand.cmd + ").err == nil; expected \"" + testCommand.err.Error() + "\"")
		} else if testCommand.err != nil && err != nil { // If testCommand failed and should have
			if testCommand.err.Error() != err.Error() {
				t.Error("RunCmd(" + testCommand.cmd + ").err == \"" + err.Error() + "\"; expected \"" + testCommand.err.Error() + "\"")
			}
		}
	}
}
