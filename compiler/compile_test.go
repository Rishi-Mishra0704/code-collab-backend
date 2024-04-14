package compiler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCodeHandler(t *testing.T) {
	testCases := []struct {
		Name         string
		Payload      interface{}
		ExpectedCode int
	}{
		{
			Name:         "Valid payload - JS",
			Payload:      map[string]interface{}{"language": "js", "code": "console.log('Hello, World!')"},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Valid payload - Python",
			Payload:      map[string]interface{}{"language": "py", "code": "print('Hello, World!')"},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Valid payload - Ruby",
			Payload:      map[string]interface{}{"language": "rb", "code": "print('Hello, World!')"},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Valid payload - Go",
			Payload:      map[string]interface{}{"language": "go", "code": "package main\n\nimport (\n\t\"fmt\"\n\t\"time\"\n)\n\nfunc main() {\n\tcurrentTime := time.Now()\n\tformattedTime := currentTime.Format(\"Monday, January 2, 2006 15:04:05\")\n\tfmt.Println(\"Formatted time:\", formattedTime)\n}"},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Valid payload - Java",
			Payload:      map[string]interface{}{"language": "java", "code": "public class Main {\n\tpublic static void main(String[] args) {\n\t\tSystem.out.println(\"Hello, World!\");\n\t}\n}"},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "Invalid payload - Unsupported",
			Payload:      map[string]interface{}{"language": "test_lang", "code": "dsfsd"},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			payload, err := json.Marshal(tc.Payload)
			assert.NoError(t, err)

			req, err := http.NewRequest("POST", "/execute", bytes.NewReader(payload))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ExecuteCodeHandler)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.ExpectedCode, rr.Code)

			if tc.ExpectedCode == http.StatusOK {
				var response map[string]string
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Contains(t, response, "output")
				assert.Empty(t, response["error"])

			} else if tc.ExpectedCode == http.StatusBadRequest {
				// Read the response body
				responseBody := rr.Body.String()

				// Assert that the response body matches the expected error message
				assert.Equal(t, "Unsupported language\n", responseBody)

			} else {
				var response map[string]string
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Empty(t, response)
			}
		})
	}
}
