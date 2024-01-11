package interpreter

import "slices"

// A Stack is an object capable of performing stack operations
type Stack interface {
	// Push adds a value to the Stack
	Push(val rune)
	// PushN adds multiple values to the Stack
	PushN(vals ...rune)

	// Pop removes the last value from the Stack and returns it
	Pop() (rune, error)
	// PopN removes the last N values from the Stack and returns it
	PopN(n int) ([]rune, error)

	// Duplicate adds another copy of the top value to the Stack
	Duplicate() error
	// Swap switches the places of the top two values on the Stack
	Swap() error

	// Rshift moves the entire stack to the right ([1,2,3,4] -> [4,1,2,3])
	Rshift() error
	// Lshift moves the entire stack to the right ([1,2,3,4] -> [2,3,4,1])
	Lshift() error
	// TopShift takes the top three values and shifts them right ([1,2,3,4] -> [1,4,2,3])
	TopShift() error

	// Reverse reverse the order of the Stack
	Reverse()
	// Length returns the size of the stack
	Length() int

	// Register toggles an element into or out of the register
	Register() error

	// Empty returns whether or not the pool is empty
	Empty() bool
	// Clear empties the stack
	Clear()
}

// A stack holds runes (int32) as well as a single-element register
type stack struct {
	pool     []rune
	register rune

	registerInUse bool // An unfortunate consequence of Go's zero values
}

func (s *stack) Push(val rune) {
	s.pool = append(s.pool, val)
}

func (s *stack) PushN(vals ...rune) {
	s.pool = append(s.pool, vals...)
}

func (s *stack) Pop() (val rune, err error) {
	if s.Empty() {
		err = ErrStackEmpty
		return
	}

	last := len(s.pool) - 1
	val = s.pool[last]     // Retrieve item
	s.pool = s.pool[:last] // Shrink stack
	return
}

func (s *stack) PopN(n int) (vals []rune, err error) {
	vals = make([]rune, n)

	for idx := range vals {
		val, err := s.Pop()
		if err != nil {
			return vals, err
		}

		vals[idx] = val
	}
	return
}

func (s *stack) Duplicate() error {
	val, err := s.Pop()
	if err != nil {
		return err
	}

	s.Push(val)
	s.Push(val)
	return nil
}

func (s *stack) Swap() error {
	if s.Empty() || s.Length() < 2 {
		return ErrStackEmpty
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
		return ErrStackEmpty
	}

	last, err := s.Pop()
	if err != nil {
		return err
	}

	s.pool = append([]rune{last}, s.pool...)
	return nil
}

func (s *stack) Lshift() error {
	if s.Empty() {
		return ErrStackEmpty
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

	last := vals[len(vals)-1]
	vals = append([]rune{last}, vals[1:len(vals)-1]...)
	s.PushN(vals...)
	return nil
}

func (s *stack) Reverse() {
	slices.Reverse(s.pool)
}

func (s *stack) Length() int {
	return len(s.pool)
}

func (s *stack) Empty() bool {
	return len(s.pool) == 0
}

func (s *stack) Clear() {
	s.pool = make([]rune, 0)
}

func (s *stack) Register() error {
	if s.registerInUse {
		s.Push(s.register)
		s.register = rune(0)
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
