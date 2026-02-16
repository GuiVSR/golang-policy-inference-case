package models

import "github.com/Knetic/govaluate"

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

type Node struct {
	Name       string
	Attributes map[string]interface{}
}

type Edge struct {
	From      string
	To        string
	Condition *govaluate.EvaluableExpression
}

type Graph struct {
	Nodes map[string]*Node
	Edges []*Edge
	Start string
}

type ResultNodes map[string]map[string]string

type Edges []*Edge
