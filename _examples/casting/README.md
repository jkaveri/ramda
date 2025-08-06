# Casting Composition Examples

This directory contains examples demonstrating the casting utility functions in the Ramda Go library.

## Overview

The examples showcase how the new parameter order (function-first, data-last) enables better functional composition and currying.

## Key Benefits

1. **Function-First Parameter Order**: Functions come before data, making them more composable
2. **Safe Type Conversion**: All casting functions provide safe fallbacks for failed conversions
3. **Functional Composition**: Easy to build pipelines of transformations
4. **Error Handling**: Built-in error handling with customizable defaults

## Examples Included

### 1. Basic Casting
Demonstrates safe type conversion with default values:
```go
toIntWithDefault := func(s string) int {
    return ramda.Cast(strconv.Atoi, 0, s)
}
```

### 2. Function Composition
Shows how to compose transformation functions:
```go
doubleThenString := func(val int) string {
    doubled := curriedDouble(val)
    return curriedToString(doubled)
}
```

### 3. Error-Safe Transformations
Demonstrates handling conversion errors gracefully:
```go
parseInt := func(s string) int {
    return ramda.TransformWithError(strconv.Atoi, 0, s)
}
```

### 4. Type Assertion with Function Defaults
Shows how to use type assertion with function-generated default values:
```go
safeStringExtractor := func(val interface{}) string {
    return ramda.AsWithDefaultFn(func() string { return "DEFAULT" }, val)
}
```

### 5. Pipeline Processing
Shows how to build complex data processing pipelines:
```go
processString := func(s string) string {
    asInt := toInt(s)
    doubled := curriedDouble(asInt)
    return curriedToString(doubled)
}
```

### 6. Pointer Casting
Demonstrates safe pointer handling and optional value management:
```go
// Safe pointer dereferencing
value := ramda.FromPtr(optionalInt) // Returns 0 if nil

// With custom defaults using Default
value := ramda.Default(ramda.FromPtr(optionalInt), 100) // Returns 100 if nil

// With function-generated defaults using DefaultFn
value := ramda.DefaultFn(func() int { return 999 }, ramda.FromPtr(optionalInt))

// Creating pointers
ptr := ramda.ToPtr(42) // Always returns a pointer
optional := ramda.NilIfEmpty(42) // Returns nil for zero values
```

## Running the Example

```bash
go run examples/casting/casting_composition.go
```

## Expected Output

```
Cast with new parameter order: 123, 0
Composed transformation: Result: 10
TransformWithError with new parameter order: 456, 0
Pipeline result for '1': Result: 2
Pipeline result for '2': Result: 4
Pipeline result for 'abc': Result: 0
Pipeline result for '3': Result: 6
Complex pipeline: Processed: 123.00

Pointer casting examples:
FromPtr(nil int): 0
FromPtr(nil string):
FromPtr(&42): 42
FromPtr(&hello): hello
Default(FromPtr(nil), 100): 100
Default(FromPtr(nil), default): default
Default(FromPtr(&42), 100): 42
DefaultFn(fn, FromPtr(nil)): 999
DefaultFn(fn, FromPtr(&42)): 42
ToPtr(42): 0x1400000e1b0
ToPtr(0): 0x1400000e1b8
ToPtr(hello): 0x14000010100
ToPtr(empty): 0x14000010108
ToPtr(true): 0x1400000e1c0
ToPtr(false): 0x1400000e1c8
NilIfEmpty(42): 0x1400000e1d0
NilIfEmpty(0): <nil>
NilIfEmpty(hello): 0x14000010110
NilIfEmpty(empty): <nil>
NilIfEmpty(true): 0x1400000e1d8
NilIfEmpty(false): <nil>

Pipeline with pointer handling:
  Data[0] (123): Processed: 246
  Data[1] (abc): Processed: 0
  Data[2] (nil): Processed: 0
  Data[3] (456): Processed: 912
```

## Key Functions Used

- `ramda.Cast` - Safe type conversion with default value
- `ramda.Transform` - Apply transformation function
- `ramda.TransformWithError` - Safe transformation with error handling
- `ramda.Compose` - Function composition
- `ramda.Curry` - Function currying (demonstrated conceptually)
- `ramda.FromPtr` - Safe pointer dereferencing
- `ramda.Default` - Provide default value if zero
- `ramda.DefaultFn` - Provide function-generated default if zero
- `ramda.ToPtr` - Always create pointer
- `ramda.NilIfEmpty` - Create optional pointer (nil for zero values)
