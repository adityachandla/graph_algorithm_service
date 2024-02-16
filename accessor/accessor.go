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
}

type GraphAccessor interface {
	GetNeighbours(request GraphAccessRequest) []uint32
	GetStats() string
}

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
		NodeId: request.Src,
		Label:  request.Label,
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

func (g *GrpcGraphAccessor) GetStats() string {
	req := &pb.StatsRequest{}
	response, err := g.client.GetStats(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return response.Stats
}

func (g *GrpcGraphAccessor) Close() {
	g.conn.Close()
}
