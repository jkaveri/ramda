package ramda

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

// Identity returns the input boolean value unchanged.
// It's useful as a default predicate function.
func Identity(t bool) bool {
	return t
}

// Zero returns true if the input value is the zero value for its type.
// It's useful for checking if a value has been initialized.
// This is a pure generic function which will have better performance than Empty.
func Zero[T comparable](a T) bool {
	var zero T
	return a == zero
}

// Empty returns true if the input value is empty (nil, empty string, empty slice, etc.).
// For slices, maps, and channels, it checks for nil or zero length.
// For other types, it checks if the value is zero.
func Empty(a any) bool {
	v := reflect.ValueOf(a)

	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.IsNil() || v.Len() == 0
	default:
		return v.IsZero()
	}
}

// NonEmpty returns true if the input value is not empty.
// It's the inverse of Empty.
func NonEmpty(a any) bool {
	return !Empty(a)
}

// Equal returns true if two values of the same type are equal.
// It uses Go's built-in equality comparison.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Nil returns true if the input value is nil.
func Nil(a any) bool {
	return a == nil
}

// NotNil returns true if the input value is not nil.
func NotNil(a any) bool {
	return a != nil
}

// IsString returns true if the input value is a string.
func IsString(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.String
}

// IsNumber returns true if the input value is a numeric type.
func IsNumber(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.Int || reflect.TypeOf(a).Kind() == reflect.Float64
}

// IsPositive returns true if the input signed number is greater than zero.
func IsPositive[T constraints.Signed](a T) bool {
	return a > 0
}

// IsNegative returns true if the input signed number is less than zero.
func IsNegative[T constraints.Signed](a T) bool {
	return a < 0
}

// IsEven returns true if the input integer is even.
func IsEven[T constraints.Integer](a T) bool {
	return a%2 == 0
}

// IsOdd returns true if the input integer is odd.
func IsOdd[T constraints.Integer](a T) bool {
	return a%2 != 0
}
