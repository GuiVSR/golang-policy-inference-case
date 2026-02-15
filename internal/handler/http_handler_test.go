package http_handler

import (
	"encoding/json"
	"testing"

	"lab/internal/models"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleGetRequest(t *testing.T) {
	testCases := []struct {
		name           string
		req            events.LambdaFunctionURLRequest
		expectedStatus int
	}{
		{
			name: "HealthCheck",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/healthcheck",
						Method: "GET",
					},
				},
			},
			expectedStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := HandleGetRequest(tc.req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.StatusCode != tc.expectedStatus {
				t.Fatalf("expected %d, got %d", tc.expectedStatus, resp.StatusCode)
			}
		})
	}
}

func TestHandlePostRequest(t *testing.T) {
	reqBody := models.InferRequest{
		PolicyDot: `digraph Policy { start [result=""]; ok [result="approved=true"]; start -> ok [cond="age>=18"] }`,
		Input:     map[string]interface{}{"age": 20},
	}
	b, _ := json.Marshal(reqBody)

	testCases := []struct {
		name           string
		req            events.LambdaFunctionURLRequest
		expectedStatus int
	}{
		{
			name: "InferSuccess",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: string(b),
			},
			expectedStatus: 200,
		},
		{
			name: "BadJSON",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: "not-json",
			},
			expectedStatus: 400,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := HandlePostRequest(tc.req)
			if resp.StatusCode != tc.expectedStatus {
				t.Fatalf("expected %d, got %d", tc.expectedStatus, resp.StatusCode)
			}
		})
	}
}
