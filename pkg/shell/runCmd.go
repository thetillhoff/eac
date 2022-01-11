package shell

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

// Takes shell command, executes it and returns output as string
func RunCmd(command string) (string, error) {
	var (
		output string
	)
	logs.Info("Command", command)

	cmd := exec.Command("/bin/sh", "-c", command)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanOut := bufio.NewScanner(stdout)
	scanOut.Split(bufio.ScanWords)
	for scanOut.Scan() {
		m := scanOut.Text()
		output = output + m + " " // Spaces have to be added manually
	}
	err := cmd.Wait()
	// if err != nil {
	// 	logs.Error("command failed", command, reflect.TypeOf(err), err)
	// }

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalln("ERR Error while running command:", err, output)
	// }

	return strings.TrimSpace(string(output)), err
}
