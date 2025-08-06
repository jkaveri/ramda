# Slice Operations Examples

This directory contains examples demonstrating the slice utility functions in the `rslice` package.

## Overview

The `rslice` package provides functional programming utilities for working with slices in Go. All functions are generic and type-safe, making them suitable for any slice type.

## Key Functions

### Transformation Functions
- `rslice.Map` - Transform each element using a function
- `rslice.Filter` - Select elements based on a predicate
- `rslice.Reduce` - Accumulate values using a function

### Search and Query Functions
- `rslice.Find` - Find the first element matching a predicate
- `rslice.Any` - Check if any element matches a predicate
- `rslice.All` - Check if all elements match a predicate

### Slice Manipulation
- `rslice.Take` - Get the first n elements
- `rslice.Drop` - Remove the first n elements
- `rslice.Reverse` - Reverse the order of elements
- `rslice.Unique` - Remove duplicate elements
- `rslice.Flatten` - Flatten nested slices

### Combination and Grouping
- `rslice.Zip` - Combine two slices into pairs
- `rslice.GroupBy` - Group elements by a key function
- `rslice.SortBy` - Sort using a custom comparison function

### Conversion Functions
- `rslice.IndexBy` - Create a map from a slice using a key function
- `rslice.ToSet` - Create a set-like map from a slice

## Examples

### 1. Map - Transform elements
```go
numbers := []int{1, 2, 3, 4, 5}
doubled := rslice.Map(func(n int) int { return n * 2 }, numbers)
// Result: []int{2, 4, 6, 8, 10}
```

### 2. Filter - Select elements
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
even := rslice.Filter(func(n int) bool { return n%2 == 0 }, numbers)
// Result: []int{2, 4, 6, 8, 10}
```

### 3. Reduce - Accumulate values
```go
numbers := []int{1, 2, 3, 4, 5}
sum := rslice.Reduce(func(acc, n int) int { return acc + n }, 0, numbers)
// Result: 15
```

### 4. Find - Find first matching element
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
found, exists := rslice.Find(func(n int) bool { return n > 7 }, numbers)
// Result: found = 8, exists = true
```

### 5. Any and All - Check conditions
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
hasEven := rslice.Any(func(n int) bool { return n%2 == 0 }, numbers)
// Result: true

allPositive := rslice.All(func(n int) bool { return n > 0 }, numbers)
// Result: true
```

### 6. Take and Drop - Slice operations
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
firstThree := rslice.Take(3, numbers)
// Result: []int{1, 2, 3}

withoutFirstTwo := rslice.Drop(2, numbers)
// Result: []int{3, 4, 5, 6, 7, 8, 9, 10}
```

### 7. Unique - Remove duplicates
```go
numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
unique := rslice.Unique(numbers)
// Result: []int{1, 2, 3, 4}
```

### 8. Flatten - Flatten nested slices
```go
matrix := [][]int{{1, 2}, {3, 4}, {5, 6, 7}, {8}}
flattened := rslice.Flatten(matrix)
// Result: []int{1, 2, 3, 4, 5, 6, 7, 8}
```

### 9. Zip - Combine two slices
```go
numbers := []int{1, 2, 3, 4}
letters := []string{"a", "b", "c", "d", "e"}
zipped := rslice.Zip(numbers, letters)
// Result: []struct{First int; Second string}{
//   {1, "a"}, {2, "b"}, {3, "c"}, {4, "d"}
// }
```

### 10. Reverse - Reverse slice order
```go
numbers := []int{1, 2, 3, 4, 5}
reversed := rslice.Reverse(numbers)
// Result: []int{5, 4, 3, 2, 1}
```

### 11. GroupBy - Group elements by key
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
grouped := rslice.GroupBy(func(n int) string {
    if n%2 == 0 {
        return "even"
    }
    return "odd"
}, numbers)
// Result: map[string][]int{
//   "even": {2, 4, 6, 8, 10},
//   "odd": {1, 3, 5, 7, 9}
// }
```

### 12. SortBy - Sort with custom comparison
```go
numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
sorted := rslice.SortBy(func(a, b int) bool { return a < b }, numbers)
// Result: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
```

### 13. IndexBy - Create map from slice
```go
type User struct {
    ID   string
    Name string
    Age  int
}

users := []User{
    {ID: "1", Name: "Alice", Age: 25},
    {ID: "2", Name: "Bob", Age: 30},
    {ID: "3", Name: "Charlie", Age: 35},
}

indexed := rslice.IndexBy(func(u User) string { return u.ID }, users)
// Result: map[string]User{
//   "1": {ID: "1", Name: "Alice", Age: 25},
//   "2": {ID: "2", Name: "Bob", Age: 30},
//   "3": {ID: "3", Name: "Charlie", Age: 35}
// }
```

### 14. ToSet - Create set from slice
```go
numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
numberSet := rslice.ToSet(numbers)
// Result: map[int]struct{}{
//   1: {}, 2: {}, 3: {}, 4: {}, 5: {}
// }

// Check membership
if _, exists := numberSet[3]; exists {
    fmt.Println("3 is in the set")
}
```

### 15. Function composition
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Filter even numbers, double them, take first 3, reverse
result := rslice.Reverse(
    rslice.Take(3,
        rslice.Map(func(n int) int { return n * 2 },
            rslice.Filter(func(n int) bool { return n%2 == 0 }, numbers))))
// Result: []int{6, 4, 2}
```

## Running the Examples

To run the slice operations examples:

```bash
go run _examples/slice_operations/slice_example.go
```

## Expected Output

The example will demonstrate all 15 slice operations with detailed output showing:

1. **Map** - Transforming elements with functions
2. **Filter** - Selecting elements based on predicates
3. **Reduce** - Accumulating values
4. **Find** - Finding first matching elements
5. **Any/All** - Checking conditions across slices
6. **Take/Drop** - Basic slice operations
7. **Unique** - Removing duplicates
8. **Flatten** - Flattening nested structures
9. **Zip** - Combining slices
10. **Reverse** - Reversing order
11. **GroupBy** - Grouping by keys
12. **SortBy** - Custom sorting
13. **IndexBy** - Creating maps from slices
14. **ToSet** - Creating sets
15. **Function composition** - Chaining operations

## Use Cases

These slice functions are particularly useful for:

- **Data Processing**: Transform, filter, and aggregate data
- **Functional Programming**: Compose operations in a functional style
- **Data Analysis**: Group, sort, and analyze collections
- **API Development**: Process request/response data
- **Testing**: Create test data and verify results
- **Performance**: Efficient operations on large datasets

## Benefits

- **Type Safety**: All functions use Go generics for compile-time type checking
- **Immutability**: Functions return new slices rather than modifying inputs
- **Composability**: Functions can be easily chained together
- **Performance**: Optimized implementations for common operations
- **Readability**: Clear, descriptive function names
- **Consistency**: Uniform API across all slice operations
