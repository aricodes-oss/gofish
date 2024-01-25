package parser

var RuneToToken = map[rune]Token{
	// Cardinal
	'>': RIGHT,
	'<': LEFT,
	'v': DOWN,
	'^': UP,

	// Mirrors
	'/':  MIRFORWARD,
	'\\': MIRBACKWARD,
	'|':  MIRVERTICAL,
	'_':  MIRHORIZONTAL,
	'#':  MIRREVERSE,
	'x':  RANDOM,

	// Conditionals
	'!': SKIP,
	'?': SKIPIF,
	'.': JUMP,

	// Literals/operands
	'+':  ADD,
	'-':  SUB,
	'*':  MUL,
	',':  DIV,
	'%':  MOD,
	'=':  EQ,
	'(':  LT,
	')':  GT,
	'\'': QUOTE,
	'"':  DOUBLEQUOTE,

	// Stack manipulation
	':': DUPLICATE,
	'~': REMOVE,
	'$': SWAP,
	'@': TOPSHIFT,
	'}': RSHIFT,
	'{': LSHIFT,
	'r': REVERSE,
	'l': LENGTH,
	'[': NEW,
	']': CONSUME,

	// I/O
	'o': OCHAR,
	'n': ONUM,
	'i': READ,

	// Reflection/misc
	'&':  REGISTER,
	'g':  PEEK,
	'p':  POKE,
	' ':  WHITESPACE,
	'\n': EOL,
	';':  HALT,
}
