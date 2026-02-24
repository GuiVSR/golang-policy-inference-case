package http_handler

import (
	"lab/internal/logger"

	"github.com/aws/aws-lambda-go/events"
)

func HandleBadRequest(err error, requestID string) events.LambdaFunctionURLResponse {
	logger.LogInfo(err.Error(), requestID)
	return events.LambdaFunctionURLResponse{
		Body:       `{{"error": "Bad Request"}}`,
		StatusCode: 400,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func handleInternalError(err error, requestID string) events.LambdaFunctionURLResponse {
	logger.LogInfo(err.Error(), requestID)
	return events.LambdaFunctionURLResponse{
		Body:       `{{"error": "Internal Server Error"}}`,
		StatusCode: 500,
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
