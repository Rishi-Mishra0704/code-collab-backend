package compiler

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	models "github.com/Rishi-Mishra0704/code-collab-backend/models"
)

// executeJavaCode executes the provided Java code and returns the output or any errors
func executeJavaCode(ctx context.Context, code string) (*models.CodeResponse, error) {
	// Create a temporary directory for the Java files
	tmpDir, err := os.MkdirTemp("", "java")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up the temporary directory

	// Write the Java code to a temporary file
	javaFile := filepath.Join(tmpDir, "Main.java")
	if err := os.WriteFile(javaFile, []byte(code), 0644); err != nil {
		return nil, fmt.Errorf("failed to write Java code to file: %v", err)
	}

	// Compile the Java code
	cmd := exec.CommandContext(ctx, "javac", javaFile)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to compile Java code: %v", err)
	}

	// Execute the compiled bytecode
	cmd = exec.CommandContext(ctx, "java", "-cp", tmpDir, "Main")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.CodeResponse{
			Output: string(output),
			Error:  err.Error(),
		}, fmt.Errorf("failed to execute Java code: %v", err)
	}

	// No errors, return output
	return &models.CodeResponse{
		Output: string(output),
		Error:  "",
	}, nil
}

func executeJavaCodeWithContext(code string) (string, error) {
	ctx := context.Background()
	resp, err := executeJavaCode(ctx, code)
	if err != nil {
		return "", err
	}
	return resp.Output, nil
}
