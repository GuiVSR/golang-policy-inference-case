package http_handler

import (
	"encoding/json"
	"lab/internal/models"

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
	var input models.InferRequest
	err := json.Unmarshal([]byte(request.Body), &input)
	if err != nil {
		return HandleBadRequest(err, request.RequestContext.RequestID)
	}
	switch path {
	case "/infer":
		return Infer(input.PolicyDot, input.Input)
	case "/visualize":
		return VisualizeGraph(input.PolicyDot)
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
