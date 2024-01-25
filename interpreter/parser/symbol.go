package parser

// Symbol represents a <Token,string> pair in the codebox.
type Symbol struct {
	token Token
	raw   rune
}

func NewSymbol(raw rune) *Symbol {
	sym := new(Symbol)
	sym.raw = raw
	if found, known := RuneToToken[rune(raw)]; known {
		sym.token = found
	} else {
		sym.token = VAL
	}

	return sym
}

// Is returns whether Token t matches the stored Token.
func (s Symbol) Is(t Token) bool {
	return s.token == t
}

// Raw returns the stored raw character.
func (s Symbol) Raw() rune {
	return s.raw
}
