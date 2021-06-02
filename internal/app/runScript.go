package app

import (
	"os/exec"
	"path"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

func RunScript(shell string, appName string, appsDirPath string, platform string, script string, appContinueOnError bool, args ...string) (string, error) {
	var (
		cmd *exec.Cmd
	)

	tmpFolder := createTmpFolder(appContinueOnError)

	scriptWithArgs := []string{}
	scriptWithArgs = append(scriptWithArgs, path.Join(appsDirPath, appName, platform, script)) // f.e. 'apps/eac/linux/install.sh'
	scriptWithArgs = append(scriptWithArgs, args...)                                           // f.e. 'apps/eac/linux/install.sh 1.2.3' (version as $1)
	scriptWithArgs = append(scriptWithArgs, tmpFolder)                                         // f.e. 'apps/eac/linux/install.sh 1.2.3 /tmp/1234' (tmpFolder as $2)

	command := strings.Join(scriptWithArgs, " ") + ""

	if strings.Contains(shell, " ") { // f.e. '/bin/sh -c'
		splittedShell := strings.Split(shell, " ")
		cmd = exec.Command(splittedShell[0], splittedShell[1], command) // sadly, the param '-c' cannot be prepended to 'command', go doesn't like that somehow (invalid param)
	} else {
		cmd = exec.Command(shell, command)
	}
	outBytes, err := cmd.CombinedOutput()
	out := strings.TrimSuffix(string(outBytes), "\n")
	if err != nil {
		deleteTmpFolder(tmpFolder, true)
		logs.Err("Failed to run script '"+script+"':", appContinueOnError, out, err)
	}

	deleteTmpFolder(tmpFolder, appContinueOnError)

	return out, err
}
