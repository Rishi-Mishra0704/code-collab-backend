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
			} else {
				var response map[string]string
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Empty(t, response)
			}
		})
	}
}