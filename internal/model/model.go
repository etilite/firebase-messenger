package model

type SendRequest struct {
	Tokens       []string          `json:"tokens"`
	Notification Notification      `json:"notification"`
	Data         map[string]string `json:"data,omitempty"`
}

type Notification struct {
	Title string `json:"title"` // Обязательно
	Body  string `json:"body"`  // Обязательно
}

type SendResponse struct {
	SuccessCount int               `json:"successCount"`
	FailureCount int               `json:"failureCount"`
	Responses    []TokenSendResult `json:"responses"`
}

type TokenSendResult struct {
	Success   bool       `json:"success"`
	MessageID string     `json:"messageId,omitempty"`
	Error     *SendError `json:"error,omitempty"`
}

type SendError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
