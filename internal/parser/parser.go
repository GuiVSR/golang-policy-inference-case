package parser

import (
	"github.com/awalterschulze/gographviz"
	"github.com/awalterschulze/gographviz/ast"
)

func ParsePolicy(dotString string) (*ast.Graph, error) {
	g, err := gographviz.ParseString(dotString)
	if err != nil {
		return nil, err
	}
	return g, nil
}
