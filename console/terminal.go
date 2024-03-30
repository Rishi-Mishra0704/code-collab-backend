package console

import (
	"bytes"
	"fmt"
	"io"
	"log" // Import log package for logging
	"os/exec"
	"runtime"
	"strings"
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

	log.Printf("Executing command: %s\n", command)

	// Check if the command is an interactive session
	isInteractive := false
	if strings.Contains(command, "python") || strings.Contains(command, "irb") || strings.Contains(command, "node") {
		isInteractive = true
	}

	cmd := exec.Command(shell, "-c", command)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stdout pipe: %w", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stderr pipe: %w", err)
	}

	var outputBuf bytes.Buffer

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command: %w", err)
	}

	if isInteractive {
		// For interactive sessions, we need to read from both stdout and stderr continuously
		go func() {
			io.Copy(&outputBuf, stdoutPipe)
		}()
		go func() {
			io.Copy(&outputBuf, stderrPipe)
		}()
	} else {
		// For non-interactive sessions, we just need to read from stdout
		_, err := io.Copy(&outputBuf, stdoutPipe)
		if err != nil {
			return "", fmt.Errorf("error reading output: %w", err)
		}
	}

	// Wait for the command to finish executing
	if err := cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			// Command did not exit cleanly, but the error is not an ExitError
			return outputBuf.String(), fmt.Errorf("error waiting for command to finish: %w", err)
		}
		// Command exited with non-zero status
		return outputBuf.String(), fmt.Errorf("command exited with non-zero status: %s", exitErr)
	}

	return outputBuf.String(), nil
}
