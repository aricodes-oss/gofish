package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
