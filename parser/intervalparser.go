package parser

import (
	"regexp"
	"strconv"
)

type NodeInterval struct {
	Name     string
	Interval Interval
}

type Interval struct {
	Start, End uint32
}

var rangeRegex = regexp.MustCompile(`(\w+)|(\d+)|(\d+)`)

func ParseNodeInterval(line string) (NodeInterval, error) {
	matches := rangeRegex.FindAllString(line, -1)
	if len(matches) != 3 {
		return NodeInterval{}, ParseError
	}
	startUint, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		return NodeInterval{}, ParseError
	}
	endUint, err := strconv.ParseUint(matches[2], 10, 32)
	if err != nil {
		return NodeInterval{}, ParseError
	}
	return NodeInterval{
		Name:     matches[0],
		Interval: Interval{uint32(startUint), uint32(endUint)},
	}, nil
}
