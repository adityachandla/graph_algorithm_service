package main

import (
	"flag"
	"log"
	"math/rand"
	"time"
)

//go:generate protoc --go-grpc_out=generated --go_out=generated --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative graph_access.proto
var (
	address     = flag.String("address", "localhost:20301", "Address to host")
	repetitions = flag.Int("repetitions", 5, "Number of times that each query should be repeated")
)

const PERSON_LOW = 0
const PERSON_HIGH = 10294

const POST_LOW = 10295
const POST_HIGH = 1_131_520

func oneHop(graphAccess *GraphAccessor, count int) {
	// Find out where the person is located.
	for i := 0; i < count; i++ {
		start := time.Now()
		personId := PERSON_LOW + (rand.Uint32() % (PERSON_HIGH - PERSON_LOW))
		location := graphAccess.GetNeighbours(personId, 23)
		duration := time.Now().Sub(start).Milliseconds()
		log.Printf("Found %d locations in %d\n", len(location), duration)
	}
}

func twoHop(graphAccess *GraphAccessor, count int) {
	// Find friends of friends.
	for i := 0; i < count; i++ {
		start := time.Now()
		personId := PERSON_LOW + (rand.Uint32() % (PERSON_HIGH - PERSON_LOW))
		friendsOfFriends := make(map[uint32]struct{})
		friends := graphAccess.GetNeighbours(personId, 21)
		for idx := range friends {
			twoHopNeighbours := graphAccess.GetNeighbours(friends[idx], 21)
			for _, v := range twoHopNeighbours {
				friendsOfFriends[v] = struct{}{}
			}
		}
		duration := time.Now().Sub(start).Milliseconds()
		log.Printf("Found %d friends of friends in %d\n", len(friendsOfFriends), duration)
	}
}

func threeHop(graphAccess *GraphAccessor, count int) {
	// What places did a post creator study in.
	for i := 0; i < count; i++ {
		start := time.Now()
		postId := POST_LOW + (rand.Uint32() % (POST_HIGH - POST_LOW))
		creator := graphAccess.GetNeighbours(postId, 10)[0]
		places := make(map[uint32]struct{})
		universities := graphAccess.GetNeighbours(creator, 4)
		for _, uni := range universities {
			uniPlace := graphAccess.GetNeighbours(uni, 18)[0]
			places[uniPlace] = struct{}{}
		}
		duration := time.Now().Sub(start).Milliseconds()
		log.Printf("Found %d places in %d\n", len(places), duration)
	}
}

func main() {
	flag.Parse()
	graphAccess := InitializeGraphAccess(*address)
	// These are dfs type queries
	oneHop(graphAccess, *repetitions)
	twoHop(graphAccess, *repetitions)
	threeHop(graphAccess, *repetitions)
}
