package stack

import (
	"gofish/interpreter/errors"
	"slices"
)

// A Stack is an object capable of performing stack operations.
type Stack interface {
	// Push adds a value to the Stack.
	Push(val float64)
	// PushN adds multiple values to the Stack.
	PushN(vals ...float64)

	// Pop removes the last value from the Stack and returns it.
	Pop() (float64, error)
	// PopN removes the last N values from the Stack and returns it.
	PopN(n int) ([]float64, error)
	// PopAll removes all values from the Stack and returns them.
	PopAll() ([]float64, error)

	// Duplicate adds another copy of the top value to the Stack.
	Duplicate() error
	// Swap switches the places of the top two values on the Stack.
	Swap() error

	// Rshift moves the entire stack to the right ([1,2,3,4] -> [4,1,2,3]).
	Rshift() error
	// Lshift moves the entire stack to the right ([1,2,3,4] -> [2,3,4,1]).
	Lshift() error
	// TopShift takes the top three values and shifts them right ([1,2,3,4] -> [1,4,2,3]).
	TopShift() error

	// Reverse reverse the order of the Stack.
	Reverse()
	// Length returns the size of the stack.
	Length() int

	// New returns a new stack with `x` elements off the previous one.
	// `x` is popped off the stack.
	New() (Stack, error)
	// Consume takes a new Stack and appends its elements to this one.
	Consume(child Stack)

	// Register toggles an element into or out of the register.
	Register() error

	// Empty returns whether or not the pool is empty.
	Empty() bool
	// Clear empties the stack.
	Clear()
}

// A stack holds float64s as well as a single-element register.
type stack struct {
	pool     []float64
	register float64

	registerInUse bool // An unfortunate consequence of Go's zero values
}

// NewStack returns a new Stack object.
func NewStack() Stack {
	return new(stack)
}

func (s *stack) Push(val float64) {
	s.pool = append(s.pool, val)
}

func (s *stack) PushN(vals ...float64) {
	s.pool = append(s.pool, vals...)
}

func (s *stack) Pop() (val float64, err error) {
	if s.Empty() {
		err = errors.ErrStackEmpty
		return
	}

	last := len(s.pool) - 1
	val = s.pool[last]     // Retrieve item
	s.pool = s.pool[:last] // Shrink stack
	return
}

func (s *stack) PopN(n int) (vals []float64, err error) {
	vals = make([]float64, n)

	for idx := range vals {
		val, err := s.Pop()
		if err != nil {
			return vals, err
		}

		vals[idx] = val
	}
	return
}

func (s *stack) PopAll() ([]float64, error) {
	return s.PopN(s.Length())
}

func (s *stack) Duplicate() error {
	val, err := s.Pop()
	if err != nil {
		return err
	}

	s.PushN(val, val)
	return nil
}

func (s *stack) Swap() error {
	if s.Empty() || s.Length() < 2 {
		return errors.ErrStackEmpty
	}

	vals, err := s.PopN(2)
	if err != nil {
		return err
	}

	// Flip the values around and pop them back on the stack
	slices.Reverse(vals)
	s.PushN(vals...)
	return nil
}

func (s *stack) Rshift() error {
	if s.Empty() {
		return errors.ErrStackEmpty
	}

	last, err := s.Pop()
	if err != nil {
		return err
	}

	s.pool = append([]float64{last}, s.pool...)
	return nil
}

func (s *stack) Lshift() error {
	if s.Empty() {
		return errors.ErrStackEmpty
	}

	first := s.pool[0]
	s.pool = append(s.pool[1:], first)
	return nil
}

func (s *stack) TopShift() error {
	vals, err := s.PopN(3)
	if err != nil {
		return err
	}
	slices.Reverse(vals)

	last := vals[len(vals)-1]
	vals = append([]float64{last}, vals[:len(vals)-1]...)
	s.PushN(vals...)
	return nil
}

func (s *stack) Reverse() {
	slices.Reverse(s.pool)
}

func (s *stack) New() (ret Stack, err error) {
	x, err := s.Pop()
	if err != nil {
		return nil, err
	}

	ret = new(stack)
	size := s.Length()

	// When values are transferred to a new stack their order is preserved
	// (as opposed to pushed and popped) so we're avoiding allocations by
	// directly manipulating the pool
	vals := s.pool[size-int(x) : size]
	s.pool = s.pool[:size-int(x)]
	ret.PushN(vals...)
	return
}

func (s *stack) Consume(child Stack) {
	child.Reverse()
	elements, _ := child.PopAll()
	s.PushN(elements...)
}

func (s *stack) Length() int {
	return len(s.pool)
}

func (s *stack) Empty() bool {
	return len(s.pool) == 0
}

func (s *stack) Clear() {
	s.pool = make([]float64, 0)
}

func (s *stack) Register() error {
	if s.registerInUse {
		s.Push(s.register)
		s.register = float64(0)
	} else {
		top, err := s.Pop()
		if err != nil {
			return err
		}
		s.register = top
	}

	s.registerInUse = !s.registerInUse
	return nil
}
