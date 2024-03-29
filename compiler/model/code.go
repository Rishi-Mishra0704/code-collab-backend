package models

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type CodeResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}
