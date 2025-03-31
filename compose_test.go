package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	// Test case from documentation
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }
	composed := Compose(square, addOne, double)

	// Test with input 5: ((5 * 2) + 1)^2 = 121
	result := composed(5)
	assert.Equal(t, 121, result)

	// Test with single function
	singleFn := Compose(double)
	assert.Equal(t, 10, singleFn(5))

	// Test with empty function list (should return input)
	emptyFn := Compose[int]()
	assert.Equal(t, 5, emptyFn(5))
}
