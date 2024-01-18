package parser

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/adityachandla/graph_algorithm_service/fileutil"
)

type Query struct {
	Node  uint32
	Edges []Edge
}

type Edge struct {
	Label uint32
	Dir   Direction
}

type QueryGenerator struct {
	EdgeMap     map[string]uint32
	IntervalMap map[string]Interval
}

func (gen *QueryGenerator) Generate(queryStr *QueryStr, repetitions int) []Query {
	//All queries should share the same edge list.
	edges := make([]Edge, len(queryStr.Edges))
	for i, e := range queryStr.Edges {
		labelUint, ok := gen.EdgeMap[e.Label]
		if !ok {
			panic(fmt.Sprintf("%s not found in edgeMap", e.Label))
		}
		edges[i] = Edge{labelUint, e.Direction}
	}
	res := make([]Query, repetitions)
	interval, ok := gen.IntervalMap[queryStr.Name]
	if !ok {
		panic(fmt.Sprintf("%s not found in intervalMap", queryStr.Name))
	}
	for i := range res {
		nodeId := uint32(rand.Int63n(int64(interval.End-interval.Start+1))) + interval.Start
		res[i].Node = nodeId
		res[i].Edges = edges
	}
	return res
}

func ParseQueries(filename string) []QueryStr {
	r := fileutil.NewReader(filename)
	res := make([]QueryStr, 0, 8)
	line, err := r.ReadLine()
	for err == nil {
		var query QueryStr
		query, err = ParseQuery(line)
		if err == nil {
			res = append(res, query)
		}
		line, err = r.ReadLine()
	}
	return res
}

func ParseEdgeLabels(filename string) map[string]uint32 {
	r := fileutil.NewReader(filename)
	res := make(map[string]uint32)
	//Discard the header
	_, err := r.ReadLine()
	line, err := r.ReadLine()
	for err == nil {
		var edgeLabel EdgeLabel
		edgeLabel, err = ParseEdgeLabel(line)
		if err != nil {
			panic(err)
		}
		res[strings.ToUpper(edgeLabel.Name)] = edgeLabel.Id
		line, err = r.ReadLine()
	}
	return res
}

func ParseNodeIntervals(filename string) map[string]Interval {
	r := fileutil.NewReader(filename)
	res := make(map[string]Interval)
	//Discard the header
	_, err := r.ReadLine()
	line, err := r.ReadLine()
	for err == nil {
		var nodeInterval NodeInterval
		nodeInterval, err = ParseNodeInterval(line)
		if err != nil {
			panic(err)
		}
		res[strings.ToUpper(nodeInterval.Name)] = nodeInterval.Interval
		line, err = r.ReadLine()
	}
	return res
}
