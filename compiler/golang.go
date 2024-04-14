package compiler

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	models "github.com/Rishi-Mishra0704/code-collab-backend/models"
)

// executeGoCode executes the provided Go code and returns the output or any errors
func executeGoCode(ctx context.Context, code string) (*models.CodeResponse, error) {
	// Create a temporary file for the Go code
	tmpfile, err := os.CreateTemp("", "code*.go")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the temporary file

	// Write the code to the temporary file
	if _, err := tmpfile.WriteString(code); err != nil {
		return nil, fmt.Errorf("failed to write code to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close temporary file: %v", err)
	}

	// Create a new context with timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Prepare the command to compile and execute the Go code
	cmd := exec.CommandContext(ctx, "go", "run", tmpfile.Name())

	// Setup output buffers
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err = cmd.Run()

	// Check for context cancellation
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("execution timed out")
	}

	// Check for other errors
	if err != nil {
		return &models.CodeResponse{
			Output: stdout.String(),
			Error:  stderr.String(),
		}, fmt.Errorf("execution failed: %v", err)
	}

	// No errors, return output
	return &models.CodeResponse{
		Output: stdout.String(),
		Error:  stderr.String(),
	}, nil
}

func executeGoCodeWithContext(code string) (string, error) {
	ctx := context.Background()
	resp, err := executeGoCode(ctx, code)
	if err != nil {
		return "", err
	}
	return resp.Output, nil
}
