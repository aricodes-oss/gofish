package interpreter

import "errors"

// ><> has 4 different types of errors, but we are only concerned with these
// Notably the ><> spec states that the only "acceptable" error message is
// "something smells fishy..." - we'll follow that, but use unique exit codes
// to identify the actual problem.
var (
	ErrDivByZero          = errors.New("Division by zero")
	ErrInvalidInstruction = errors.New("Invalid instruction")
	ErrStackEmpty         = errors.New("Stack is empty")
)
