package http_handler

import (
	"github.com/aws/aws-lambda-go/events"
)

func HandleGetRequest(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	path := request.RequestContext.HTTP.Path
	switch path {
	case "/healthcheck":
		return healthCheck(), nil
	default:
		return handleBadRequest(), nil
	}

}

func HandlePostRequest(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	path := request.RequestContext.HTTP.Path
	switch path {
	case "/infer":
		return infer(), nil
	default:
		return handleBadRequest(), nil
	}
}

func healthCheck() events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		Body:       `{{"event": "Hello World"}}`,
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func handleBadRequest() events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		Body:       `{{"error": "Bad Request"}}`,
		StatusCode: 400,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func infer() events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		Body:       `{{"TODO": "Resto da aplicação continua daqui"}}`,
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}
