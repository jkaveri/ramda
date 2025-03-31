package ramda

// Default returns the first value if it's not zero, otherwise returns the default value.
// It's useful for providing fallback values when dealing with potentially zero values.
func Default[T comparable](val T, defaultVal T) T {
	var empty T
	if val == empty {
		return defaultVal
	}

	return val
}

// DefaultFn returns the first value if it's not zero, otherwise calls the provided function
// to get a default value. This is useful when the default value needs to be computed
// or when you want to defer the creation of the default value until it's needed.
func DefaultFn[T comparable](val T, defaultFn func() T) T {
	var empty T
	if val == empty {
		return defaultFn()
	}

	return val
}
