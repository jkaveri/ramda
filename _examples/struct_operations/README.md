# Struct Operations Examples

This directory contains examples demonstrating the struct operations functionality in the ramda library.

## Examples

### Basic Struct Operations (`struct_example.go`)
Demonstrates basic struct field access and manipulation using the standalone functions:
- `Get()` - Retrieve field values
- `Set()` - Set field values (returns error for success/failure)
- `Has()` - Check field existence
- `GetOrDefault()` - Get with default fallback
- `Fields()` - Get all field names

- `ToMap()` - Convert struct to map
- `FromMap()` - Create struct from map
- `Pick()` - Select specific fields
- `Omit()` - Exclude specific fields
- `Merge()` - Combine multiple structs
- `Clone()` - Create deep copy

### Set Function Example (`set_example.go`)
Demonstrates the `Set()` function that returns an error:
- Shows how to handle success and error cases
- Demonstrates nested field setting
- Shows various error scenarios (non-existing fields, type mismatches, etc.)

### Nested Fields (`nested_fields_example.go`)
Shows how to work with nested struct fields using dot notation:
- Access nested fields like `"Address.City"`
- Set nested field values
- Check nested field existence

### Field Conflicts (`field_conflict_example.go`)
Demonstrates how field name conflicts are handled:
- Shows behavior when embedded and container structs have fields with the same name
- Demonstrates explicit dot notation for accessing embedded fields
- Shows precedence rules for field access

### Struct Reflection Debug (`cache_debug.go`)
Shows the internal structure of structs and how field access works:
- Demonstrates embedded struct field access using dot notation
- Shows that direct access to embedded fields doesn't work
- Explains the difference between direct and nested field access

### Nested Cache Debug (`nested_cache_debug.go`)
Demonstrates nested field access and setting:
- Shows how to access deeply nested fields
- Demonstrates setting nested field values
- Tests field existence checks

## Features

- **Case-sensitive field names**: Field names must match exactly
- **Nested field support**: `"Address.City"`, `"User.Profile.Email"`
- **Embedded struct support**: Use explicit dot notation (e.g., `"User.Name"`)
- **Pointer support**: Works with both value and pointer types
- **Error handling**: `Set()` returns detailed error messages

## Important Notes

### Embedded Struct Access
The implementation requires explicit dot notation to access embedded struct fields:

```go
type User struct {
    Name string
    Age  int
}

type Person struct {
    User   // Embedded struct
    Active bool
}

person := Person{User: User{Name: "Alice", Age: 25}, Active: true}

// ❌ This doesn't work
name, found := rstruct.Get(person, "Name") // found = false

// ✅ This works
name, found := rstruct.Get(person, "User.Name") // found = true
```

### Field Name Conflicts
When embedded and container structs have fields with the same name, use explicit dot notation to access the embedded field:

```go
type Container struct {
    Embedded        // Embedded struct with Name field
    Name     string // Container also has Name field
}

// Access container field
name, found := rstruct.Get(container, "Name")

// Access embedded field
embeddedName, found := rstruct.Get(container, "Embedded.Name")
```

## Running the Examples

```bash
# Run basic struct operations
go run struct_example.go

# Run nested fields example
go run nested_fields_example.go

# Run Set function example (returns error)
go run set_example.go

# Run field conflict example
go run field_conflict_example.go

# Run struct reflection debug
go run cache_debug.go

# Run nested cache debug
go run nested_cache_debug.go
```

## Function Signatures

```go
// Get field value
Get[T any](obj T, fieldName string) (any, bool)

// Set field value (returns error)
Set[T any](obj T, fieldName string, value any) error

// Check field existence
Has[T any](obj T, fieldName string) bool

// Get with default value
GetOrDefault[T any](obj T, fieldName string, defaultValue any) any

// Get all field names
Fields[T any](obj T) []string



// Convert struct to map
ToMap[T any](obj T) map[string]any

// Create struct from map
FromMap[T any](data map[string]any) (T, bool)

// Select specific fields
Pick[T any](obj T, fields []string) T

// Exclude specific fields
Omit[T any](obj T, fields []string) T

// Combine multiple structs
Merge[T any](structs ...T) T

// Create deep copy
Clone[T any](obj T) T
```
