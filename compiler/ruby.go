package compiler

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/Rishi-Mishra0704/code-collab-backend/models"
)

// executeRubyCode executes the provided Ruby code and returns the output or any errors
func executeRubyCode(ctx context.Context, code string) (*models.CodeResponse, error) {
	// Create a new context with timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Prepare the command to execute Ruby code
	cmd := exec.CommandContext(ctx, "ruby", "-e", code)

	// Setup output buffers
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()

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

func executeRubyCodeWithContext(code string) (string, error) {
	ctx := context.Background()
	resp, err := executeRubyCode(ctx, code)
	if err != nil {
		return "", err
	}
	return resp.Output, nil
}
