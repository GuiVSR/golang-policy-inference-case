package parser

import (
	"strings"
	"testing"

	"github.com/awalterschulze/gographviz/ast"
)

func TestParsePolicy_Valid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		check func(t *testing.T, g *ast.Graph)
	}{
		{
			name:  "ExampleDot",
			input: `digraph Policy { start [result=""]; approved [ result="approved=true,segment=prime" ]; rejected [ result="approved=false" ]; review [ result="approved=false,segment=manual" ]; start -> approved [cond="age>=18 && score>700"]; start -> review [cond="age>=18 && score<=700"]; start -> rejected [cond="age<18"]; }`,
			check: func(t *testing.T, g *ast.Graph) {
				output := g.String()
				if !strings.Contains(output, `result="approved=true,segment=prime"`) {
					t.Fatalf("approved node attributes not preserved")
				}
				if !strings.Contains(output, `cond="age>=18 && score>700"`) {
					t.Fatalf("edge cond not preserved")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g, err := ParsePolicy(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if g == nil {
				t.Fatalf("expected graph, got nil")
			}
			tc.check(t, g)
		})
	}
}

func TestParsePolicy_Invalid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "IncompleteDot",
			input: "digraph { start -> ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParsePolicy(tc.input)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
