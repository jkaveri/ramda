# Ramda

Ramda is a Go library that provides functional programming utilities for working with slices and collections. Inspired by the [Ramda.js](https://ramdajs.com/) library, it offers a collection of pure functions that make working with slices in Go more functional and composable.

## Features

- Pure functions with no side effects
- Generic type support for maximum flexibility
- Immutable operations (original slices are not modified)
- Comprehensive set of slice operations
- Type-safe implementations

## Installation

```bash
go get github.com/jkaveri/ramda
```

## Usage

Import the package:

```go
import "github.com/jkaveri/ramda"
```

## Available Functions

### Map
Transform each element in a slice using a function.

```go
numbers := []int{1, 2, 3, 4, 5}
doubled := ramda.Map(func(n int) int { return n * 2 }, numbers)
// Result: []int{2, 4, 6, 8, 10}
```

### Filter
Keep only elements that satisfy a predicate function.

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
even := ramda.Filter(func(n int) bool { return n%2 == 0 }, numbers)
// Result: []int{2, 4, 6, 8, 10}
```

### Reduce
Accumulate values in a slice using a function.

```go
numbers := []int{1, 2, 3, 4, 5}
sum := ramda.Reduce(func(acc, n int) int { return acc + n }, 0, numbers)
// Result: 15
```

### Find
Find the first element that satisfies a predicate.

```go
numbers := []int{1, 2, 3, 4, 5}
found, exists := ramda.Find(func(n int) bool { return n > 3 }, numbers)
// Result: found = 4, exists = true
```

### Any/All
Check if any or all elements satisfy a predicate.

```go
numbers := []int{1, 2, 3, 4, 5}
hasEven := ramda.Any(func(n int) bool { return n%2 == 0 }, numbers)
// Result: true

allEven := ramda.All(func(n int) bool { return n%2 == 0 }, numbers)
// Result: false
```

### Take/Drop
Get or remove elements from the beginning of a slice.

```go
numbers := []int{1, 2, 3, 4, 5}
firstThree := ramda.Take(3, numbers)
// Result: []int{1, 2, 3}

withoutFirstTwo := ramda.Drop(2, numbers)
// Result: []int{3, 4, 5}
```

### Unique
Remove duplicate elements while preserving order.

```go
numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
unique := ramda.Unique(numbers)
// Result: []int{1, 2, 3, 4}
```

### Flatten
Convert a 2D slice into a 1D slice.

```go
matrix := [][]int{{1, 2}, {3, 4}, {5}}
flattened := ramda.Flatten(matrix)
// Result: []int{1, 2, 3, 4, 5}
```

### Zip
Combine two slices into pairs.

```go
numbers := []int{1, 2, 3}
letters := []string{"a", "b", "c"}
zipped := ramda.Zip(numbers, letters)
// Result: []struct{First: int, Second: string}{
//   {1, "a"}, {2, "b"}, {3, "c"},
// }
```

### Reverse
Create a new slice with elements in reverse order.

```go
numbers := []int{1, 2, 3, 4, 5}
reversed := ramda.Reverse(numbers)
// Result: []int{5, 4, 3, 2, 1}
```

### GroupBy
Group elements by a key function.

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
grouped := ramda.GroupBy(func(n int) string {
    if n%2 == 0 {
        return "even"
    }
    return "odd"
}, numbers)
// Result: map[string][]int{
//   "even": []int{2, 4, 6, 8, 10},
//   "odd":  []int{1, 3, 5, 7, 9},
// }
```

### SortBy
Sort a slice using a custom comparison function.

```go
numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
sorted := ramda.SortBy(func(a, b int) bool { return a < b }, numbers)
// Result: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.



## License

This project is licensed under the MIT License - see the LICENSE file for details.
