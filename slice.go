package ramda

import (
	"sort"
)

// Map applies a function to each element in a slice and returns a new slice
// containing the results. It transforms each element using the provided function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := Map(func(n int) int { return n * 2 }, numbers)
//	// Result: []int{2, 4, 6, 8, 10}
func Map[T, R any](fn func(T) R, slice []T) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
// It removes elements for which the predicate returns false.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	even := Filter(func(n int) bool { return n%2 == 0 }, numbers)
//	// Result: []int{2, 4, 6, 8, 10}
func Filter[T any](fn func(T) bool, slice []T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce applies a function to each element in a slice, accumulating the result.
// It starts with the initial value and applies the function from left to right.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sum := Reduce(func(acc, n int) int { return acc + n }, 0, numbers)
//	// Result: 15
func Reduce[T, R any](fn func(R, T) R, initial R, slice []T) R {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Find returns the first element in the slice that satisfies the predicate function,
// along with a boolean indicating whether such an element was found.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found, exists := Find(func(n int) bool { return n > 3 }, numbers)
//	// Result: found = 4, exists = true
func Find[T any](fn func(T) bool, slice []T) (T, bool) {
	var zero T
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	return zero, false
}

// Any returns true if at least one element in the slice satisfies the predicate function.
// It stops searching as soon as it finds a matching element.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := Any(func(n int) bool { return n%2 == 0 }, numbers)
//	// Result: true
func Any[T any](fn func(T) bool, slice []T) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// All returns true if every element in the slice satisfies the predicate function.
// It stops searching as soon as it finds a non-matching element.
//
// Example:
//
//	numbers := []int{2, 4, 6, 8, 10}
//	allEven := All(func(n int) bool { return n%2 == 0 }, numbers)
//	// Result: true
func All[T any](fn func(T) bool, slice []T) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Take returns the first n elements from the slice.
// If n is greater than the slice length, it returns the entire slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	firstThree := Take(3, numbers)
//	// Result: []int{1, 2, 3}
func Take[T any](n int, slice []T) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

// Drop returns the slice with the first n elements removed.
// If n is greater than the slice length, it returns an empty slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	withoutFirstTwo := Drop(2, numbers)
//	// Result: []int{3, 4, 5}
func Drop[T any](n int, slice []T) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}

// Unique returns a new slice containing only unique elements from the input slice.
// It removes all duplicates while preserving the order of first occurrence.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
//	unique := Unique(numbers)
//	// Result: []int{1, 2, 3, 4}
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Flatten converts a 2D slice into a 1D slice by concatenating all inner slices.
// It preserves the order of elements as they appear in the input.
//
// Example:
//
//	matrix := [][]int{{1, 2}, {3, 4}, {5}}
//	flattened := Flatten(matrix)
//	// Result: []int{1, 2, 3, 4, 5}
func Flatten[T any](slice [][]T) []T {
	var result []T
	for _, v := range slice {
		result = append(result, v...)
	}
	return result
}

// Zip combines two slices into a slice of pairs, where each pair contains elements
// at the same index from both input slices. The result is truncated to the length
// of the shorter input slice.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	letters := []string{"a", "b", "c"}
//	zipped := Zip(numbers, letters)
//	// Result: []struct{First: int, Second: string}{
//	//   {1, "a"}, {2, "b"}, {3, "c"},
//	// }
func Zip[T, U any](slice1 []T, slice2 []U) []struct {
	First  T
	Second U
} {
	minLen := len(slice1)
	if len(slice2) < minLen {
		minLen = len(slice2)
	}
	result := make([]struct {
		First  T
		Second U
	}, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = struct {
			First  T
			Second U
		}{
			First:  slice1[i],
			Second: slice2[i],
		}
	}
	return result
}

// Reverse returns a new slice with the elements in reverse order.
// It does not modify the input slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	reversed := Reverse(numbers)
//	// Result: []int{5, 4, 3, 2, 1}
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// GroupBy returns a map where the keys are the results of applying the function
// to each element, and the values are slices of elements that produced that key.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	grouped := GroupBy(func(n int) string {
//	  if n%2 == 0 {
//	    return "even"
//	  }
//	  return "odd"
//	}, numbers)
//	// Result: map[string][]int{
//	//   "even": []int{2, 4, 6, 8, 10},
//	//   "odd":  []int{1, 3, 5, 7, 9},
//	// }
func GroupBy[T any, K comparable](fn func(T) K, slice []T) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// SortBy sorts a slice using the provided comparison function.
// The comparison function should return true if the first argument should come
// before the second argument in the sorted result.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
//	sorted := SortBy(func(a, b int) bool { return a < b }, numbers)
//	// Result: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
func SortBy[T any](fn func(T, T) bool, slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return fn(result[i], result[j])
	})
	return result
}
