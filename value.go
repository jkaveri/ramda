package ramda

import (
	"fmt"
	"strconv"
)

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
func DefaultFn[T comparable](defaultFn func() T, val T) T {
	var empty T
	if val == empty {
		return defaultFn()
	}

	return val
}

// Cast safely converts a value from one type to another using a conversion function.
// If the conversion fails, it returns the default value.
//
// Example:
//
//	toInt := func(s string) (int, error) { return strconv.Atoi(s) }
//	result := Cast(toInt, 0, "123") // 123
//	result2 := Cast(toInt, 0, "abc") // 0
func Cast[T, R any](converter func(T) (R, error), defaultVal R, val T) R {
	result, err := converter(val)
	if err != nil {
		return defaultVal
	}
	return result
}

// CastFn safely converts a value from one type to another using a conversion function.
// If the conversion fails, it calls the provided function to get a default value.
func CastFn[T, R any](converter func(T) (R, error), defaultFn func() R, val T) R {
	result, err := converter(val)
	if err != nil {
		return defaultFn()
	}
	return result
}

// ToString converts any value to its string representation.
// For basic types, it uses fmt.Sprintf, for custom types it uses their String() method if available.
func ToString[T any](val T) string {
	return fmt.Sprintf("%v", val)
}

// ToInt converts a string to an integer, returning 0 if conversion fails.
func ToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

// ToInt64 converts a string to an int64, returning 0 if conversion fails.
func ToInt64(s string) int64 {
	result, _ := strconv.ParseInt(s, 10, 64)
	return result
}

// ToFloat64 converts a string to a float64, returning 0.0 if conversion fails.
func ToFloat64(s string) float64 {
	result, _ := strconv.ParseFloat(s, 64)
	return result
}

// ToBool converts a string to a boolean, returning false if conversion fails.
func ToBool(s string) bool {
	result, _ := strconv.ParseBool(s)
	return result
}

// FromInt converts an integer to a string.
func FromInt(i int) string {
	return strconv.Itoa(i)
}

// FromInt64 converts an int64 to a string.
func FromInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// FromFloat64 converts a float64 to a string.
func FromFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// FromBool converts a boolean to a string.
func FromBool(b bool) string {
	return strconv.FormatBool(b)
}

// As converts a value to a specific type using a type assertion.
// Returns the zero value of the target type if the assertion fails.
func As[T any](val interface{}) T {
	if result, ok := val.(T); ok {
		return result
	}
	var zero T
	return zero
}

// AsWithDefault converts a value to a specific type using a type assertion.
// Returns the provided default value if the assertion fails.
func AsWithDefault[T any](val interface{}, defaultVal T) T {
	if result, ok := val.(T); ok {
		return result
	}
	return defaultVal
}

// AsWithDefaultFn converts a value to a specific type using a type assertion.
// Calls the provided function to get a default value if the assertion fails.
func AsWithDefaultFn[T any](defaultFn func() T, val interface{}) T {
	if result, ok := val.(T); ok {
		return result
	}
	return defaultFn()
}

// Transform applies a transformation function to a value and returns the result.
// This is useful for chaining transformations in a functional style.
func Transform[T, R any](transformer func(T) R, val T) R {
	return transformer(val)
}

// TransformWithError applies a transformation function that may return an error.
// If the transformation fails, it returns the default value.
func TransformWithError[T, R any](transformer func(T) (R, error), defaultVal R, val T) R {
	result, err := transformer(val)
	if err != nil {
		return defaultVal
	}
	return result
}

// TransformWithErrorFn applies a transformation function that may return an error.
// If the transformation fails, it calls the provided function to get a default value.
func TransformWithErrorFn[T, R any](transformer func(T) (R, error), defaultFn func() R, val T) R {
	result, err := transformer(val)
	if err != nil {
		return defaultFn()
	}
	return result
}

// FromPtr converts a pointer to its value, returning the zero value if the pointer is nil.
// This is useful for safely dereferencing pointers without panic.
func FromPtr[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// ToPtr converts a value to a pointer, always creating a pointer.
// This is useful for creating pointers to values.
func ToPtr[T any](val T) *T {
	return &val
}

// NilIfEmpty converts a value to a pointer. If the value is the zero value, it returns nil.
// This is useful for creating optional values.
func NilIfEmpty[T comparable](val T) *T {
	var zero T
	if val == zero {
		return nil
	}
	return &val
}
