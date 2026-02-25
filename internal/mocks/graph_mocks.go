package mocks

import (
	"lab/internal/models"

	"github.com/Knetic/govaluate"
)

var MockSimpleGraph = func() *models.Graph {
	condApproved, _ := govaluate.NewEvaluableExpression("age >= 18 && score > 700")
	condRejected, _ := govaluate.NewEvaluableExpression("age < 18")
	condReview, _ := govaluate.NewEvaluableExpression("age >= 18 && score <= 700")

	return &models.Graph{
		Nodes: map[string]*models.Node{
			"start": {
				Name:       "start",
				Attributes: map[string]interface{}{"result": ""},
				Edges: []*models.Edge{
					{
						From:      "start",
						To:        "approved",
						Condition: condApproved,
					},
					{
						From:      "start",
						To:        "rejected",
						Condition: condRejected,
					},
					{
						From:      "start",
						To:        "review",
						Condition: condReview,
					},
				},
			},
			"approved": {
				Name:       "approved",
				Attributes: map[string]interface{}{"result": "approved=true,segment=prime"},
				Edges:      []*models.Edge{},
			},
			"rejected": {
				Name:       "rejected",
				Attributes: map[string]interface{}{"result": "approved=false"},
				Edges:      []*models.Edge{},
			},
			"review": {
				Name:       "review",
				Attributes: map[string]interface{}{"result": "approved=false,segment=manual"},
				Edges:      []*models.Edge{},
			},
		},
		Start: "start",
	}
}()

var MockGraphWithInvalidCondition = func() *models.Graph {
	condApproved, _ := govaluate.NewEvaluableExpression("age >= 18 && score > 700")
	condRejected, _ := govaluate.NewEvaluableExpression("age < 18")
	condReview, _ := govaluate.NewEvaluableExpression("age >= 18 && scoe <= 700")

	return &models.Graph{
		Nodes: map[string]*models.Node{
			"start": {
				Name:       "start",
				Attributes: map[string]interface{}{"result": ""},
				Edges: []*models.Edge{
					{
						From:      "start",
						To:        "approved",
						Condition: condApproved,
					},
					{
						From:      "start",
						To:        "rejected",
						Condition: condRejected,
					},
					{
						From:      "start",
						To:        "review",
						Condition: condReview,
					},
				},
			},
			"approved": {
				Name:       "approved",
				Attributes: map[string]interface{}{"result": "approved=true,segment=prime"},
				Edges:      []*models.Edge{},
			},
			"rejected": {
				Name:       "rejected",
				Attributes: map[string]interface{}{"result": "approved=false"},
				Edges:      []*models.Edge{},
			},
			"review": {
				Name:       "review",
				Attributes: map[string]interface{}{"result": "approved=false,segment=manual"},
				Edges:      []*models.Edge{},
			},
		},
		Start: "start",
	}
}()

var MockGraphWithEmptyCondition = func() *models.Graph {
	condApproved, _ := govaluate.NewEvaluableExpression("")
	condReview, _ := govaluate.NewEvaluableExpression("age >= 18 && scoe <= 700")

	return &models.Graph{
		Nodes: map[string]*models.Node{
			"start": {
				Name:       "start",
				Attributes: map[string]interface{}{"result": ""},
				Edges: []*models.Edge{
					{
						From:      "start",
						To:        "approved",
						Condition: condApproved,
					},
					{
						From:      "start",
						To:        "rejected",
						Condition: nil,
					},
					{
						From:      "start",
						To:        "review",
						Condition: condReview,
					},
				},
			},
			"approved": {
				Name:       "approved",
				Attributes: map[string]interface{}{"result": "approved=true,segment=prime"},
				Edges:      []*models.Edge{},
			},
			"rejected": {
				Name:       "rejected",
				Attributes: map[string]interface{}{"result": "approved=false"},
				Edges:      []*models.Edge{},
			},
			"review": {
				Name:       "review",
				Attributes: map[string]interface{}{"result": "approved=false,segment=manual"},
				Edges:      []*models.Edge{},
			},
		},
		Start: "start",
	}
}()
