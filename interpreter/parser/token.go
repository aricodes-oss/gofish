package parser

//go:generate stringer -type=Token

// Token represents a lexical token.
type Token int

const (
	/** Directional **/
	// Cardinal
	RIGHT Token = iota // >
	LEFT               // <
	DOWN               // v
	UP                 // ^

	// Mirrors
	MIRFORWARD    // /
	MIRBACKWARD   // \
	MIRVERTICAL   // |
	MIRHORIZONTAL // _
	MIRREVERSE    // #
	RANDOM        // x

	// Conditionals
	SKIP   // !
	SKIPIF // ?
	JUMP   // .

	/** Literals/operators **/
	VAL // 0-f
	ADD // +
	SUB // -
	MUL // *
	DIV // ,
	MOD // %
	EQ  // =

	/** Stack Manipulation **/
	DUPLICATE // :
	REMOVE    // ~
	SWAP      // $
	TOPSHIFT  // @
	RSHIFT    // }
	LSHIFT    // {
	REVERSE   // r
	LENGTH    // l
	NEW       // [
	CONSUME   // ]

	/** I/O **/
	OCHAR // o
	ONUM  // n
	READ  // i

	/** Reflection/misc **/
	REGISTER // &
	PEEK     // g
	POKE     // p
	HALT     // :

	WHITESPACE
	EOF
)
