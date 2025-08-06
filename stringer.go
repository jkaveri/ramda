package ramda

import (
	"github.com/jkaveri/ramda/rslice"
)

// Stringer is an interface that wraps the String method.
type Stringer interface {
	String() string
}

// Stringify converts a value that implements the Stringer interface to a string.
// It's a convenience wrapper around the String() method.
func Stringify[T Stringer](item T) string {
	return item.String()
}

// StringifySlice converts a slice of values that implement the Stringer interface
// to a slice of strings. It applies Stringify to each element in the slice.
func StringifySlice[T Stringer](items []T) []string {
	return rslice.Map(Stringify, items)
}
