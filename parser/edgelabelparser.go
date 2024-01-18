package parser

import (
	"regexp"
	"strconv"
)

type EdgeLabel struct {
	Name string
	Id   uint32
}

var edgeLabelRegex = regexp.MustCompile(`(\w+)|(\d+)`)

func ParseEdgeLabel(line string) (EdgeLabel, error) {
	matches := edgeLabelRegex.FindAllString(line, -1)
	if len(matches) != 2 {
		return EdgeLabel{}, ParseError
	}
	id, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		return EdgeLabel{}, ParseError
	}
	return EdgeLabel{
		Name: matches[0],
		Id:   uint32(id),
	}, nil
}
