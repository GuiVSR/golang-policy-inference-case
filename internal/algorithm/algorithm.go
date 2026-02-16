package algorithm

import (
	"fmt"

	"lab/internal/models"

	"github.com/Knetic/govaluate"
)

func evalCondition(expr *govaluate.EvaluableExpression, vars map[string]interface{}) (bool, error) {
	result, err := expr.Evaluate(vars)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return result.(bool), nil
}

func dfs(g *models.Graph, current string, input map[string]interface{}, visited map[string]bool) map[string]string {
	if visited[current] {
		return nil
	}
	visited[current] = true
	result := make(map[string]string)
	node := g.Nodes[current]
	if node != nil && node.Attributes != nil {
		if res, ok := node.Attributes["result"]; ok {
			if resStr, ok := res.(string); ok {
				result[current] = resStr
			}
		}
	}
	for _, edge := range g.Edges {
		if edge.From == current {
			var ok bool
			var err error
			if edge.Condition != nil {
				ok, err = evalCondition(edge.Condition, input)
				if err != nil {
					continue
				}
			} else {
				ok = true
			}
			if ok {
				subResult := dfs(g, edge.To, input, visited)
				for k, v := range subResult {
					result[k] = v
				}
			}
		}
	}
	return result
}

func EvaluatePolicy(g *models.Graph, input map[string]interface{}) (map[string]string, error) {
	visited := make(map[string]bool)
	res := dfs(g, g.Start, input, visited)
	return res, nil
}
