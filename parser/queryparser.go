package parser

import (
	"fmt"
	"regexp"
)

type Direction byte

const (
	OUTGOING = iota
	INCOMING
	BOTH
)

func (d Direction) String() string {
	switch d {
	case OUTGOING:
		return "outgoing"
	case INCOMING:
		return "incoming"
	case BOTH:
		return "both"
	}
	panic("Invalid number")
}

type QueryStr struct {
	Name  string
	Edges []EdgeStr
}

type EdgeStr struct {
	Label     string
	Direction Direction
}

var SOURCE_FORMAT = regexp.MustCompile(`^(\w+)`)
var EDGE_FORMAT = regexp.MustCompile(`\(([<>]{1,2}),\s?(\w+)\)`)

var ParseError = fmt.Errorf("Unable to parse string")

// The format of the query is as follows:
func ParseQuery(input string) (QueryStr, error) {
	res := QueryStr{}
	src := SOURCE_FORMAT.FindAllString(input, -1)
	if len(src) == 0 {
		return res, ParseError
	}
	res.Name = src[0]
	edges := EDGE_FORMAT.FindAllStringSubmatch(input, -1)
	if len(edges) == 0 {
		return res, ParseError
	}
	res.Edges = make([]EdgeStr, len(edges))
	for i := 0; i < len(edges); i++ {
		if len(edges[i]) != 3 {
			return res, ParseError
		}
		res.Edges[i].Label = edges[i][2]
		switch edges[i][1] {
		case ">":
			res.Edges[i].Direction = OUTGOING
		case "<":
			res.Edges[i].Direction = INCOMING
		case "<>":
			res.Edges[i].Direction = BOTH
		default:
			return res, ParseError
		}
	}
	return res, nil
}
