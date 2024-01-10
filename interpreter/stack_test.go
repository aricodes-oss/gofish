package interpreter

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAX_SAMPLE_SIZE = 128

func setup(t *testing.T) (*assert.Assertions, *stack) {
	return assert.New(t), new(stack)
}

func randLetter() rune {
	return rune(rand.Int31())
}

func randNLetters(n int) (out []rune) {
	out = make([]rune, n)
	for idx := range out {
		out[idx] = randLetter()
	}
	return
}

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
	letters := randNLetters(rand.Intn(MAX_SAMPLE_SIZE))

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
	letters := randNLetters(rand.Intn(MAX_SAMPLE_SIZE))

	stack.PushN(letters...)
	assert.Equal(len(letters), stack.Length())

	// Reverse the input array to match pop order and compare
	slices.Reverse(letters)
	received, _ := stack.PopN(len(letters))

	for idx, val := range letters {
		assert.Equal(val, received[idx])
	}
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
		received, _ = stack.PopN(2)
	)
	assert.ElementsMatch(expected, received)
}
