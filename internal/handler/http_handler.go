package http_handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"lab/internal/algorithm"
	"lab/internal/logger"
	"lab/internal/models"
	"lab/internal/parser"

	"github.com/aws/aws-lambda-go/events"
	"github.com/goccy/go-graphviz"
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
		return handleBadRequest(err, request.RequestContext.RequestID)
	}
	switch path {
	case "/infer":
		return infer(input.PolicyDot, input.Input)
	case "/visualize":
		return visualizeGraph(input.PolicyDot)
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

func infer(dotString string, input map[string]interface{}) events.LambdaFunctionURLResponse {
	parsed, err := parser.ParsePolicy(dotString)
	if err != nil {
		return handleBadRequest(err, err.Error())
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

func visualizeGraph(dotString string) events.LambdaFunctionURLResponse {
	g, _ := parser.BuildRenderGraph(dotString)
	gv, _ := graphviz.New(context.Background())
	defer gv.Close()

	var buf bytes.Buffer
	gv.Render(context.Background(), g, graphviz.PNG, &buf)

	return events.LambdaFunctionURLResponse{
		Body:       base64.StdEncoding.EncodeToString(buf.Bytes()),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":        "image/png",
			"Content-Disposition": "inline; filename=\"graph.png\"",
		},
		IsBase64Encoded: true,
	}
}

func encodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
