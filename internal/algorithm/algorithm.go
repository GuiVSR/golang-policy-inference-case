package algorithm

import (
	"fmt"
	"lab/internal/models"
	"maps"

	"github.com/Knetic/govaluate"
)

func evalCondition(expr *govaluate.EvaluableExpression, vars map[string]interface{}) (bool, error) {
	if expr == nil {
		return true, nil
	}
	result, err := expr.Evaluate(vars)
	if err != nil {
		return false, err
	}
	if resBool, ok := result.(bool); ok {
		return resBool, nil
	}
	return false, fmt.Errorf("unexpected result type: %T", result)
}

func dfs(g *models.Graph, current string, input map[string]interface{}, visited map[string]bool) (map[string]string, error) {
	if visited[current] {
		return nil, nil
	}
	results := make(map[string]string)
	currNode := g.Nodes[current]
	visited[current] = true

	for k, v := range currNode.Attributes {
		results[k] = v.(string)
	}

	for _, edge := range currNode.Edges {
		if visited[edge.To] {
			continue
		} else {
			res, err := evalCondition(edge.Condition, input)
			if err != nil {
				return nil, err
			}
			if res {
				visited[edge.To] = false
				edgeResults, err := dfs(g, edge.To, input, visited)
				if err != nil {
					return nil, err
				}
				maps.Copy(results, edgeResults)
			}
		}
	}

	return results, nil
}

func EvaluatePolicy(g *models.Graph, input map[string]interface{}) (map[string]string, error) {
	visited := make(map[string]bool)
	visited[g.Start] = false
	res, err := dfs(g, g.Start, input, visited)
	if err != nil {
		return nil, err
	}
	return res, nil
}
