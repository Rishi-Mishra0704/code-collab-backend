package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func getUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w ", err)
	}
	return strings.TrimSpace(input), nil
}

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

func CallTerminal() {
	shell := getShell()

	for {
		command, err := getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if command == "exit" {
			break
		}
		cmd := exec.Command(shell, "-c", command)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error creating stdout pipe:", err)
			continue
		}
		err = cmd.Start()
		if err != nil {
			fmt.Println("Error starting command:", err)
			continue
		}
		output, err := io.ReadAll(stdout)
		if err != nil {
			fmt.Println("Error reading output:", err)
			continue
		}
		fmt.Println(string(output))

		cmd.Wait()
	}
	fmt.Println("Exiting...")
}
