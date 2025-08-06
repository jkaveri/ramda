# Ramda Go Examples

This directory contains examples demonstrating the various features and utilities provided by the Ramda Go library.

## Available Examples

### Casting Examples (`casting/`)
Demonstrates the casting utility functions and how the function-first parameter order enables better functional composition.

**Key Features:**
- Safe type conversion with default values
- Function composition and currying
- Error handling with graceful fallbacks
- Building data processing pipelines

**Run the example:**
```bash
go run _examples/casting/casting_composition.go
```

### Indexing Examples (`indexing/`)
Demonstrates the `IndexBy` function for creating lookup maps from slices.

**Key Features:**
- Flexible key generation from slice elements
- Overwrite behavior for duplicate keys
- Type-safe generic implementation
- Composable with other functional utilities

**Run the example:**
```bash
go run _examples/indexing/index_by_example.go
```

### Map Operations Examples (`map_operations/`)
Demonstrates comprehensive map utility functions for transformation, filtering, and safe access.

**Key Features:**
- Transform map keys, values, and entries
- Filter map entries based on predicates
- Safe value retrieval with defaults
- Map composition and merging
- Extract keys, values, and entries

**Run the example:**
```bash
go run _examples/map_operations/map_example.go
```

### Slice Operations Examples (`slice_operations/`)
Demonstrates comprehensive slice utility functions for transformation, filtering, and manipulation.

**Key Features:**
- Transform, filter, and reduce slice elements
- Search and query functions (Find, Any, All)
- Slice manipulation (Take, Drop, Reverse, Unique)
- Combination and grouping (Zip, GroupBy, SortBy)
- Conversion functions (IndexBy, ToSet)

**Run the example:**
```bash
go run _examples/slice_operations/slice_example.go
```

### Struct Operations Examples (`struct_operations/`)
Demonstrates comprehensive struct utility functions for dynamic field access and manipulation.

**Key Features:**
- Dynamic field access and modification (Get, Set, Has)
- Safe field access with defaults (GetOrDefault)
- Field information and conversion (Fields, ToMap, FromMap)
- Field selection and combination (Pick, Omit, Merge)
- Utility functions (Clone, IsEmpty)

**Run the example:**
```bash
go run _examples/struct_operations/struct_example.go
```

## Project Structure

```
_examples/
├── README.md              # This file
├── casting/              # Casting utility examples
│   ├── README.md         # Detailed casting examples documentation
│   └── casting_composition.go  # Main casting example
├── indexing/             # Indexing utility examples
│   ├── README.md         # Detailed indexing examples documentation
│   └── index_by_example.go  # Main indexing example
├── map_operations/       # Map utility examples
│   ├── README.md         # Detailed map operations documentation
│   └── map_example.go    # Main map operations example
├── slice_operations/     # Slice utility examples
│   ├── README.md         # Detailed slice operations documentation
│   └── slice_example.go  # Main slice operations example
└── struct_operations/    # Struct utility examples
    ├── README.md         # Detailed struct operations documentation
    └── struct_example.go # Main struct operations example
```

## Running Examples

All examples can be run from the project root directory using:

```bash
go run _examples/[subdirectory]/[filename].go
```

## Contributing Examples

When adding new examples:

1. Create a new subdirectory for your example category
2. Include a README.md explaining the example
3. Ensure the example demonstrates key functional programming concepts
4. Test that the example runs successfully
5. Update this main README.md to include your new example
