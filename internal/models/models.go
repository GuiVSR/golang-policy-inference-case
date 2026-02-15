package models

type InferRequest struct {
	PolicyDot string                 `json:"policy_dot"`
	Input     map[string]interface{} `json:"input"`
}

type ResponsePayload struct {
	Message string `json:"message"`
}

type LogEntry struct {
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"requestId"`
	Function  string                 `json:"function"`
	Duration  int64                  `json:"duration,omitempty"`
	Error     string                 `json:"error,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
}
