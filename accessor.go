package main

import (
	"context"
	"log"

	"github.com/adityachandla/graph_algorithm_service/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GraphAccessor struct {
	conn   *grpc.ClientConn
	client generated.GraphAccessClient
}

func InitializeGraphAccess(address string) *GraphAccessor {
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to localhost:20301")
	}
	client := generated.NewGraphAccessClient(conn)
	return &GraphAccessor{conn, client}
}

func (g *GraphAccessor) GetNeighbours(node, label uint32) []uint32 {
	request := &generated.AccessRequest{
		NodeId:   node,
		Label:    label,
		Incoming: false,
	}
	response, err := g.client.GetNeighbours(context.Background(), request)
	if err != nil {
		panic(err)
	}
	return response.Neighbours
}

func (g *GraphAccessor) Close() {
	g.conn.Close()
}
