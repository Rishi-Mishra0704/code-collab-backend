package console

import (
	"fmt"
	"io"
	"os/exec"
	"runtime"
)

func getShell() string {
	switch runtime.GOOS {
	case "windows":
		return "powershell"
	case "darwin":
		return "zsh"
	default:
		return "bash"
	}
}

func CallTerminal(command string) (string, error) {
	shell := getShell()

	cmd := exec.Command(shell, "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stdout pipe: %w", err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command: %w", err)
	}

	output, err := io.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("error reading output: %w", err)
	}

	// Wait for the command to finish executing
	if err := cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			// Command did not exit cleanly, but the error is not an ExitError
			return "", fmt.Errorf("error waiting for command to finish: %w", err)
		}
		// Command exited with non-zero status
		return string(output), fmt.Errorf("command exited with non-zero status: %s", exitErr)
	}

	return string(output), nil
}
