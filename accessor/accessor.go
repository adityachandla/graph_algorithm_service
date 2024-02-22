package accessor

import (
	"context"
	"log"

	pb "github.com/adityachandla/graph_algorithm_service/generated"
	"github.com/adityachandla/graph_algorithm_service/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GraphAccessRequest struct {
	Src, Label uint32
	Dir        parser.Direction
	QueryId    int
}

type GraphAccessor interface {
	StartQuery(Algo) int
	GetNeighbours(request GraphAccessRequest) []uint32
	EndQuery(int)
	GetStats() string
}

type Algo byte

const (
	BFS Algo = iota
	DFS
)

type GrpcGraphAccessor struct {
	conn   *grpc.ClientConn
	client pb.GraphAccessClient
}

func InitializeGraphAccess(address string) *GrpcGraphAccessor {
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to localhost:20301")
	}
	client := pb.NewGraphAccessClient(conn)
	return &GrpcGraphAccessor{conn, client}
}

func (g *GrpcGraphAccessor) GetNeighbours(request GraphAccessRequest) []uint32 {
	req := &pb.AccessRequest{
		NodeId:  request.Src,
		Label:   request.Label,
		QueryId: int32(request.QueryId),
	}
	if request.Dir == parser.OUTGOING {
		req.Direction = pb.AccessRequest_OUTGOING
	} else if request.Dir == parser.INCOMING {
		req.Direction = pb.AccessRequest_INCOMING
	} else {
		req.Direction = pb.AccessRequest_BOTH
	}
	response, err := g.client.GetNeighbours(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return response.Neighbours
}

func (g *GrpcGraphAccessor) StartQuery(algorithm Algo) int {
	req := pb.StartQueryRequest{}
	if algorithm == BFS {
		req.Algorithm = pb.StartQueryRequest_BFS
	} else {
		req.Algorithm = pb.StartQueryRequest_DFS
	}
	resp, err := g.client.StartQuery(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	return int(resp.QueryId)
}

func (g *GrpcGraphAccessor) EndQuery(id int) {
	req := pb.EndQueryRequest{QueryId: int32(id)}
	_, err := g.client.EndQuery(context.Background(), &req)
	if err != nil {
		panic(err)
	}
}

func (g *GrpcGraphAccessor) GetStats() string {
	req := &pb.StatsRequest{}
	response, err := g.client.GetStats(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return response.Stats
}

func (g *GrpcGraphAccessor) Close() {
	err := g.conn.Close()
	if err != nil {
		panic(err)
	}
}
