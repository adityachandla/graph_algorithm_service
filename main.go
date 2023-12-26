package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adityachandla/graph_algorithm_service/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto
func main() {

	conn, err := grpc.Dial("localhost:20301",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to localhost:20301")
	}
	defer conn.Close()
	client := generated.NewGraphAccessClient(conn)
	request := &generated.AccessRequest{
		NodeId:   22,
		Label:    1,
		Incoming: false,
	}
	response, err := client.GetNeighbours(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Neighbours)
}
