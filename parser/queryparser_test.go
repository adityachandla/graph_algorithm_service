package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatches(t *testing.T) {
	res, err := ParseQuery("Person (>,Knows)(<>,works)(<,Goes)")
	assert.Nil(t, err)
	assert.Equal(t, "Person", res.Name)
	assert.Equal(t, 3, len(res.Edges))
	assert.Equal(t, []EdgeStr{{"Knows", OUTGOING}, {"works", BOTH}, {"Goes", INCOMING}}, res.Edges)
}

func TestErrors(t *testing.T) {
	_, err := ParseQuery("(>,hello)")
	assert.Equal(t, ParseError, err)
	_, err = ParseQuery("Person (>,hello)(><,jello)")
	assert.Equal(t, ParseError, err)
}
