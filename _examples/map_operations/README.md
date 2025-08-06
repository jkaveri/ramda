# Map Operations Examples

This directory contains examples demonstrating the map utility functions in the Ramda Go library.

## Overview

The map operations provide functional utilities for working with Go maps, including transformation, filtering, merging, and safe value retrieval operations.

## Key Functions

### Transformation Functions
- `rmap.TransformKeys` - Transform map keys using a function
- `rmap.TransformValues` - Transform map values using a function
- `rmap.TransformEntries` - Transform both keys and values using a function

### Filtering and Selection
- `rmap.Filter` - Filter map entries based on a predicate
- `rmap.Pick` - Create a new map with only specified keys
- `rmap.Omit` - Create a new map excluding specified keys

### Extraction and Conversion
- `rmap.Keys` - Extract all keys as a slice
- `rmap.Values` - Extract all values as a slice
- `rmap.Entries` - Convert map to slice of key-value pairs
- `rmap.FromEntries` - Convert slice of key-value pairs to map

### Combination
- `rmap.Merge` - Combine multiple maps into one

### Safe Access
- `rmap.Has` - Check if a key exists
- `rmap.Get` - Retrieve value with existence check
- `rmap.GetOrElse` - Retrieve value with default fallback
- `rmap.GetOrElseFn` - Retrieve value with computed default

### Properties
- `rmap.Size` - Get number of key-value pairs
- `rmap.IsEmpty` - Check if map is empty

## Examples

### 1. TransformKeys - Transform map keys
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
uppercase := rmap.TransformKeys(func(k string) string { return strings.ToUpper(k) }, original)
// Result: map[string]int{"A": 1, "B": 2, "C": 3}
```

### 2. TransformValues - Transform map values
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
doubled := rmap.TransformValues(func(v int) int { return v * 2 }, original)
// Result: map[string]int{"a": 2, "b": 4, "c": 6}
```

### 3. TransformEntries - Transform both keys and values
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
transformed := rmap.TransformEntries(func(k string, v int) (string, string) {
    return strings.ToUpper(k), fmt.Sprintf("value_%d", v)
}, original)
// Result: map[string]string{"A": "value_1", "B": "value_2", "C": "value_3"}
```

### 4. Filter - Filter map entries
```go
numbers := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}
evenValues := rmap.Filter(func(k string, v int) bool { return v%2 == 0 }, numbers)
// Result: map[string]int{"b": 2, "d": 4, "f": 6}
```

### 5. Keys and Values - Extract keys and values
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
keys := rmap.Keys(original)     // Result: []string{"a", "b", "c"}
values := rmap.Values(original) // Result: []int{1, 2, 3}
```

### 6. Entries and FromEntries - Convert between map and entries
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
entries := rmap.Entries(original)
// Result: []struct{Key string; Value int}{{Key: "a", Value: 1}, ...}

reconstructed := rmap.FromEntries(entries)
// Result: map[string]int{"a": 1, "b": 2, "c": 3}
```

### 7. Merge - Combine multiple maps
```go
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"b": 3, "c": 4}
map3 := map[string]int{"d": 5, "e": 6}
merged := rmap.Merge(map1, map2, map3)
// Result: map[string]int{"a": 1, "b": 3, "c": 4, "d": 5, "e": 6}
```

### 8. Pick and Omit - Selectively include/exclude keys
```go
config := map[string]string{
    "host": "localhost", "port": "8080", "database": "mydb",
    "username": "admin", "password": "secret", "timeout": "30s",
}

// Pick only connection-related keys
connectionKeys := []string{"host", "port", "database"}
connectionConfig := rmap.Pick(connectionKeys, config)
// Result: map[string]string{"host": "localhost", "port": "8080", "database": "mydb"}

// Omit sensitive keys
sensitiveKeys := []string{"password", "username"}
publicConfig := rmap.Omit(sensitiveKeys, config)
// Result: map[string]string{"host": "localhost", "port": "8080", "database": "mydb", "timeout": "30s"}
```

### 9. Has and Get - Check existence and retrieve values
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}

// Check if keys exist
exists := rmap.Has("b", original) // true
notExists := rmap.Has("d", original) // false

// Get values with existence check
if value, exists := rmap.Get("b", original); exists {
    fmt.Printf("Value: %d\n", value) // Value: 2
}
```

### 10. GetOrElse and GetOrElseFn - Safe value retrieval
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}

// GetOrElse with default value
value := rmap.GetOrElse("x", 0, original) // 0

