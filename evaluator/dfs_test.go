package evaluator

import (
	"testing"

	"github.com/adityachandla/graph_algorithm_service/parser"
	"github.com/stretchr/testify/assert"
)

func TestDfs(t *testing.T) {
	dfs := NewDfsEvaluator(dummyGraph)
	query := parser.Query{
		Node:  1,
		Edges: []parser.Edge{{Label: 0, Dir: parser.OUTGOING}, {Label: 2, Dir: parser.BOTH}},
	}
	res := dfs.Evaluate(query)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, []uint32{3}, res)

	query = parser.Query{
		Node:  3,
		Edges: []parser.Edge{{Label: 2, Dir: parser.INCOMING}, {Label: 0, Dir: parser.BOTH}},
	}
	res = dfs.Evaluate(query)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, []uint32{1}, res)
}
