# Ramda

Ramda is a comprehensive Go library that provides functional programming utilities for working with slices, maps, structs, and collections. Inspired by the [Ramda.js](https://ramdajs.com/) library, it offers a collection of pure functions that make working with Go data structures more functional and composable.

## Features

- **Pure functions** with no side effects
- **Generic type support** for maximum flexibility
- **Immutable operations** (original data structures are not modified)
- **Comprehensive set of utilities** for slices, maps, and structs
- **Type-safe implementations** with compile-time guarantees
- **Functional composition** and currying support
- **Predicate functions** for common operations
- **Value transformation** and casting utilities
- **Extensive examples** demonstrating all features

## Installation

```bash
go get github.com/jkaveri/ramda
```

## Usage

Import the main package:

```go
import "github.com/jkaveri/ramda"
```

Or import specific modules:

```go
import (
    "github.com/jkaveri/ramda"
    "github.com/jkaveri/ramda/rmap"
    "github.com/jkaveri/ramda/rslice"
    "github.com/jkaveri/ramda/rstruct"
)
```

## Core Utilities

### Functional Composition

#### Compose
Combine multiple functions into a single function that applies them from right to left.

```go
double := func(x int) int { return x * 2 }
addOne := func(x int) int { return x + 1 }
square := func(x int) int { return x * x }
composed := ramda.Compose(square, addOne, double)
result := composed(5) // ((5 * 2) + 1)^2 = 121
```

#### Curry
Create curried versions of functions for partial application.

```go
add := func(a, b int) int { return a + b }
curriedAdd := ramda.Curry(add)
addOne := curriedAdd(1)
result := addOne(2) // 3
```

### Predicate Functions

Common predicate functions for filtering and validation:

```go
numbers := []int{0, 1, 2, 3, 4, 5}

// Check for zero values
zeros := ramda.Filter(ramda.Zero, numbers) // []int{0}

// Check for empty values
empty := ramda.Filter(ramda.Empty, []string{"", "hello", ""}) // []string{"", ""}

// Check for even/odd numbers
even := ramda.Filter(ramda.IsEven, numbers) // []int{0, 2, 4}
odd := ramda.Filter(ramda.IsOdd, numbers)   // []int{1, 3, 5}

// Check for positive/negative numbers
positive := ramda.Filter(ramda.IsPositive, numbers) // []int{1, 2, 3, 4, 5}
```

### Value Utilities

Safe value handling and type conversion:

```go
// Default values
value := ramda.Default("", "hello") // "hello"
value = ramda.Default("existing", "hello") // "existing"

// Safe casting
number := ramda.Cast(strconv.Atoi, 0, "123") // 123
number = ramda.Cast(strconv.Atoi, 0, "abc")  // 0

// Type conversion
str := ramda.ToString(42) // "42"
num := ramda.ToInt("123") // 123
```

## Slice Operations (`rslice`)

### Basic Operations

```go
numbers := []int{1, 2, 3, 4, 5}

// Transform elements
doubled := rslice.Map(func(n int) int { return n * 2 }, numbers)
// Result: []int{2, 4, 6, 8, 10}

// Filter elements
even := rslice.Filter(func(n int) bool { return n%2 == 0 }, numbers)
// Result: []int{2, 4}

// Reduce elements
sum := rslice.Reduce(func(acc, n int) int { return acc + n }, 0, numbers)
// Result: 15
```

### Search and Query

```go
numbers := []int{1, 2, 3, 4, 5}

// Find first matching element
found, exists := rslice.Find(func(n int) bool { return n > 3 }, numbers)
// Result: found = 4, exists = true

// Check if any/all elements match
hasEven := rslice.Any(func(n int) bool { return n%2 == 0 }, numbers) // true
allEven := rslice.All(func(n int) bool { return n%2 == 0 }, numbers) // false
```

### Manipulation

```go
numbers := []int{1, 2, 3, 4, 5}

// Take/Drop elements
firstThree := rslice.Take(3, numbers)     // []int{1, 2, 3}
withoutFirstTwo := rslice.Drop(2, numbers) // []int{3, 4, 5}

// Remove duplicates
unique := rslice.Unique([]int{1, 2, 2, 3, 3, 3, 4})
// Result: []int{1, 2, 3, 4}

// Reverse
reversed := rslice.Reverse(numbers) // []int{5, 4, 3, 2, 1}
```

### Combination and Grouping

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Group by criteria
grouped := rslice.GroupBy(func(n int) string {
    if n%2 == 0 {
        return "even"
    }
    return "odd"
}, numbers)
// Result: map[string][]int{
//   "even": []int{2, 4, 6, 8, 10},
//   "odd":  []int{1, 3, 5, 7, 9},
// }

// Create lookup map
indexed := rslice.IndexBy(func(n int) int { return n * n }, numbers)
// Result: map[int]int{1: 1, 4: 2, 9: 3, 16: 4, 25: 5, ...}

// Zip two slices
letters := []string{"a", "b", "c"}
zipped := rslice.Zip(numbers[:3], letters)
// Result: []struct{First: int, Second: string}{
//   {1, "a"}, {2, "b"}, {3, "c"},
// }
```

## Map Operations (`rmap`)

### Transformation

```go
original := map[string]int{"a": 1, "b": 2, "c": 3}

// Transform keys
uppercase := rmap.TransformKeys(func(k string) string {
    return strings.ToUpper(k)
}, original)
// Result: map[string]int{"A": 1, "B": 2, "C": 3}

// Transform values
doubled := rmap.TransformValues(func(v int) int {
    return v * 2
}, original)
// Result: map[string]int{"a": 2, "b": 4, "c": 6}

// Transform entries
transformed := rmap.TransformEntries(func(k string, v int) (string, string) {
    return strings.ToUpper(k), fmt.Sprintf("value_%d", v)
}, original)
// Result: map[string]string{"A": "value_1", "B": "value_2", "C": "value_3"}
```

### Filtering and Access

```go
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

// Filter entries
evenValues := rmap.Filter(func(k string, v int) bool {
    return v%2 == 0
}, data)
// Result: map[string]int{"b": 2, "d": 4}

// Safe value access
value, exists := rmap.Get("a", data) // 1, true
value = rmap.GetOrElse("x", 0, data) // 0 (default value)

// Extract keys and values
keys := rmap.Keys(data)   // []string{"a", "b", "c", "d"}
values := rmap.Values(data) // []int{1, 2, 3, 4}
```

### Composition

```go
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"c": 3, "d": 4}

