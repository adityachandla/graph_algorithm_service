package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	qg := QueryGenerator{
		EdgeMap:     map[string]uint32{"one": 1, "two": 2},
		IntervalMap: map[string]Interval{"person": {1, 300}},
	}
	queries := qg.Generate(&QueryStr{"person", []EdgeStr{{"one", BOTH}}}, 10)
	assert.Equal(t, 10, len(queries))
	for _, q := range queries {
		assert.True(t, q.node >= 1 && q.node <= 300)
		assert.Equal(t, 1, len(q.edges))
		assert.Equal(t, BOTH, int(q.edges[0].dir))
		assert.Equal(t, uint32(1), q.edges[0].label)
	}
}
