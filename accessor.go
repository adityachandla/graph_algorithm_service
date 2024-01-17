package main

import (
	"context"
	"log"

	pb "github.com/adityachandla/graph_algorithm_service/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GraphAccessor struct {
	conn   *grpc.ClientConn
	client pb.GraphAccessClient
}

func InitializeGraphAccess(address string) *GraphAccessor {
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to localhost:20301")
	}
	client := pb.NewGraphAccessClient(conn)
	return &GraphAccessor{conn, client}
}

func (g *GraphAccessor) GetNeighbours(req *pb.AccessRequest) []uint32 {
	response, err := g.client.GetNeighbours(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return response.Neighbours
}

func (g *GraphAccessor) Close() {
	g.conn.Close()
}
