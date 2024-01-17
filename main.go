package main

import (
	"flag"
	"fmt"

	pb "github.com/adityachandla/graph_algorithm_service/generated"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto
var (
	address     = flag.String("address", "localhost:20301", "Address to host")
	repetitions = flag.Int("repetitions", 5, "Number of times that each query should be repeated")
)

func main() {
	flag.Parse()
	graphAccess := InitializeGraphAccess(*address)
	req := &pb.AccessRequest{
		NodeId:    625,
		Label:     13,
		Direction: pb.AccessRequest_INCOMING,
	}
	fmt.Println(graphAccess.GetNeighbours(req))
}
