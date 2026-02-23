package http_handler

import (
	"fmt"
	"lab/internal/mocks"
	"testing"

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
	testCases := []struct {
		name           string
		req            events.LambdaFunctionURLRequest
		expectedStatus int
		expectedResult string
	}{

		{
			name: "Simple Graph - Approved",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockSimple, mocks.InputSimpleApproved),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"segment\":\"prime\"}",
		},
		{
			name: "Simple Graph - Review",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockSimple, mocks.InputSimpleReview),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"false\",\"segment\":\"manual\"}",
		},
		{
			name: "Simple Graph - Rejected",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockSimple, mocks.InputSimpleRejected),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"false\"}",
		},
		{
			name: "Complex Graph - Prime",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexPrime),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"experienced_worker\":\"true\",\"high_income\":\"true\",\"isInDebt\":\"false\",\"legal_age\":\"true\",\"segment\":\"prime\"}",
		},
		{
			name: "Complex Graph - Subprime",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexSubprime),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"income_compensates\":\"true\",\"legal_age\":\"true\",\"segment\":\"manual\"}",
		},
		{
			name: "Complex Graph - Manual Review",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexManualReview),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"priority\":\"high\",\"segment\":\"manual\"}",
		},
		{
			name: "Complex Graph - Conditional Approval",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexConditionalApproval),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"conditions\":\"lower_rate\",\"legal_age\":\"true\",\"segment\":\"conditional\"}",
		},
		{
			name: "Complex Graph - Conditional Approval Experienced Worker",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexConditionalApprovalExperiencedWorker),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"experienced_worker\":\"true\",\"income_compensates\":\"true\",\"legal_age\":\"true\",\"segment\":\"manual\"}",
		},
		{
			name: "Complex Graph - Conditional Approval Experienced Worker",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexConditionalApprovalHighIncome),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"experienced_worker\":\"true\",\"income_compensates\":\"true\",\"isInDebt\":\"false\",\"legal_age\":\"true\",\"segment\":\"manual\"}",
		},
		{
			name: "Complex Graph - Conditional Approval High Income",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexConditionalApprovalHighIncome),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"experienced_worker\":\"true\",\"income_compensates\":\"true\",\"isInDebt\":\"false\",\"legal_age\":\"true\",\"segment\":\"manual\"}",
		},
		{
			name: "Complex Graph - Near Prime",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexNearPrime),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"true\",\"income_compensates\":\"true\",\"legal_age\":\"true\",\"segment\":\"manual\"}",
		},
		{
			name: "Complex Graph - Rejected",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Path:   "/infer",
						Method: "POST",
					},
				},
				Body: fmt.Sprintf(`{
					"policy_dot": "%s",
					"input": %s
				}`, mocks.MockComplex, mocks.InputComplexRejected),
			},
			expectedStatus: 200,
			expectedResult: "{\"approved\":\"false\"}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := HandlePostRequest(tc.req)
			if resp.StatusCode != tc.expectedStatus {
				t.Fatalf("expected %d, got %d", tc.expectedStatus, resp.StatusCode)
			}
			if resp.Body != tc.expectedResult {
				t.Fatalf("expected %s, got %s", tc.expectedResult, resp.Body)
			}
			_ = resp
		})
	}
}
