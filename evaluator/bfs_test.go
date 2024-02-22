package evaluator

import (
	"testing"

	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/parser"
	"github.com/stretchr/testify/assert"
)

type testEdge struct {
	dest, label uint32
	outgoing    bool
}

type testAccessor struct {
	adjacency map[uint32][]testEdge
}

func (t *testAccessor) GetNeighbours(req accessor.GraphAccessRequest) []uint32 {
	res := make([]uint32, 0)
	for _, edge := range t.adjacency[req.Src] {
		if edge.label == req.Label {
			if (edge.outgoing && req.Dir != parser.INCOMING) ||
				(!edge.outgoing && req.Dir != parser.OUTGOING) {
				res = append(res, edge.dest)
			}
		}
	}
	return res
}

func (t *testAccessor) StartQuery(accessor.Algo) int {
	return 1
}

func (t *testAccessor) EndQuery(int) {

}

func (t *testAccessor) GetStats() string {
	return ""
}

var dummyGraph = &testAccessor{
	adjacency: map[uint32][]testEdge{
		1: {{2, 0, true}, {3, 1, true}},
		2: {{3, 2, true}, {1, 0, false}},
		3: {{2, 2, false}, {1, 1, false}},
	},
}

func TestBfs(t *testing.T) {
	bfs := NewBfsEvaluator(dummyGraph)
	query := parser.Query{
		Node:  1,
		Edges: []parser.Edge{{Label: 0, Dir: parser.OUTGOING}, {Label: 2, Dir: parser.BOTH}},
	}
	bfs.Start(query)
	res := bfs.Evaluate()
	bfs.End()
	assert.Equal(t, 1, len(res))
	assert.Equal(t, []uint32{3}, res)

	query = parser.Query{
		Node:  3,
		Edges: []parser.Edge{{Label: 2, Dir: parser.INCOMING}, {Label: 0, Dir: parser.BOTH}},
	}
	bfs.Start(query)
	res = bfs.Evaluate()
	bfs.End()
	assert.Equal(t, 1, len(res))
	assert.Equal(t, []uint32{1}, res)
}
