package parser

// Big thank you to the Gopher Academy Blog for this excellent post,
// which I took much of the Scanner type inspiration from
// https://blog.gopheracademy.com/advent-2014/parsers-lexers/

import (
	"bufio"
	"io"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() (rune, error) {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return 0, err
	}
	return ch, nil
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, raw rune) {
	// Read the next rune
	ch, err := s.read()
	if err != nil {
		return EOF, 0
	}
	raw = ch

	// Here is typically when we would consume all whitespace, but
	// in ><> every character is significant - making matching easy
	if found, known := RuneToToken[ch]; known {
		tok = found
	} else {
		tok = VAL
	}

	return
}
