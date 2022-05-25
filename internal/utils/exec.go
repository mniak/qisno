package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var execDebug bool

func SetExecDebug(value bool) {
	execDebug = value
}

func Exec(program string, args ...string) (string, error) {
	return ExecContext(context.Background(), program, args...)
}

func ExecContext(ctx context.Context, program string, args ...string) (string, error) {
	if execDebug {
		fmt.Fprintf(os.Stderr, "Exec: %s %s\n", program, strings.Join(args, " "))
	}
	cmd := exec.CommandContext(ctx, program, args...)
	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer
	err := cmd.Run()

	if execDebug {
		fmt.Fprintf(os.Stderr, "Exec stdout: %s\n", stdoutBuffer.String())
		fmt.Fprintf(os.Stderr, "Exec stderr: %s\n", stderrBuffer.String())
	}

	if err != nil {
		return stdoutBuffer.String(), err
	}
	return stdoutBuffer.String(), err
}

func ExecInteractive(program string, args ...string) error {
	return ExecInteractiveContext(context.Background(), program, args...)
}

func ExecInteractiveContext(ctx context.Context, program string, args ...string) error {
	if execDebug {
		fmt.Fprintf(os.Stderr, "Exec: %s %s\n", program, strings.Join(args, " "))
	}
	cmd := exec.CommandContext(ctx, program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
