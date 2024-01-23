package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/evaluator"
	"github.com/adityachandla/graph_algorithm_service/parser"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto
var (
	address     = flag.String("address", "localhost:20301", "Address to host")
	nodeMapFile = flag.String("nodeMap", "nodeMap.csv", "The file that contains node to range mapping")
	repetitions = flag.Int("repetitions", 5, "The number of times each query should be repeated")
	algorithm   = flag.String("algorithm", "bfs", "Evaluate with bfs/dfs order")
)

const edgeMapFile = "edgeMap.csv"
const queryFile = "queries.txt"

func main() {
	flag.Parse()
	rand.Seed(17041998)
	graphAccess := accessor.InitializeGraphAccess(*address)
	edgeMap := parser.ParseEdgeLabels(edgeMapFile)
	intervalMap := parser.ParseNodeIntervals(*nodeMapFile)
	queryStrs := parser.ParseQueries(queryFile)
	qg := parser.QueryGenerator{
		EdgeMap:     edgeMap,
		IntervalMap: intervalMap,
	}
	queries := make([]parser.Query, 0, len(queryStrs)*(*repetitions))
	for queryIdx := range queryStrs {
		q := qg.Generate(&queryStrs[queryIdx], *repetitions)
		queries = append(queries, q...)
	}

	var eval evaluator.Interface
	if *algorithm == "bfs" {
		eval = evaluator.NewBfsEvaluator(graphAccess)
	} else if *algorithm == "dfs" {
		eval = evaluator.NewDfsEvaluator(graphAccess)
	}

	for i := range queries {
		start := time.Now()
		res := eval.Evaluate(queries[i])
		diff := time.Now().Sub(start)
		log.Printf("%d results in %v\n", len(res), diff)
	}

}
