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
	testCases := []struct {
		name           string
		reqBody        models.InferRequest
		expectedStatus int
	}{
		{
			name: "InferToSuccess",
			reqBody: models.InferRequest{
				PolicyDot: `digraph Policy { start [result=""]; approved [ result="approved=true,segment=prime" ]; rejected [ result="approved=false" ]; review [ result="approved=false,segment=manual" ]; start -> approved [cond="age>=18 && score>700"]; start -> review [cond="age>=18 && score<=700"]; start -> rejected [cond="age<18"]; }`,
				Input:     map[string]interface{}{"age": 20, "score": 750},
			},
			expectedStatus: 200,
		},
		{
			name: "InferToRejectionAge15",
			reqBody: models.InferRequest{
				PolicyDot: `digraph Policy { start [result=""]; approved [ result="approved=true,segment=prime" ]; rejected [ result="approved=false" ]; review [ result="approved=false,segment=manual" ]; start -> approved [cond="age>=18 && score>700"]; start -> review [cond="age>=18 && score<=700"]; start -> rejected [cond="age<18"]; }`,
				Input:     map[string]interface{}{"age": 15, "score": 300},
			},
			expectedStatus: 200,
		},
		{
			name: "InferVeryComplexDigraph",
			reqBody: models.InferRequest{
				PolicyDot: `digraph Policy { start [result=""]; age_check [result=""]; income_check [result=""]; credit_check [result=""]; employment_check [result=""]; debt_check [result=""]; savings_check [result=""]; expenses_check [result=""]; home_check [result=""]; education_check [result=""]; marital_check [result=""]; approved [result="approved=true,segment=prime"]; rejected [result="approved=false"]; start -> age_check [cond="age >= 18"]; age_check -> income_check [cond="income > 30000"]; income_check -> credit_check [cond="credit_score > 650"]; credit_check -> employment_check [cond="employment_years >= 2"]; employment_check -> debt_check [cond="debt_to_income_ratio < 0.4"]; debt_check -> savings_check [cond="savings_amount > 10000"]; savings_check -> expenses_check [cond="monthly_expenses < income * 0.3"]; expenses_check -> home_check [cond="home_ownership == true"]; home_check -> education_check [cond="education_level == 'bachelor' || education_level == 'master'"]; education_check -> marital_check [cond="marital_status == 'married'"]; marital_check -> approved; start -> rejected [cond="age < 18"]; age_check -> rejected [cond="income <= 30000"]; income_check -> rejected [cond="credit_score <= 650"]; credit_check -> rejected [cond="employment_years < 2"]; employment_check -> rejected [cond="debt_to_income_ratio >= 0.4"]; debt_check -> rejected [cond="savings_amount <= 10000"]; savings_check -> rejected [cond="monthly_expenses >= income * 0.3"]; expenses_check -> rejected [cond="home_ownership == false"]; home_check -> rejected [cond="education_level != 'bachelor' && education_level != 'master'"]; education_check -> rejected [cond="marital_status != 'married'"]; }`,
				Input: map[string]interface{}{
					"age":                   30,
					"income":                60000,
					"credit_score":          750,
					"employment_years":      5,
					"debt_to_income_ratio":  0.3,
					"savings_amount":        20000,
					"monthly_expenses":      15000,
					"home_ownership":        true,
					"education_level":       "bachelor",
					"marital_status":        "married",
				},
			},
			expectedStatus: 200,
		},
		{
			name:           "BadJSON",
			reqBody:        models.InferRequest{}, // Not used for BadJSON
			expectedStatus: 400,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var req events.LambdaFunctionURLRequest
			if tc.name == "BadJSON" {
				req = events.LambdaFunctionURLRequest{
					RequestContext: events.LambdaFunctionURLRequestContext{
						HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
							Path:   "/infer",
							Method: "POST",
						},
					},
					Body: "not-json",
				}
			} else {
				b, _ := json.Marshal(tc.reqBody)
				req = events.LambdaFunctionURLRequest{
					RequestContext: events.LambdaFunctionURLRequestContext{
						HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
							Path:   "/infer",
							Method: "POST",
						},
					},
					Body: string(b),
				}
			}
			resp := HandlePostRequest(req)
			if resp.StatusCode != tc.expectedStatus {
				t.Fatalf("expected %d, got %d", tc.expectedStatus, resp.StatusCode)
			}
		})
	}
}
