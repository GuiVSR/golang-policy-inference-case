package parser

import (
	"lab/internal/models"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/awalterschulze/gographviz"
	"github.com/awalterschulze/gographviz/ast"
)

func ParsePolicy(dotString string) (*models.Graph, error) {
	rawGraph, err := gographviz.ParseString(dotString)
	if err != nil {
		return nil, err
	}
	graph := BuildGraph(rawGraph)
	return graph, nil
}

func parseResult(resultStr string) *models.Node {
	attrs := make(map[string]interface{})
	resultStr = strings.Trim(resultStr, "\"")
	if resultStr == "" {
		return &models.Node{Attributes: attrs}
	}

	pairs := strings.Split(resultStr, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(strings.TrimSpace(pair), "=", 2)
		if len(kv) == 2 {
			attrs[kv[0]] = kv[1]
		}
	}
	return &models.Node{Attributes: attrs}
}

func BuildGraph(graph *ast.Graph) *models.Graph {
	nodes := make(map[string]*models.Node)
	statementList := graph.StmtList
	var edges []*models.Edge
	var startNode string

	if nodeStmt, ok := statementList[0].(*ast.NodeStmt); ok {
		startNode = string(nodeStmt.NodeID.ID)
	}

	for i := 1; i < len(statementList); i++ {
		nodeStmt, ok := statementList[i].(*ast.NodeStmt)

		if ok && nodeStmt.Attrs != nil {
			for _, attr := range nodeStmt.Attrs {
				temp := attr.String()
				nodes[string(nodeStmt.NodeID.ID)] = parseResult(temp)
			}
		}

		edgeStmt, ok := statementList[i].(*ast.EdgeStmt)

		if ok && len(edgeStmt.EdgeRHS) > 0 {
			sourceID := string(edgeStmt.Source.GetID())
			destID := string(edgeStmt.EdgeRHS[0].Destination.GetID())
			var condition *govaluate.EvaluableExpression

			if edgeStmt.Attrs != nil {
				for _, attrList := range edgeStmt.Attrs {
					for _, attr := range attrList {
						if strings.Contains(attr.String(), "cond") {
							exprStr := strings.TrimPrefix(attr.String(), "cond=")
							exprStr = strings.Trim(exprStr, `"'`)
							condition, _ = govaluate.NewEvaluableExpression(exprStr)
						}
					}
				}
			}

			newEdge := &models.Edge{
				From:      sourceID,
				To:        destID,
				Condition: condition,
			}
			edges = append(edges, newEdge)
		}

	}

	graphModel := models.Graph{
		Nodes: nodes,
		Edges: edges,
		Start: startNode,
	}

	return &graphModel
}
