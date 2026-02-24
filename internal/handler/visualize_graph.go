package http_handler

import (
	"lab/internal/parser"

	"github.com/aws/aws-lambda-go/events"
)

func VisualizeGraph(dotString string) events.LambdaFunctionURLResponse {
	i, err := parser.DotToBase64Image(dotString)
	if err != nil {
		return HandleBadRequest(err, err.Error())
	}
	return events.LambdaFunctionURLResponse{
		Body:       `{{"image": "` + i + `"}}`,
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}
