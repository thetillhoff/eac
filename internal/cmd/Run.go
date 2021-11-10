package cmd

import (
	"os/exec"
	"strings"
)

func Run(command string, args ...string) (string, error) {
	var (
		tempFolder string
		err        error
		output     []byte
	)

	tempFolder = createTempFolder()
	defer deleteTempFolder(tempFolder)

	args = append([]string{"-c", command}, args...)
	args = append(args, tempFolder)

	output, err = exec.Command("/bin/sh", args...).Output()

	return strings.TrimSuffix(string(output), "\n"), err
}
