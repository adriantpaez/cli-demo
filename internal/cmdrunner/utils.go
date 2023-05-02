package cmdrunner

import "os/exec"

func runCmd(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.Output()
}
