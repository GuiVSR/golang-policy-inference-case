package parser

import (
	"lab/internal/models"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/awalterschulze/gographviz"
	"github.com/awalterschulze/gographviz/ast"
	"github.com/goccy/go-graphviz/cgraph"
)

func ParsePolicy(dotString string) (*models.Graph, error) {
	rawGraph, err := gographviz.ParseString(dotString)
	if err != nil {
		return nil, err
	}
	graph := BuildGraph(rawGraph)
	return graph, nil
}

func parseNode(rawNode *ast.NodeStmt) *models.Node {
	attrs := make(map[string]interface{})

	for _, rawAttrs := range rawNode.Attrs {
		for _, rawAttr := range rawAttrs {
			sanitized := strings.Trim(rawAttr.Value.String(), `\"`)
			sanitized = strings.Trim(sanitized, ` `)
			hasAttributes := strings.Contains(sanitized, "=")

			if hasAttributes {
				splitted := strings.Split(sanitized, `,`)
				for newAttr := range splitted {
					kv := strings.Split(splitted[newAttr], `=`)
					attrs[kv[0]] = kv[1]
				}
			}
		}
	}

	return &models.Node{
		Name:       rawNode.NodeID.ID.String(),
		Attributes: attrs,
	}
}

func parseEdge(rawEdge *ast.EdgeStmt) *models.Edge {
	destID := string(rawEdge.EdgeRHS[0].Destination.GetID())
	condition := parseCondition(rawEdge)

	return &models.Edge{
		From:      rawEdge.Source.GetID().String(),
		To:        destID,
		Condition: condition,
	}
}

func BuildGraph(graph *ast.Graph) *models.Graph {
	nodes := make(map[string]*models.Node)
	statementList := graph.StmtList
	var startNode string

	nodeStmt, ok := statementList[0].(*ast.NodeStmt)
	if ok {
		startNode = string(nodeStmt.NodeID.ID)
	}

	for i := 0; i < len(statementList); i++ {
		nodeStmt, isNode := statementList[i].(*ast.NodeStmt)

		if isNode {
			nodes[string(nodeStmt.NodeID.ID)] = parseNode(nodeStmt)
		} else {
			if edgeStmt, ok := statementList[i].(*ast.EdgeStmt); ok {
				sourceID := string(edgeStmt.Source.GetID())
				newEdge := parseEdge(edgeStmt)

				if nodes[sourceID].Edges == nil {
					nodes[sourceID].Edges = []*models.Edge{newEdge}
				} else {
					nodes[sourceID].Edges = append(nodes[sourceID].Edges, newEdge)
				}
			}
		}
	}

	graphModel := models.Graph{
		Nodes: nodes,
		Start: startNode,
	}

	return &graphModel
}

func parseCondition(edgeStmt *ast.EdgeStmt) *govaluate.EvaluableExpression {
	if edgeStmt.Attrs != nil {
		for _, attrList := range edgeStmt.Attrs {
			for _, attr := range attrList {
				if strings.Contains(attr.String(), "cond") {
					exprStr := strings.TrimPrefix(attr.String(), "cond=")
					exprStr = strings.Trim(exprStr, `\"`)
					condition, err := govaluate.NewEvaluableExpression(exprStr)
					if err != nil {
						return nil
					}
					return condition
				}
			}
		}
	}
	return nil
}

func BuildRenderGraph(dotString string) (*cgraph.Graph, error) {
	renderGraph := gographviz.NewGraph()
	rawGraph, _ := gographviz.ParseString(dotString)
	statementList := rawGraph.StmtList

	for _, stmt := range statementList {
		if nodeStmt, ok := stmt.(*ast.NodeStmt); ok {
			nodeID := string(nodeStmt.NodeID.ID)
			attrs := make(map[string]string)
			if nodeStmt.Attrs != nil {
				for _, attr := range nodeStmt.Attrs {
					attrs[attr[0].Field.String()] = attr[0].Value.String()
				}
			}
			renderGraph.AddNode("G", nodeID, attrs)
		}

		if edgeStmt, ok := stmt.(*ast.EdgeStmt); ok {
			sourceID := string(edgeStmt.Source.GetID())
			destID := string(edgeStmt.EdgeRHS[0].Destination.GetID())
			attrs := make(map[string]string)
			if edgeStmt.Attrs != nil {
				for _, attrList := range edgeStmt.Attrs {
					for _, attr := range attrList {
						attrs[attr.Field.String()] = attr.Value.String()
					}
				}
			}
			renderGraph.AddEdge(sourceID, destID, true, attrs)
		}
	}

	dotOutput := renderGraph.String()
	renderableGraph, _ := cgraph.ParseBytes([]byte(dotOutput))

	return renderableGraph, nil
}
