package rslice

// Map applies a function to each element of a slice and returns a new slice with the results.
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

// Filter creates a new slice containing only the elements that satisfy the predicate function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	even := Filter(func(n int) bool { return n%2 == 0 }, numbers)
//	// Result: []int{2, 4, 6}
func Filter[T any](fn func(T) bool, slice []T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce applies a function to each element of a slice, accumulating the result.
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

// Find returns the first element that satisfies the predicate function, along with a boolean indicating if an element was found.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found, exists := Find(func(n int) bool { return n > 3 }, numbers)
//	// Result: found = 4, exists = true
func Find[T any](fn func(T) bool, slice []T) (T, bool) {
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Any returns true if at least one element in the slice satisfies the predicate function.
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

// All returns true if all elements in the slice satisfy the predicate function.
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

// Take returns the first n elements from a slice.
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
		return append([]T{}, slice...)
	}
	return append([]T{}, slice[:n]...)
}

// Drop returns a slice with the first n elements removed.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	withoutFirstTwo := Drop(2, numbers)
//	// Result: []int{3, 4, 5}
func Drop[T any](n int, slice []T) []T {
	if n <= 0 {
		return append([]T{}, slice...)
	}
	if n >= len(slice) {
		return []T{}
	}
	return append([]T{}, slice[n:]...)
}

// Unique returns a new slice with duplicate elements removed.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4}
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

// Flatten converts a slice of slices into a single slice.
//
// Example:
//
//	matrix := [][]int{{1, 2}, {3, 4}, {5}}
//	flattened := Flatten(matrix)
//	// Result: []int{1, 2, 3, 4, 5}
func Flatten[T any](slice [][]T) []T {
	var result []T
	for _, subSlice := range slice {
		result = append(result, subSlice...)
	}
	return result
}

// Zip combines two slices into a slice of pairs.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	letters := []string{"a", "b", "c"}
//	zipped := Zip(numbers, letters)
//	// Result: []struct{First int; Second string}{{1, "a"}, {2, "b"}, {3, "c"}}
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
		}{slice1[i], slice2[i]}
	}

	return result
}

// Reverse returns a new slice with elements in reverse order.
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

// GroupBy groups elements of a slice by a key function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	grouped := GroupBy(func(n int) string {
//		if n%2 == 0 {
//			return "even"
//		}
//		return "odd"
//	}, numbers)
//	// Result: map[string][]int{"even": {2, 4, 6}, "odd": {1, 3, 5}}
func GroupBy[T any, K comparable](fn func(T) K, slice []T) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// SortBy sorts a slice using a comparison function.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	sorted := SortBy(func(a, b int) bool { return a < b }, numbers)
//	// Result: []int{1, 1, 2, 3, 4, 5, 6, 9}
func SortBy[T any](fn func(T, T) bool, slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	// Simple bubble sort for demonstration
	// In practice, you might want to use a more efficient sorting algorithm
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if !fn(result[j], result[j+1]) {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// IndexBy creates a map from a slice using a function to generate keys.
// Each element in the slice becomes a value in the map, with the key determined
// by applying the provided function to the element. If multiple elements produce
// the same key, the last element will overwrite previous ones.
//
// Example:
//
//	users := []struct{ID string; Name string}{
//	  {ID: "1", Name: "Alice"},
//	  {ID: "2", Name: "Bob"},
//	  {ID: "3", Name: "Charlie"},
//	}
//	indexed := IndexBy(func(u struct{ID string; Name string}) string { return u.ID }, users)
//	// Result: map[string]struct{ID string; Name string}{
//	//   "1": {ID: "1", Name: "Alice"},
//	//   "2": {ID: "2", Name: "Bob"},
//	//   "3": {ID: "3", Name: "Charlie"},
//	// }
func IndexBy[T any, K comparable](fn func(T) K, slice []T) map[K]T {
	result := make(map[K]T, len(slice))
	for _, v := range slice {
		key := fn(v)
		result[key] = v
	}
	return result
}

// ToSet creates a set-like map from a slice of comparable values.
// The result is a map[T]struct{} which can be used for efficient membership checking.
// Duplicate values are automatically handled since map keys must be unique.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
//	numberSet := ToSet(numbers)
//	// Result: map[int]struct{}{
//	//   1: {},
//	//   2: {},
//	//   3: {},
//	//   4: {},
//	//   5: {},
//	// }
//
//	// Check membership
//	if _, exists := numberSet[2]; exists {
//	  fmt.Println("2 is in the set")
//	}
func ToSet[T comparable](slice []T) map[T]struct{} {
	result := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		result[v] = struct{}{}
	}
	return result
}
