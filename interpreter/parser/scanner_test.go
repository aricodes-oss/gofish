package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	assert := assert.New(t)
	var tests = []struct {
		name     string
		input    string
		expected []Token
	}{
		{"Fish", "><>", []Token{RIGHT, LEFT, RIGHT}},
		{"Konami", "^^vv<><>", []Token{UP, UP, DOWN, DOWN, LEFT, RIGHT, LEFT, RIGHT}},
		{"Mirrors", "/\\|_#x", []Token{MIRFORWARD, MIRBACKWARD, MIRVERTICAL, MIRHORIZONTAL, MIRREVERSE, RANDOM}},
		{"Conditionals", "!?.", []Token{SKIP, SKIPIF, JUMP}},
		{"Literals", "0123456789abcdef", []Token{VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL, VAL}},
		{"Operands", "+-*,%=()'\"", []Token{ADD, SUB, MUL, DIV, MOD, EQ, LT, GT, QUOTE, DOUBLEQUOTE}},
		{"Stack Ops", ":~$@}{rl[]", []Token{DUPLICATE, REMOVE, SWAP, TOPSHIFT, RSHIFT, LSHIFT, REVERSE, LENGTH, NEW, CONSUME}},
		{"I/O", "oni", []Token{OCHAR, ONUM, READ}},
		{"Reflection", "&gp", []Token{REGISTER, PEEK, POKE}},
		{"Misc", " ;\n", []Token{WHITESPACE, HALT, EOL}},
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
