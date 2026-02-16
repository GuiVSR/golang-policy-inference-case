package http_handler

import (
	"encoding/json"
	"lab/internal/algorithm"
	"lab/internal/logger"
	"lab/internal/models"
	"lab/internal/parser"

	"github.com/aws/aws-lambda-go/events"
)

func HandleGetRequest(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	path := request.RequestContext.HTTP.Path
	switch path {
	case "/healthcheck":
		return healthCheck(), nil
	default:
		return handleNotFound(), nil
	}

}

func HandlePostRequest(request events.LambdaFunctionURLRequest) events.LambdaFunctionURLResponse {
	path := request.RequestContext.HTTP.Path
	var data models.InferRequest
	err := json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		return handleBadRequest(err, request.RequestContext.RequestID)
	}
	switch path {
	case "/infer":
		return infer(data.PolicyDot, data.Input)
	default:
		return handleNotFound()
	}
}

func healthCheck() events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		Body:       `{{"event": "Hello World"}}`,
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func handleBadRequest(err error, requestID string) events.LambdaFunctionURLResponse {
	logger.LogInfo(err.Error(), requestID)
	return events.LambdaFunctionURLResponse{
		Body:       `{{"error": "Bad Request"}}`,
		StatusCode: 400,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func handleNotFound() events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		Body:       `{{"error": "Not Found"}}`,
		StatusCode: 404,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func infer(dotString string, input map[string]interface{}) events.LambdaFunctionURLResponse {
	parsed, _ := parser.ParsePolicy(dotString)
	output, _ := algorithm.EvaluatePolicy(parsed, input)
	jsonBody, _ := json.Marshal(output)
	return events.LambdaFunctionURLResponse{
		Body:       string(jsonBody),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}
