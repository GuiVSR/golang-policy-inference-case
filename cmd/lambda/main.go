package main

import (
	http_handler "lab/internal/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestPayload struct {
	PolicyDot string                 `json:"policy_dot"`
	Input     map[string]interface{} `json:"input"`
}

type ResponsePayload struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	switch request.RequestContext.HTTP.Method {
	case "GET":
		getResponse, err := http_handler.HandleGetRequest(request)
		if err != nil {
			return getResponse, err
		}
		return getResponse, nil
	case "POST":
		postResponse, err := http_handler.HandlePostRequest(request)
		if err != nil {
			return postResponse, err
		}
		return postResponse, nil
	default:
		defaultResponse := events.LambdaFunctionURLResponse{
			Body:       `{"error": "Método não permitido. Use POST"}`,
			StatusCode: 405,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}
		return defaultResponse, nil
	}
}
