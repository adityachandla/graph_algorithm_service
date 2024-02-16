package main

import (
	"flag"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/adityachandla/graph_algorithm_service/accessor"
	"github.com/adityachandla/graph_algorithm_service/evaluator"
	"github.com/adityachandla/graph_algorithm_service/parser"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto
var (
	address     = flag.String("address", "localhost:20301", "Address to host")
	nodeMapFile = flag.String("nodeMap", "nodeMap10.csv", "The file that contains node to range mapping")
	repetitions = flag.Int("repetitions", 5, "The number of times each query should be repeated")
	algorithm   = flag.String("algorithm", "bfs", "Evaluate with bfs/dfs order")
	parallelism = flag.Int("parallelism", 1, "Number of queries to evaluate in parallel")
	queryFile   = flag.String("query", "queries.txt", "File with queries")
)

const edgeMapFile = "edgeMap.csv"

func main() {
	flag.Parse()
	rand.Seed(17041998)
	queries := generateQueries()
	for _, queryArray := range queries {
		evaluateQueries(queryArray)
	}
}

func generateQueries() [][]parser.Query {
	edgeMap := parser.ParseEdgeLabels(edgeMapFile)
	intervalMap := parser.ParseNodeIntervals(*nodeMapFile)
	queryStrs := parser.ParseQueries(*queryFile)
	qg := parser.QueryGenerator{
		EdgeMap:     edgeMap,
		IntervalMap: intervalMap,
	}
	queries := make([][]parser.Query, len(queryStrs))
	for queryIdx := range queryStrs {
		q := qg.Generate(&queryStrs[queryIdx], queryIdx, *repetitions)
		queries[queryIdx] = q
	}
	return queries
}

func evaluateQueries(queries []parser.Query) {
	graphAccess := accessor.InitializeGraphAccess(*address)

	queryChannel := make(chan parser.Query)
	var wg sync.WaitGroup
	for i := 0; i < *parallelism; i++ {
		var eval evaluator.Interface
		if *algorithm == "bfs" {
			eval = evaluator.NewBfsEvaluator(graphAccess)
		} else if *algorithm == "dfs" {
			eval = evaluator.NewDfsEvaluator(graphAccess)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			runQuery(queryChannel, eval)
		}()
	}
	for i := range queries {
		queryChannel <- queries[i]
	}
	close(queryChannel)
	wg.Wait()
	log.Println(graphAccess.GetStats())
}

func runQuery(queryChannel <-chan parser.Query, eval evaluator.Interface) {
	for q := range queryChannel {
		start := time.Now()
		res := eval.Evaluate(q)
		diff := time.Now().Sub(start)
		log.Printf("QueryId=%d Results=%d time=%dms\n", q.Id, len(res), diff.Milliseconds())
	}
}
