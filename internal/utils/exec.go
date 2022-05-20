package utils

import (
	"bytes"
	"os/exec"
)

func ExecSimple(program string, args ...string) (string, error) {
	cmd := exec.Command(program, args...)
	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer
	err := cmd.Run()
	if err != nil {
		return stdoutBuffer.String(), err
	}
	return stdoutBuffer.String(), err
}
