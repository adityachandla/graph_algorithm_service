package evaluator

import (
	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/parser"
)

type bfsNode struct {
	nodeId uint32
	// This is the index of the label from the original
	// query's label array
	labelIdx int
}

func EvaluateBFS(q parser.Query, access accessor.GraphAccessor) []uint32 {
	res := make([]uint32, 0, 4)
	bfsQueue := NewQueue[bfsNode]()
	edges := q.Edges
	seen := make(map[uint32]struct{})
	bfsQueue.AddToFront(bfsNode{q.Node, 0})
	for !bfsQueue.Empty() {
		toProcess := bfsQueue.PopBack()
		request := accessor.GraphAccessRequest{
			Src:   toProcess.nodeId,
			Label: edges[toProcess.labelIdx].Label,
			Dir:   edges[toProcess.labelIdx].Dir,
		}
		neighbours := access.GetNeighbours(request)
		if toProcess.labelIdx == len(edges)-1 {
			res = append(res, neighbours...)
		} else {
			for _, n := range neighbours {
				if _, ok := seen[n]; !ok {
					bfsQueue.AddToFront(bfsNode{n, toProcess.labelIdx + 1})
					seen[n] = struct{}{}
				}
			}
		}
	}
	return res
}
