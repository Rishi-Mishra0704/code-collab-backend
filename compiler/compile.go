package compiler

import (
	"encoding/json"
	"net/http"
	"strings"

	models "github.com/Rishi-Mishra0704/code-collab-backend/models"
)

func ExecuteCodeHandler(w http.ResponseWriter, r *http.Request) {
	var codeReq models.CodeRequest
	var output string
	var err error
	if err := json.NewDecoder(r.Body).Decode(&codeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch codeReq.Language {
	case "go":
		output, err = executeGoCodeWithContext(codeReq.Code)
	case "js":
		output, err = executeNodeCodeWithContext(codeReq.Code)
	case "py":
		output, err = executePythonCodeWithContext(codeReq.Code)
	case "rb":
		output, err = executeRubyCodeWithContext(codeReq.Code)
	case "java":
		output, err = executeJavaCodeWithContext(codeReq.Code)
	case "dart":
		output, err = executeDartCodeWithContext(codeReq.Code)
	default:
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.CodeResponse{
		Output: strings.TrimRight(output, "\n"),
		Error:  "",
	}
	json.NewEncoder(w).Encode(response)
}
