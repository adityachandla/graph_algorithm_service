package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	valid, err := ParseNodeInterval("person|22|88")
	assert.Nil(t, err)
	assert.Equal(t, NodeInterval{"person", Interval{22, 88}}, valid)

	_, err = ParseNodeInterval("|81|91")
	assert.Equal(t, ParseError, err)
}
