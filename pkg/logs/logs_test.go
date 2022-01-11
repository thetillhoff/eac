package logs

import "testing"

func TestInfo(*testing.T) {
	Info("INFO")
	// Output: INFO
}

func TestSuccess(*testing.T) {
	Success("SUCCESS")
	// Output: SUCCESS
}

func TestWarn(*testing.T) {
	Warn("WARN")
	// Output: WARN
}

func TestError(*testing.T) {
	ContinueOnError = true
	Error("ERROR")
	// Output: ERROR
}
