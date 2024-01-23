package evaluator

import (
	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/parser"
)

type Interface interface {
	Evaluate(q parser.Query) []uint32
}

type bfsNode struct {
	nodeId uint32
	// This is the index of the label from the original
	// query's label array
	labelIdx int
}

type nodeLevel struct {
	nodeId uint32
	idx    int
}

type BFSEvaluator struct {
	access accessor.GraphAccessor
	result []uint32
	queue  *Queue[bfsNode]
	seen   map[nodeLevel]struct{}
	edges  []parser.Edge
}

func NewBfsEvaluator(access accessor.GraphAccessor) *BFSEvaluator {
	return &BFSEvaluator{access: access}
}

func (eval *BFSEvaluator) initialize(q parser.Query) {
	eval.result = make([]uint32, 0, 4)
	eval.queue = NewQueue[bfsNode]()
	eval.edges = q.Edges
	eval.seen = make(map[nodeLevel]struct{})
	eval.queue.AddToFront(bfsNode{q.Node, 0})
}

func (eval *BFSEvaluator) Evaluate(q parser.Query) []uint32 {
	eval.initialize(q)
	for !eval.queue.Empty() {
		toProcess := eval.queue.PopBack()
		eval.processNode(toProcess)
	}
	return eval.result
}

func (eval *BFSEvaluator) processNode(toProcess bfsNode) {
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

func (eval *BFSEvaluator) addToQueue(neighbours []uint32, idx int) {
	for _, n := range neighbours {
		if _, ok := eval.seen[nodeLevel{n, idx}]; !ok {
			eval.queue.AddToFront(bfsNode{n, idx})
			eval.seen[nodeLevel{n, idx}] = struct{}{}
		}
	}
}
