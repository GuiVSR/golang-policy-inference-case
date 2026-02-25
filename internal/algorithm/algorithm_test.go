package algorithm

import (
	"encoding/json"
	"lab/internal/mocks"
	"lab/internal/models"
	"reflect"
	"testing"
)

func TestEvaluatePolicy(t *testing.T) {
	testCases := []struct {
		name           string
		graph          *models.Graph
		input          string
		expectedResult map[string]string
	}{
		{
			name:           "Simple Graph - Approved",
			graph:          mocks.MockSimpleGraph,
			input:          mocks.InputSimpleApproved,
			expectedResult: map[string]string{"result": "approved=true,segment=prime"},
		},
		{
			name:           "Simple Graph - Review",
			graph:          mocks.MockSimpleGraph,
			input:          mocks.InputSimpleReview,
			expectedResult: map[string]string{"result": "approved=false,segment=manual"},
		},
		{
			name:           "Simple Graph - Rejected",
			graph:          mocks.MockSimpleGraph,
			input:          mocks.InputSimpleRejected,
			expectedResult: map[string]string{"result": "approved=false"},
		},
		{
			name:           "Simple Graph - Empty Condition",
			graph:          mocks.MockSimpleGraph,
			input:          mocks.InputSimpleRejected,
			expectedResult: map[string]string{"result": "approved=false"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var parsedInput map[string]interface{}
			err := json.Unmarshal([]byte(tc.input), &parsedInput)
			if err != nil {
				t.Fatalf("Failed to unmarshal input: %v", err)
			}
			resultMap, err := EvaluatePolicy(tc.graph, parsedInput)

			if !reflect.DeepEqual(resultMap, tc.expectedResult) {
				t.Fatalf("Expected %v, got %v", tc.expectedResult, resultMap)
			}
		})
	}
}

func TestEvaluatePolicyErrors(t *testing.T) {
	testCases := []struct {
		name  string
		graph *models.Graph
		input string
	}{
		{
			name:  "Invalid Condition Expression",
			graph: mocks.MockGraphWithInvalidCondition,
			input: mocks.InvalidInput1,
		},
		{
			name:  "Missing Input Field",
			graph: mocks.MockSimpleGraph,
			input: mocks.MissingFieldInput,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var parsedInput map[string]interface{}
			_ = json.Unmarshal([]byte(tc.input), &parsedInput)

			if tc.graph == nil {
				return
			}

			_, err := EvaluatePolicy(tc.graph, parsedInput)

			if err == nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}
