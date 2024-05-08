package parser

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	type ts = []Token
	assert := assert.New(t)

	var tests = []struct {
		name     string
		input    string
		expected ts
	}{
		{"Fish", "><>", ts{RIGHT, LEFT, RIGHT}},
		{"Konami", "^^vv<><>", ts{UP, UP, DOWN, DOWN, LEFT, RIGHT, LEFT, RIGHT}},
		{"Mirrors", "/\\|_#x", ts{MIRFORWARD, MIRBACKWARD, MIRVERTICAL, MIRHORIZONTAL, MIRREVERSE, RANDOM}},
		{"Conditionals", "!?.", ts{SKIP, SKIPIF, JUMP}},
		{"Literals", "0123456789abcdef", ts{VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL}},
		{"Operands", "+-*,%=()'\"", ts{ADD, SUB, MUL, DIV, MOD, EQ, LT, GT, QUOTE, DOUBLEQUOTE}},
		{"Stack Ops", ":~$@}{rl[]", ts{DUPLICATE, REMOVE, SWAP, TOPSHIFT, RSHIFT, LSHIFT, REVERSE, LENGTH, NEW, CONSUME}},
		{"I/O", "oni", ts{OCHAR, ONUM, READ}},
		{"Reflection", "&gp", ts{REGISTER, PEEK, POKE}},
		{"Misc", " ;\n", ts{WHITESPACE, HALT, EOL}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))

			for _, expected := range tt.expected {
				received, _ := scanner.Scan()
				assert.Equal(expected, received)
			}
		})
	}
}

func TestNoUnintendedDuplilcateTokens(t *testing.T) {
	assert := assert.New(t)

	for key, value := range RuneToToken {
		t.Run(fmt.Sprintf("%c Doesn't match EOF", key), func(t *testing.T) {
			for o_key, o_value := range RuneToToken {
				if key == o_key {
					// If we're comparing against ourself, we should match
					assert.Equal(value, o_value)
				} else if slices.Contains([]rune{'\'', '"'}, key) && slices.Contains([]rune{'\'', '"'}, o_key) {
					// Since QUOTE and DOUBLEQUOTE are intentionally equal, make sure they match.
					assert.Equal(value, o_value)
				} else {
					// In any other case the constants should be different.
					assert.NotEqual(value, o_value)
				}
			}
		})
	}
}
