package logs

import (
	"os"

	"github.com/fatih/color"
)

var (
	normal = color.New(color.FgWhite)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	red    = color.New(color.FgRed)

	ContinueOnError = false
	Verbose         = false
)

func Success(message string) {
	green.Fprint(os.Stdout, "SUC ")
	normal.Fprintln(os.Stdout, message)
}

func Info(message string, objs ...interface{}) {
	if Verbose {
		green.Fprint(os.Stdout, "INF ")
		normal.Fprintln(os.Stdout, message)
		if len(objs) > 0 {
			normal.Fprintln(os.Stdout, objs...)
		}
	}
}
func Warn(message string, objs ...interface{}) {
	yellow.Fprint(os.Stdout, "WRN ")
	normal.Fprintln(os.Stdout, message)
	if len(objs) > 0 {
		for _, obj := range objs {
			if obj != nil {
				normal.Fprintln(os.Stdout, obj)
			}
		}
	}
}
func Error(message string, objs ...interface{}) {
	red.Fprint(os.Stderr, "ERR ")
	normal.Fprintln(os.Stderr, message)
	if len(objs) > 0 {
		for _, obj := range objs {
			if errObj, ok := obj.(error); ok && !ContinueOnError {
				if errObj != nil {
					normal.Fprintln(os.Stderr, errObj)
				}
			} else {
				normal.Fprintln(os.Stderr, obj)
			}
		}
	}
	if !ContinueOnError {
		os.Exit(1)
	}
}
