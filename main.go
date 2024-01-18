package main

import (
	"flag"
	"fmt"

	"github.com/adityachandla/graph_algorithm_service/parser"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto
var (
	address     = flag.String("address", "localhost:20301", "Address to host")
	nodeMapFile = flag.String("nodeMap", "nodeMap.csv", "The file that contains node to range mapping")
)

const edgeMapFile = "edgeMap.csv"
const queryFile = "queries.txt"

func main() {
	flag.Parse()
	//graphAccess := InitializeGraphAccess(*address)
	edgeMap := parser.ParseEdgeLabels(edgeMapFile)
	intervalMap := parser.ParseNodeIntervals(*nodeMapFile)
	queries := parser.ParseQueries(queryFile)
	fmt.Println(edgeMap)
	fmt.Println(intervalMap)
	fmt.Println(queries)
}
