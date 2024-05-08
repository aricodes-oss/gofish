package codebox

import (
	"gofish/interpreter/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodebox(t *testing.T) {
	type cb = [][]parser.Token
	assert := assert.New(t)

	var tests = []struct {
		name     string
		input    string
		expected cb
	}{
		{"cycler", "./test_data/cycler", cb{{parser.RIGHT, parser.DOWN}, {parser.UP, parser.LEFT}}},
		{"mirror shield", "./test_data/mirror_shield", cb{{parser.MIRFORWARD, parser.MIRHORIZONTAL, parser.MIRBACKWARD}, {parser.MIRVERTICAL, parser.MIRREVERSE, parser.MIRVERTICAL}, {parser.WHITESPACE, parser.RANDOM, parser.WHITESPACE}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := NewCodebox()

			cb.BuildCodebox(tt.input)
			assert.Equal(tt.expected, cb.grid)
		})
	}
}