// GetOrElseFn with computed default
value := rmap.GetOrElseFn("y", func() int { return 999 }, original) // 999
```

### 11. Size and IsEmpty - Map properties
```go
original := map[string]int{"a": 1, "b": 2, "c": 3}
size := rmap.Size(original)        // 3
isEmpty := rmap.IsEmpty(original)  // false

emptyMap := map[string]int{}
size := rmap.Size(emptyMap)        // 0
isEmpty := rmap.IsEmpty(emptyMap)  // true
```

### 12. Function composition with map operations
```go
userData := map[string]interface{}{
    "name": "John Doe", "age": 30, "email": "john@example.com",
    "active": true, "score": 85.5,
}

// Transform to string values, then pick only string fields
stringValues := rmap.TransformValues(func(v interface{}) string { return fmt.Sprintf("%v", v) }, userData)
stringFields := rmap.Pick([]string{"name", "email"}, stringValues)

// Filter numeric values, then get their keys
numericEntries := rmap.Filter(func(k string, v interface{}) bool {
    switch v.(type) {
    case int, float64:
        return true
    default:
        return false
    }
}, userData)
numericKeys := rmap.Keys(numericEntries)
```

## Expected Output

```
Map Operations Examples
======================

1. MapKeys - Transform map keys:
  Original: map[a:1 b:2 c:3]
  Uppercase keys: map[A:1 B:2 C:3]

2. MapValues - Transform map values:
  Original: map[a:1 b:2 c:3]
  Doubled values: map[a:2 b:4 c:6]
  String values: map[a:value_1 b:value_2 c:value_3]

3. MapEntries - Transform both keys and values:
  Original: map[a:1 b:2 c:3]
  Transformed: map[A:value_1 B:value_2 C:value_3]

4. FilterMap - Filter map entries:
  Original: map[a:1 b:2 c:3 d:4 e:5 f:6]
  Even values: map[b:2 d:4 f:6]
  Keys starting with 'a': map[a:1]

5. Keys and Values - Extract keys and values:
  Original: map[a:1 b:2 c:3]
  Keys: [a b c]
  Values: [2 3 1]

6. Entries and FromEntries - Convert between map and entries:
  Original: map[a:1 b:2 c:3]
  Entries: [{b 2} {c 3} {a 1}]
  Reconstructed: map[a:1 b:2 c:3]

7. Merge - Combine multiple maps:
  Map1: map[a:1 b:2]
  Map2: map[b:3 c:4]
  Map3: map[d:5 e:6]
  Merged: map[a:1 b:3 c:4 d:5 e:6]

8. Pick and Omit - Selectively include/exclude keys:
  Original config: map[database:mydb host:localhost password:secret port:8080 timeout:30s username:admin]
  Connection config (Pick): map[database:mydb host:localhost port:8080]
  Public config (Omit): map[database:mydb host:localhost port:8080 timeout:30s]

9. Has and Get - Check existence and retrieve values:
  Original: map[a:1 b:2 c:3]
  Has 'a': true
  Has 'd': false
  Get 'b': 2 (exists: true)
  Get 'x': not found (exists: false)

10. GetOrElse and GetOrElseFn - Safe value retrieval:
  Original: map[a:1 b:2 c:3]
  GetOrElse 'a' with default 0: 1
  GetOrElse 'x' with default -1: -1
  GetOrElseFn 'b' with computed default: 2
  GetOrElseFn 'y' with computed default: 999

11. Size and IsEmpty - Map properties:
  Original: map[a:1 b:2 c:3]
  Size: 3
  IsEmpty: false
  Empty map: map[]
  Size: 0
  IsEmpty: true

12. Function composition with map operations:
  Original user data: map[active:true age:30 email:john@example.com name:John Doe score:85.5]
  String fields only: map[email:john@example.com name:John Doe]
  Numeric field keys: [age score]
```

## Use Cases

- **Configuration Management**: Pick/omit sensitive fields, merge configs
- **Data Transformation**: Convert between different data formats
- **API Response Processing**: Extract specific fields, transform data
- **Data Validation**: Check required fields, filter invalid entries
- **Safe Data Access**: Handle missing keys gracefully
- **Data Analysis**: Extract keys/values for processing
- **Map Composition**: Build complex maps from simpler ones
- **Type Conversion**: Transform map types for different contexts

## Key Features

1. **Type Safety**: Full generic type safety with compile-time checks
2. **Immutability**: Functions return new maps rather than modifying inputs
3. **Composability**: Functions can be easily combined and chained
4. **Safe Access**: Built-in safety for missing keys
5. **Flexible Transformation**: Support for key, value, and entry transformations
6. **Efficient Operations**: Optimized for common map operations
7. **Functional Style**: Consistent with functional programming principles
