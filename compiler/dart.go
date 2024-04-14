package compiler

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	models "github.com/Rishi-Mishra0704/code-collab-backend/models"
)

// executeDartCode executes the provided Dart code and returns the output or any errors
func executeDartCode(ctx context.Context, code string) (*models.CodeResponse, error) {
	// Create a temporary directory for the Dart files
	tmpDir, err := os.MkdirTemp("", "dart")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up the temporary directory

	// Write the Dart code to a temporary file
	dartFile := filepath.Join(tmpDir, "main.dart")
	if err := os.WriteFile(dartFile, []byte(code), 0644); err != nil {
		return nil, fmt.Errorf("failed to write Dart code to file: %v", err)
	}

	// Compile the Dart code
	cmd := exec.CommandContext(ctx, "dart", dartFile)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return &models.CodeResponse{
			Output: stdout.String(),
			Error:  stderr.String(),
		}, fmt.Errorf("failed to compile Dart code: %v", err)
	}

	// Execute the compiled Dart code
	cmd = exec.CommandContext(ctx, "dart", dartFile)
	stdout.Reset()
	stderr.Reset()
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return &models.CodeResponse{
			Output: stdout.String(),
			Error:  stderr.String(),
		}, fmt.Errorf("failed to execute Dart code: %v", err)
	}

	// No errors, return output
	return &models.CodeResponse{
		Output: stdout.String(),
		Error:  stderr.String(),
	}, nil
}

func executeDartCodeWithContext(code string) (string, error) {
	ctx := context.Background()
	resp, err := executeDartCode(ctx, code)
	if err != nil {
		return "", err
	}
	return resp.Output, nil
}
