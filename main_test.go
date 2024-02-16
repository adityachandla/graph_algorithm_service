package main

import (
	"fmt"
	"testing"

	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/evaluator"
	"github.com/adityachandla/graph_algorithm_service/parser"
	"github.com/stretchr/testify/assert"
)

// Test for an empty edge.
func TestSpecificFetch(t *testing.T) {
	graphAccess := accessor.InitializeGraphAccess("localhost:20301")
	eval := evaluator.NewBfsEvaluator(graphAccess)
	query := parser.Query{
		Node:  1,
		Edges: []parser.Edge{{Label: 12, Dir: parser.OUTGOING}},
	}
	res := eval.Evaluate(query)
	fmt.Println(res)
	assert.NotEmpty(t, res)
}
