package ramda

// Curry takes a function of two arguments and returns a curried version of it.
// The curried function can be called with one argument at a time, returning
// a new function that takes the remaining argument.
//
// Example:
//
//	add := func(a, b int) int { return a + b }
//	curriedAdd := Curry(add)
//	addOne := curriedAdd(1)
//	result := addOne(2) // 3
func Curry[T1, T2, R any](fn func(T1, T2) R) func(T1) func(T2) R {
	return func(a T1) func(T2) R {
		return func(b T2) R {
			return fn(a, b)
		}
	}
}

// Curry3 takes a function of three arguments and returns a curried version of it.
// The curried function can be called with one argument at a time, returning
// a new function that takes the remaining arguments.
func Curry3[T1, T2, T3, R any](fn func(T1, T2, T3) R) func(T1) func(T2) func(T3) R {
	return func(a T1) func(T2) func(T3) R {
		return func(b T2) func(T3) R {
			return func(c T3) R {
				return fn(a, b, c)
			}
		}
	}
}

// Curry4 takes a function of four arguments and returns a curried version of it.
// The curried function can be called with one argument at a time, returning
// a new function that takes the remaining arguments.
func Curry4[T1, T2, T3, T4, R any](fn func(T1, T2, T3, T4) R) func(T1) func(T2) func(T3) func(T4) R {
	return func(a T1) func(T2) func(T3) func(T4) R {
		return func(b T2) func(T3) func(T4) R {
			return func(c T3) func(T4) R {
				return func(d T4) R {
					return fn(a, b, c, d)
				}
			}
		}
	}
}
