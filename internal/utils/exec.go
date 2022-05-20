package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var execDebug bool

func SetExecDebug(value bool) {
	execDebug = value
}

func ExecSimple(program string, args ...string) (string, error) {
	fmt.Fprintf(os.Stderr, "Exec: %s %s\n", program, strings.Join(args, " "))

	cmd := exec.Command(program, args...)
	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer
	err := cmd.Run()

	fmt.Fprintf(os.Stderr, "Exec stdout: %s\n", stdoutBuffer.String())
	fmt.Fprintf(os.Stderr, "Exec stderr: %s\n", stderrBuffer.String())

	if err != nil {
		return stdoutBuffer.String(), err
	}
	return stdoutBuffer.String(), err
}