// Merge maps
merged := rmap.Merge(map1, map2)
// Result: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

// Pick specific keys
picked := rmap.Pick([]string{"a", "c"}, merged)
// Result: map[string]int{"a": 1, "c": 3}

// Omit specific keys
omitted := rmap.Omit([]string{"a", "c"}, merged)
// Result: map[string]int{"b": 2, "d": 4}
```

## Struct Operations (`rstruct`)

### Dynamic Field Access

```go
type User struct {
    Name string
    Age  int
    Address struct {
        Street string
        City   string
    }
}

user := User{
    Name: "Alice",
    Age:  25,
    Address: struct {
        Street string
        City   string
    }{
        Street: "123 Main St",
        City:   "New York",
    },
}

// Get field values
name, found := rstruct.Get(user, "Name") // "Alice", true
city, found := rstruct.Get(user, "Address.City") // "New York", true

// Safe access with defaults
age := rstruct.GetOrDefault(user, "Age", 0) // 25
email := rstruct.GetOrDefault(user, "Email", "unknown") // "unknown"

// Set field values
err := rstruct.Set(&user, "Age", 30) // user.Age = 30
err = rstruct.Set(&user, "Address.City", "Los Angeles")
```

### Field Information and Conversion

```go
// Get all field names
fields := rstruct.Fields(user) // []string{"Name", "Age", "Address"}

// Convert to/from map
userMap := rstruct.ToMap(user)
// Result: map[string]any{
//   "Name": "Alice",
//   "Age": 25,
//   "Address": map[string]any{...},
// }

// Create struct from map
newUser, ok := rstruct.FromMap[User](userMap)

// Pick specific fields
picked := rstruct.Pick(user, []string{"Name", "Age"})

// Omit specific fields
omitted := rstruct.Omit(user, []string{"Address"})
```

## Examples

The project includes comprehensive examples demonstrating all features:

```bash
# Run casting examples
go run _examples/casting/casting_composition.go

# Run indexing examples
go run _examples/indexing/index_by_example.go

# Run map operations examples
go run _examples/map_operations/map_example.go

# Run slice operations examples
go run _examples/slice_operations/slice_example.go

# Run struct operations examples
go run _examples/struct_operations/struct_example.go
```

See the `_examples/` directory for detailed documentation and working examples of all features.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
