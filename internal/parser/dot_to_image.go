package parser

import (
	"bytes"
	"context"
	"encoding/base64"

	"github.com/goccy/go-graphviz"
)

func DotToBase64Image(dotString string) (string, error) {
	g, err := graphviz.New(context.Background())

	if err != nil {
		return "", err
	}

	defer g.Close()

	graphAst, err := graphviz.ParseBytes([]byte(dotString))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := g.Render(context.Background(), graphAst, graphviz.PNG, &buf); err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64Str, nil
}
