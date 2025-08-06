# IndexBy Function Examples

This directory contains examples demonstrating the `IndexBy` function from the Ramda Go library.

## Overview

The `IndexBy` function creates a map from a slice using a function to generate keys. Each element in the slice becomes a value in the map, with the key determined by applying the provided function to the element.

## Key Features

1. **Flexible Key Generation**: Use any function to generate keys from slice elements
2. **Overwrite Behavior**: If multiple elements produce the same key, the last element overwrites previous ones
3. **Type Safety**: Full generic type safety with compile-time checks
4. **Composable**: Can be easily combined with other functional utilities
5. **Set Operations**: Create efficient sets for membership checking
6. **Automatic Deduplication**: Duplicate values are automatically handled

## Function Signatures

```go
func IndexBy[T any, K comparable](fn func(T) K, slice []T) map[K]T
func ToSet[T comparable](slice []T) map[T]struct{}
```

## Examples Included

### 1. Basic Indexing
Demonstrates simple indexing by ID:
```go
users := []User{{ID: "1", Name: "Alice"}, {ID: "2", Name: "Bob"}}
userIndex := ramda.IndexBy(func(u User) string { return u.ID }, users)
```

### 2. Product Catalog Indexing
Shows indexing products by SKU for efficient lookup:
```go
products := []Product{{SKU: "PROD-001", Name: "Laptop"}}
productIndex := ramda.IndexBy(func(p Product) string { return p.SKU }, products)
```

### 3. Category Grouping
Demonstrates overwrite behavior when multiple items have the same key:
```go
categoryIndex := ramda.IndexBy(func(p Product) string { return p.Category }, products)
```

### 4. Age Group Classification
Shows complex key generation logic:
```go
ageGroupIndex := ramda.IndexBy(func(u User) string {
    if u.Age < 30 { return "Young" }
    else if u.Age < 40 { return "Middle-aged" }
    return "Senior"
}, users)
```

### 5. Price Range Classification
Demonstrates conditional key generation:
```go
priceRangeIndex := ramda.IndexBy(func(p Product) string {
    if p.Price < 50 { return "Budget" }
    else if p.Price < 200 { return "Mid-range" }
    return "Premium"
}, products)
```

### 6. Overwrite Behavior
Shows how duplicate keys are handled:
```go
duplicateUsers := []User{
    {ID: "1", Name: "Alice"},
    {ID: "1", Name: "AliceUpdated"}, // Same ID, different name
}
duplicateIndex := ramda.IndexBy(func(u User) string { return u.ID }, duplicateUsers)
```

### 7. Function Composition
Demonstrates combining IndexBy with other functions:
```go
expensiveProducts := ramda.Filter(func(p Product) bool { return p.Price > 100 }, products)
expensiveByCategory := ramda.IndexBy(func(p Product) string { return p.Category }, expensiveProducts)
```

### 8. Set Operations with ToSet
Demonstrates creating sets for membership checking:
```go
numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
numberSet := ramda.ToSet(numbers)

// Check membership
if _, exists := numberSet[3]; exists {
    fmt.Println("3 is in the set")
}
```

### 9. Combining ToSet with Other Functions
Shows how to create sets from filtered or transformed data:
```go
// Filter even numbers, then create a set
evenNumbers := ramda.Filter(func(n int) bool { return n%2 == 0 }, numbers)
evenNumberSet := ramda.ToSet(evenNumbers)

// Create a set of product categories
categorySet := ramda.ToSet(ramda.Map(func(p Product) string { return p.Category }, products))
```

## Running the Example

```bash
go run _examples/indexing/index_by_example.go
```

## Expected Output

```
IndexBy Function Examples
========================

1. Users indexed by ID:
  1: Alice (Age: 25)
  2: Bob (Age: 30)
  3: Charlie (Age: 35)
  4: Diana (Age: 28)

2. Products indexed by SKU:
  PROD-001: Laptop ($999.99) - Electronics
  PROD-002: Mouse ($29.99) - Electronics
  PROD-003: Desk ($199.99) - Furniture
  PROD-004: Chair ($149.99) - Furniture

3. Products indexed by Category (shows overwrite behavior):
  Electronics: Mouse ($29.99)
  Furniture: Chair ($149.99)

4. Users indexed by age group:
  Young: Diana (Age: 28)
  Middle-aged: Charlie (Age: 35)

5. Products indexed by price range:
  Budget: Mouse ($29.99)
  Mid-range: Chair ($149.99)
  Premium: Laptop ($999.99)

6. Demonstrating overwrite behavior:
  1: AliceUpdated (Age: 26)
  2: Bob (Age: 30)

7. Combining IndexBy with other functions:
  Expensive products by category:
    Electronics: Laptop ($999.99)
    Furniture: Chair ($149.99)

8. Set operations with ToSet:
  Number set (unique values):
    5
    1
    3
    2
    4
  Word set (unique values):
    hello
    world
    go
    programming
    golang
  Membership checking:
    1 is in the set
    3 is in the set
    5 is in the set
    7 is NOT in the set
    9 is NOT in the set
  Word membership checking:
    'hello' is in the set
    'go' is in the set
    'python' is NOT in the set
    'golang' is in the set

9. Combining ToSet with other functions:
  Even numbers set:
    2
    4
  Product categories set:
    Furniture
    Electronics
```

## Use Cases

- **Database Lookups**: Create efficient lookup tables from database results
- **API Response Processing**: Index API responses by ID for quick access
- **Configuration Management**: Index configuration objects by key
- **Data Analysis**: Group data by categories or ranges
- **Caching**: Create in-memory indexes for frequently accessed data
- **Membership Checking**: Create sets for efficient existence testing
- **Deduplication**: Remove duplicate values from slices
- **Set Operations**: Perform set-based operations on data

## Key Functions Used

- `ramda.IndexBy` - Create map from slice using key function
- `ramda.ToSet` - Create set-like map for membership checking
- `ramda.Filter` - Filter slice before indexing
- `ramda.Map` - Transform slice before creating set
- Function composition patterns
