package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/adityachandla/graph_algorithm_service/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto

var (
	address = flag.String("address", "localhost:20301", "Address to host")
)

type graphAccess struct {
	conn   *grpc.ClientConn
	client generated.GraphAccessClient
}

func initializeGraphAccess() *graphAccess {
	conn, err := grpc.Dial(*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to localhost:20301")
	}
	client := generated.NewGraphAccessClient(conn)
	return &graphAccess{conn, client}
}

func (g *graphAccess) GetNeighbours(node, label uint32) []uint32 {
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

func (g *graphAccess) Close() {
	g.conn.Close()
}

func shortOne(graphAccess *graphAccess, count int) {
	//Find out where the person is located
	for i := 0; i < count; i++ {
		start := time.Now()
		personId := rand.Uint32() % 10294
		location := graphAccess.GetNeighbours(personId, 23)
		if len(location) > 1 {
			fmt.Println("More than one locations")
		}
		fmt.Printf("Found person at: %d\n", location[0])
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
}

func shortThree(graphAccess *graphAccess, count int) {
	//Find a person's friends
	for i := 0; i < count; i++ {
		start := time.Now()
		personId := rand.Uint32() % 10294
		friends := graphAccess.GetNeighbours(personId, 21)
		fmt.Printf("%d has %d friends\n", personId, len(friends))
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
}

func shortFive(graphAccess *graphAccess, count int) {
	//Creator of a message
	message_low := uint32(1_131_521)
	message_high := uint32(2_870_958)
	for i := 0; i < count; i++ {
		start := time.Now()
		messageId := message_low + (rand.Uint32() % (message_high - message_low + 1))
		creator := graphAccess.GetNeighbours(messageId, 6)[0]
		fmt.Printf("%d was created by %d\n", messageId, creator)
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
}

func shortSix(graphAccess *graphAccess, count int) {
	//Get the location of a person's companies
	for i := 0; i < count; i++ {
		start := time.Now()
		personId := rand.Uint32() % 10294
		companies := graphAccess.GetNeighbours(personId, 3)
		for _, cid := range companies {
			location := graphAccess.GetNeighbours(cid, 18)
			fmt.Printf("Person %d has worked in %d\n", personId, location[0])
		}
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
}

func main() {
	flag.Parse()
	graphAccess := initializeGraphAccess()
	shortOne(graphAccess, 10)
	shortThree(graphAccess, 10)
	shortFive(graphAccess, 10)
	shortSix(graphAccess, 10)
}
