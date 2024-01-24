package interpreter

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The max depth *for testing* - actual max depth depends on default platform int size
const MAX_DEPTH = math.MaxInt16

func TestPush(t *testing.T) {
	assert, stack := setup(t)
	letter := rune(rand.Int31())

	stack.Push(letter)

	retrieved, err := stack.Pop()
	assert.Nil(err)
	assert.Equal(retrieved, letter)
}

func TestPushN(t *testing.T) {
	assert, stack := setup(t)
	letters := randLetters(rand.Intn(MAX_DEPTH))

	stack.PushN(letters...)
	assert.Equal(stack.Length(), len(letters))

	// Reverse the input array to match pop order and compare
	slices.Reverse(letters)
	for _, val := range letters {
		stackVal, _ := stack.Pop()
		assert.Equal(val, stackVal)
	}
}

func TestPop(t *testing.T) {
	assert, stack := setup(t)
	letter := randLetter()

	stack.Push(letter)
	popped, _ := stack.Pop()
	assert.Equal(letter, popped)
}

func TestPopN(t *testing.T) {
	assert, stack := setup(t)
	letters := randLetters(rand.Intn(MAX_DEPTH))

	stack.PushN(letters...)
	assert.Equal(len(letters), stack.Length())

	// Reverse the input array to match pop order and compare
	slices.Reverse(letters)
	received, _ := stack.PopN(len(letters))

	for idx, val := range letters {
		assert.Equal(val, received[idx])
	}
}

func TestPopAll(t *testing.T) {
	assert, stack := setup(t)
	letters := randLetters(MAX_DEPTH)

	stack.PushN(letters...)
	popped, _ := stack.PopAll()
	assert.Equal(len(letters), len(popped))
}

func TestDuplicate(t *testing.T) {
	assert, stack := setup(t)
	letter := randLetter()

	stack.Push(letter)
	assert.Equal(1, stack.Length())
	stack.Duplicate()
	assert.Equal(2, stack.Length())

	var (
		expected    = []rune{letter, letter}
		received, _ = stack.PopAll()
	)
	assert.Equal(expected, received)
	// assert.Equal(expected, received)
}

func TestSwap(t *testing.T) {
	assert, stack := setup(t)
	expected := randLetters(2)

	stack.PushN(expected...)
	stack.Swap()

	slices.Reverse(expected)
	received, _ := stack.PopAll()
	assert.Equal(expected, received)
}

func TestRshift(t *testing.T) {
	assert, stack := setup(t)
	expected := randLetters(3)

	stack.PushN(expected...)
	stack.Rshift() // Dropping a (hopefully) impossible error

	// manual rshift + reverse to account for LIFO
	expected = []rune{expected[1], expected[0], expected[2]}
	received, _ := stack.PopAll()
	assert.Equal(expected, received)
}

func TestLshift(t *testing.T) {
	assert, stack := setup(t)
	expected := randLetters(3)

	stack.PushN(expected...)
	stack.Lshift() // Dropping a (hopefully) impossible error

	// manual lshift + reverse to account for LIFO
	expected = []rune{expected[0], expected[2], expected[1]}
	received, _ := stack.PopN(3)
	assert.Equal(expected, received)
}

func TestTopShift(t *testing.T) {
	assert, actual := setup(t)
	expected := new(stack)
	letters := make([]rune, 0)
	for len(letters) < 4 {
		letters = randLetters(64)
	}

	expected.PushN(letters...)
	actual.PushN(letters...)

	expected.Push(3)
	child, _ := expected.New()
	child.Rshift()
	expected.Consume(child)

	actual.TopShift()

	assert.Equal(expected.pool, actual.pool)
}

func TestReverse(t *testing.T) {
	assert, stack := setup(t)
	letters := randLetters(MAX_DEPTH)
	stack.PushN(letters...)
	stack.Reverse()

	actual, _ := stack.PopAll()
	assert.Equal(letters, actual)
}

// Is this needed? No.
// Is this wanted? No.
// Am I gonna put it? Yep.
func TestLength(t *testing.T) {
	assert, stack := setup(t)
	letters := randLetters(MAX_DEPTH)
	stack.PushN(letters...)
	assert.Equal(len(letters), stack.Length())
}

func TestNew(t *testing.T) {
	assert, stack := setup(t)
	letters := randLetters(MAX_DEPTH)
	stack.PushN(letters...)

	childSize := rand.Int31n(int32(len(letters) - 1))
	stack.Push(childSize)
	child, _ := stack.New()

	assert.Equal(int(childSize), child.Length())
	assert.Equal(stack.Length(), len(letters)-int(childSize))
}

/**** Utility ****/
func setup(t *testing.T) (*assert.Assertions, *stack) {
	return assert.New(t), new(stack)
}

func randLetter() rune {
	return rune(rand.Int31())
}

func randLetters(n int) (out []rune) {
	out = make([]rune, n)
	for idx := range out {
		out[idx] = randLetter()
	}
	return
}
