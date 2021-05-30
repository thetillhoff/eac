package logs

import (
	"os"

	"github.com/fatih/color"
)

var (
	//time = color.New(color.Fg)
	normal = color.New(color.FgWhite)
	info   = color.New(color.FgGreen)
	warn   = color.New(color.FgYellow)
	err    = color.New(color.FgRed)
)

func Info(message string, objs ...interface{}) {
	info.Fprint(os.Stdout, "INF ")
	normal.Fprintln(os.Stdout, message)
	if len(objs) > 0 {
		normal.Fprintln(os.Stdout, objs)
	}
}
func Warn(message string, objs ...interface{}) {
	warn.Fprint(os.Stdout, "WRN ")
	normal.Fprintln(os.Stdout, message)
	if len(objs) > 0 {
		normal.Fprintln(os.Stdout, objs)
	}
}
func Err(message string, objs ...interface{}) {
	err.Fprint(os.Stderr, "ERR ")
	normal.Fprintln(os.Stderr, message)
	if len(objs) > 0 {
		normal.Fprintln(os.Stderr, objs)
	}
	os.Exit(1)
}
