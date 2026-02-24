package http_handler

import (
	"encoding/json"
	"lab/internal/algorithm"
	"lab/internal/parser"

	"github.com/aws/aws-lambda-go/events"
)

func Infer(dotString string, input map[string]interface{}) events.LambdaFunctionURLResponse {
	parsed, err := parser.ParsePolicy(dotString)
	if err != nil {
		return HandleBadRequest(err, err.Error())
	}

	output, err := algorithm.EvaluatePolicy(parsed, input)
	if err != nil {
		return handleInternalError(err, err.Error())
	}
	jsonBody, _ := json.Marshal(output)
	return events.LambdaFunctionURLResponse{
		Body:       string(jsonBody),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}
