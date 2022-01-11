package app

import (
	"os/exec"
	"path"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

func RunScript(appName string, appsDirPath string, platform string, script string, args ...string) (string, error) {
	var (
		cmd *exec.Cmd
	)

	tmpFolder := createTmpFolder()

	scriptWithArgs := []string{}
	scriptWithArgs = append(scriptWithArgs, path.Join(appsDirPath, appName, platform, script)) // f.e. 'apps/eac/linux/install.sh'
	scriptWithArgs = append(scriptWithArgs, args...)                                           // f.e. 'apps/eac/linux/install.sh 1.2.3' (version as $1)
	scriptWithArgs = append(scriptWithArgs, tmpFolder)                                         // f.e. 'apps/eac/linux/install.sh 1.2.3 /tmp/1234' (tmpFolder as $2)

	command := strings.Join(scriptWithArgs, " ") + ""

	cmd = exec.Command("/bin/sh", "-c", command) // sadly, the param '-c' cannot be prepended to 'command', go doesn't like that somehow (invalid param)
	outBytes, err := cmd.CombinedOutput()
	out := strings.TrimSuffix(string(outBytes), "\n")
	if err != nil {
		deleteTmpFolder(tmpFolder)
		logs.Error("Failed to run script '"+script+"':", out, err)
	}

	deleteTmpFolder(tmpFolder)

	return out, err
}
