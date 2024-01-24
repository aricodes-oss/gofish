package errors

import "errors"

// ><> has 4 different types of errors, but we are only concerned with these
// Notably the ><> spec states that the only "acceptable" error message is
// "something smells fishy..." - we'll follow that, but use unique exit codes
// to identify the actual problem.
var (
	ErrDivByZero          = errors.New("division by zero")
	ErrInvalidInstruction = errors.New("invalid instruction")
	ErrStackEmpty         = errors.New("stack is empty")
)
