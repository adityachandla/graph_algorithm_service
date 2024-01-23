package evaluator

import (
	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/parser"
)

type dfsNode struct {
	nodeId uint32
	// This is the index of the label from the original
	// query's label array
	labelIdx int
}

type DFSEvaluator struct {
	access accessor.GraphAccessor
	result []uint32
	stack  *Stack[dfsNode]
	seen   map[uint32]struct{}
	edges  []parser.Edge
}

func NewDfsEvaluator(access accessor.GraphAccessor) *DFSEvaluator {
	return &DFSEvaluator{access: access}
}

func (eval *DFSEvaluator) initialize(q parser.Query) {
	eval.result = make([]uint32, 0, 4)
	eval.stack = NewStack[dfsNode]()
	eval.edges = q.Edges
	eval.seen = make(map[uint32]struct{})
	eval.stack.Push(dfsNode{q.Node, 0})
}

func (eval *DFSEvaluator) Evaluate(q parser.Query) []uint32 {
	eval.initialize(q)
	for !eval.stack.Empty() {
		toProcess := eval.stack.Pop()
		eval.processNode(toProcess)
	}
	return eval.result
}

func (eval *DFSEvaluator) processNode(toProcess dfsNode) {
	labelIdx := toProcess.labelIdx
	request := accessor.GraphAccessRequest{
		Src:   toProcess.nodeId,
		Label: eval.edges[labelIdx].Label,
		Dir:   eval.edges[labelIdx].Dir,
	}
	neighbours := eval.access.GetNeighbours(request)
	if labelIdx == len(eval.edges)-1 {
		eval.result = append(eval.result, neighbours...)
	} else {
		eval.addToQueue(neighbours, labelIdx+1)
	}
}

func (eval *DFSEvaluator) addToQueue(neighbours []uint32, idx int) {
	for _, n := range neighbours {
		if _, ok := eval.seen[n]; !ok {
			eval.stack.Push(dfsNode{n, idx})
			eval.seen[n] = struct{}{}
		}
	}
}
