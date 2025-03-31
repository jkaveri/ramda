package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurry(t *testing.T) {
	// Test basic currying with two arguments
	add := func(a, b int) int { return a + b }
	curriedAdd := Curry(add)
	addOne := curriedAdd(1)
	assert.Equal(t, 3, addOne(2))

	// Test currying with three arguments
	add3 := func(a, b, c int) int { return a + b + c }
	curriedAdd3 := Curry3(add3)
	addOne3 := curriedAdd3(1)
	addTwo3 := addOne3(2)
	assert.Equal(t, 6, addTwo3(3))

	// Test currying with four arguments
	add4 := func(a, b, c, d int) int { return a + b + c + d }
	curriedAdd4 := Curry4(add4)
	addOne4 := curriedAdd4(1)
	addTwo4 := addOne4(2)
	addThree4 := addTwo4(3)
	assert.Equal(t, 10, addThree4(4))
}
