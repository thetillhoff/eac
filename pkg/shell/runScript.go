package shell

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

// Takes path to shell script, executes it and returns output as string
func RunLinuxScriptAt(scriptPath string) (string, error) {
	var (
		output string
	)
	logs.Info("scriptPath", scriptPath)

	cmd := exec.Command("/bin/sh", "-c", scriptPath)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanOut := bufio.NewScanner(stdout)
	scanOut.Split(bufio.ScanWords)
	for scanOut.Scan() {
		m := scanOut.Text()
		output = output + m + " " // Spaces have to be added manually
	}
	err := cmd.Wait()

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalln("ERR Error while running command:", err, output)
	// }

	return strings.TrimSpace(string(output)), err
}
