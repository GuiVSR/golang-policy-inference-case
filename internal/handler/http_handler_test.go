package http_handler

import (
	"encoding/json"
	"testing"

	"lab/internal/models"

	"github.com/aws/aws-lambda-go/events"
)

func TestHealthCheck(t *testing.T) {
	req := events.LambdaFunctionURLRequest{
		RequestContext: events.LambdaFunctionURLRequestContext{
			HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
				Path:   "/healthcheck",
				Method: "GET",
			},
		},
	}
	resp, err := HandleGetRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestPostInfer_SuccessAndBadJSON(t *testing.T) {
	reqBody := models.InferRequest{
		PolicyDot: `digraph Policy { start [result=""]; ok [result="approved=true"]; start -> ok [cond="age>=18"] }`,
		Input:     map[string]interface{}{"age": 20},
	}
	b, _ := json.Marshal(reqBody)
	req := events.LambdaFunctionURLRequest{
		RequestContext: events.LambdaFunctionURLRequestContext{
			HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
				Path:   "/infer",
				Method: "POST",
			},
		},
		Body: string(b),
	}
	resp := HandlePostRequest(req)
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 for valid infer request, got %d", resp.StatusCode)
	}

	bad := events.LambdaFunctionURLRequest{
		RequestContext: events.LambdaFunctionURLRequestContext{
			HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
				Path:   "/infer",
				Method: "POST",
			},
		},
		Body: "not-json",
	}
	badResp := HandlePostRequest(bad)
	if badResp.StatusCode != 400 {
		t.Fatalf("expected 400 for bad JSON, got %d", badResp.StatusCode)
	}
}
