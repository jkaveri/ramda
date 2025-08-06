package main

import (
	"fmt"
	"strings"

	"github.com/jkaveri/ramda/rmap"
)

func main() {
	fmt.Println("Map Operations Examples")
	fmt.Println("======================")

	// Example 1: MapKeys - Transform keys
	fmt.Println("\n1. MapKeys - Transform map keys:")
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("  Original: %v\n", original)

	uppercase := rmap.TransformKeys(func(k string) string { return strings.ToUpper(k) }, original)
	fmt.Printf("  Uppercase keys: %v\n", uppercase)

	// Example 2: MapValues - Transform values
	fmt.Println("\n2. MapValues - Transform map values:")
	fmt.Printf("  Original: %v\n", original)

	doubled := rmap.TransformValues(func(v int) int { return v * 2 }, original)
	fmt.Printf("  Doubled values: %v\n", doubled)

	stringValues := rmap.TransformValues(func(v int) string { return fmt.Sprintf("value_%d", v) }, original)
	fmt.Printf("  String values: %v\n", stringValues)

	// Example 3: MapEntries - Transform both keys and values
	fmt.Println("\n3. MapEntries - Transform both keys and values:")
	fmt.Printf("  Original: %v\n", original)

	transformed := rmap.TransformEntries(func(k string, v int) (string, string) {
		return strings.ToUpper(k), fmt.Sprintf("value_%d", v)
	}, original)
	fmt.Printf("  Transformed: %v\n", transformed)

	// Example 4: FilterMap - Filter map entries
	fmt.Println("\n4. FilterMap - Filter map entries:")
	numbers := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}
	fmt.Printf("  Original: %v\n", numbers)

	evenValues := rmap.Filter(func(k string, v int) bool { return v%2 == 0 }, numbers)
	fmt.Printf("  Even values: %v\n", evenValues)

	keysStartingWithA := rmap.Filter(func(k string, v int) bool { return strings.HasPrefix(k, "a") }, numbers)
	fmt.Printf("  Keys starting with 'a': %v\n", keysStartingWithA)

	// Example 5: Keys and Values - Extract keys and values
	fmt.Println("\n5. Keys and Values - Extract keys and values:")
	fmt.Printf("  Original: %v\n", original)

	keys := rmap.Keys(original)
	fmt.Printf("  Keys: %v\n", keys)

	values := rmap.Values(original)
	fmt.Printf("  Values: %v\n", values)

	// Example 6: Entries and FromEntries - Convert between map and entries
	fmt.Println("\n6. Entries and FromEntries - Convert between map and entries:")
	fmt.Printf("  Original: %v\n", original)

	entries := rmap.Entries(original)
	fmt.Printf("  Entries: %v\n", entries)

	reconstructed := rmap.FromEntries(entries)
	fmt.Printf("  Reconstructed: %v\n", reconstructed)

	// Example 7: Merge - Combine multiple maps
	fmt.Println("\n7. Merge - Combine multiple maps:")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}
	map3 := map[string]int{"d": 5, "e": 6}

	fmt.Printf("  Map1: %v\n", map1)
	fmt.Printf("  Map2: %v\n", map2)
	fmt.Printf("  Map3: %v\n", map3)

	merged := rmap.Merge(map1, map2, map3)
	fmt.Printf("  Merged: %v\n", merged)

	// Example 8: Pick and Omit - Selectively include/exclude keys
	fmt.Println("\n8. Pick and Omit - Selectively include/exclude keys:")
	config := map[string]string{
		"host":     "localhost",
		"port":     "8080",
		"database": "mydb",
		"username": "admin",
		"password": "secret",
		"timeout":  "30s",
	}
	fmt.Printf("  Original config: %v\n", config)

	// Pick only connection-related keys
	connectionKeys := []string{"host", "port", "database"}
	connectionConfig := rmap.Pick(connectionKeys, config)
	fmt.Printf("  Connection config (Pick): %v\n", connectionConfig)

	// Omit sensitive keys
	sensitiveKeys := []string{"password", "username"}
	publicConfig := rmap.Omit(sensitiveKeys, config)
	fmt.Printf("  Public config (Omit): %v\n", publicConfig)

	// Example 9: Has and Get - Check existence and retrieve values
	fmt.Println("\n9. Has and Get - Check existence and retrieve values:")
	fmt.Printf("  Original: %v\n", original)

	// Check if keys exist
	fmt.Printf("  Has 'a': %v\n", rmap.Has("a", original))
	fmt.Printf("  Has 'd': %v\n", rmap.Has("d", original))

	// Get values with existence check
	if value, exists := rmap.Get("b", original); exists {
		fmt.Printf("  Get 'b': %d (exists: %v)\n", value, exists)
	} else {
		fmt.Printf("  Get 'b': not found\n")
	}

	if value, exists := rmap.Get("x", original); exists {
		fmt.Printf("  Get 'x': %d (exists: %v)\n", value, exists)
	} else {
		fmt.Printf("  Get 'x': not found (exists: %v)\n", exists)
	}

	// Example 10: GetOrElse and GetOrElseFn - Safe value retrieval
	fmt.Println("\n10. GetOrElse and GetOrElseFn - Safe value retrieval:")
	fmt.Printf("  Original: %v\n", original)

	// GetOrElse with default value
	valueA := rmap.GetOrElse("a", 0, original)
	fmt.Printf("  GetOrElse 'a' with default 0: %d\n", valueA)

	valueX := rmap.GetOrElse("x", -1, original)
	fmt.Printf("  GetOrElse 'x' with default -1: %d\n", valueX)

	// GetOrElseFn with computed default
	valueB := rmap.GetOrElseFn("b", func() int { return 999 }, original)
	fmt.Printf("  GetOrElseFn 'b' with computed default: %d\n", valueB)

	valueY := rmap.GetOrElseFn("y", func() int { return 999 }, original)
	fmt.Printf("  GetOrElseFn 'y' with computed default: %d\n", valueY)

	// Example 11: Size and IsEmpty - Map properties
	fmt.Println("\n11. Size and IsEmpty - Map properties:")
	fmt.Printf("  Original: %v\n", original)
	fmt.Printf("  Size: %d\n", rmap.Size(original))
	fmt.Printf("  IsEmpty: %v\n", rmap.IsEmpty(original))

	emptyMap := map[string]int{}
	fmt.Printf("  Empty map: %v\n", emptyMap)
	fmt.Printf("  Size: %d\n", rmap.Size(emptyMap))
	fmt.Printf("  IsEmpty: %v\n", rmap.IsEmpty(emptyMap))

	// Example 12: Function composition with map operations
	fmt.Println("\n12. Function composition with map operations:")
	userData := map[string]interface{}{
		"name":   "John Doe",
		"age":    30,
		"email":  "john@example.com",
		"active": true,
		"score":  85.5,
	}
	fmt.Printf("  Original user data: %v\n", userData)

	// Transform to string values, then pick only string fields
	stringValues = rmap.TransformValues(func(v interface{}) string { return fmt.Sprintf("%v", v) }, userData)
	stringFields := rmap.Pick([]string{"name", "email"}, stringValues)
	fmt.Printf("  String fields only: %v\n", stringFields)

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
	fmt.Printf("  Numeric field keys: %v\n", numericKeys)
}
