package ramda

// Compose takes a list of functions and returns a new function that applies
// each function in sequence from right to left. The result of each function
// is passed as the argument to the next function in the chain.
//
// Example:
//
//	double := func(x int) int { return x * 2 }
//	addOne := func(x int) int { return x + 1 }
//	square := func(x int) int { return x * x }
//	composed := Compose(square, addOne, double)
//	result := composed(5) // ((5 * 2) + 1)^2 = 121
func Compose[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		result := x
		// Apply functions from right to left
		for i := len(fns) - 1; i >= 0; i-- {
			result = fns[i](result)
		}
		return result
	}
}
