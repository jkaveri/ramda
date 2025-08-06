package rmap

// TransformKeys transforms the keys of a map using the provided function.
// The values remain unchanged, but the keys are transformed according to the function.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	uppercase := TransformKeys(func(k string) string { return strings.ToUpper(k) }, original)
//	// Result: map[string]int{"A": 1, "B": 2, "C": 3}
func TransformKeys[K1, K2 comparable, V any](fn func(K1) K2, m map[K1]V) map[K2]V {
	result := make(map[K2]V, len(m))
	for k, v := range m {
		result[fn(k)] = v
	}
	return result
}

// TransformValues transforms the values of a map using the provided function.
// The keys remain unchanged, but the values are transformed according to the function.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	doubled := TransformValues(func(v int) int { return v * 2 }, original)
//	// Result: map[string]int{"a": 2, "b": 4, "c": 6}
func TransformValues[K comparable, V1, V2 any](fn func(V1) V2, m map[K]V1) map[K]V2 {
	result := make(map[K]V2, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// TransformEntries transforms both keys and values of a map using the provided function.
// The function receives both key and value, and returns a new key-value pair.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	transformed := TransformEntries(func(k string, v int) (string, string) {
//		return strings.ToUpper(k), fmt.Sprintf("value_%d", v)
//	}, original)
//	// Result: map[string]string{"A": "value_1", "B": "value_2", "C": "value_3"}
func TransformEntries[K1, K2 comparable, V1, V2 any](fn func(K1, V1) (K2, V2), m map[K1]V1) map[K2]V2 {
	result := make(map[K2]V2, len(m))
	for k, v := range m {
		newKey, newValue := fn(k, v)
		result[newKey] = newValue
	}
	return result
}

// Filter filters a map based on a predicate function that receives both key and value.
// Only entries where the predicate returns true are included in the result.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	evenValues := Filter(func(k string, v int) bool { return v%2 == 0 }, original)
//	// Result: map[string]int{"b": 2, "d": 4}
func Filter[K comparable, V any](fn func(K, V) bool, m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

// Keys returns all keys from a map as a slice.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := Keys(original)
//	// Result: []string{"a", "b", "c"} (order may vary)
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Values returns all values from a map as a slice.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := Values(original)
//	// Result: []int{1, 2, 3} (order may vary)
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// Entries returns all key-value pairs from a map as a slice of structs.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	entries := Entries(original)
//	// Result: []struct{Key string; Value int}{
//	//   {Key: "a", Value: 1},
//	//   {Key: "b", Value: 2},
//	//   {Key: "c", Value: 3},
//	// } (order may vary)
func Entries[K comparable, V any](m map[K]V) []struct {
	Key   K
	Value V
} {
	result := make([]struct {
		Key   K
		Value V
	}, 0, len(m))
	for k, v := range m {
		result = append(result, struct {
			Key   K
			Value V
		}{k, v})
	}
	return result
}

// FromEntries creates a map from a slice of key-value pairs.
// This is the inverse of Entries.
//
// Example:
//
//	entries := []struct{Key string; Value int}{
//		{Key: "a", Value: 1},
//		{Key: "b", Value: 2},
//		{Key: "c", Value: 3},
//	}
//	result := FromEntries(entries)
//	// Result: map[string]int{"a": 1, "b": 2, "c": 3}
func FromEntries[K comparable, V any](entries []struct {
	Key   K
	Value V
},
) map[K]V {
	result := make(map[K]V, len(entries))
	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}
	return result
}

// Merge combines multiple maps into a single map.
// If there are duplicate keys, the value from the later map takes precedence.
//
// Example:
//
//	map1 := map[string]int{"a": 1, "b": 2}
//	map2 := map[string]int{"b": 3, "c": 4}
//	merged := Merge(map1, map2)
//	// Result: map[string]int{"a": 1, "b": 3, "c": 4}
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Pick creates a new map containing only the specified keys from the original map.
// Keys that don't exist in the original map are ignored.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	picked := Pick([]string{"a", "c", "e"}, original)
//	// Result: map[string]int{"a": 1, "c": 3} (e is ignored as it doesn't exist)
func Pick[K comparable, V any](keys []K, m map[K]V) map[K]V {
	result := make(map[K]V)
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}
	return result
}

// Omit creates a new map containing all keys from the original map except the specified ones.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	omitted := Omit([]string{"b", "d"}, original)
//	// Result: map[string]int{"a": 1, "c": 3}
func Omit[K comparable, V any](keys []K, m map[K]V) map[K]V {
	keySet := make(map[K]struct{})
	for _, key := range keys {
		keySet[key] = struct{}{}
	}

	result := make(map[K]V)
	for k, v := range m {
		if _, exists := keySet[k]; !exists {
			result[k] = v
		}
	}
	return result
}

// Has checks if a key exists in the map.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	exists := Has("b", original) // true
//	notExists := Has("d", original) // false
func Has[K comparable, V any](key K, m map[K]V) bool {
	_, exists := m[key]
	return exists
}

// Get retrieves a value from the map and returns it along with a boolean indicating if the key exists.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	value, exists := Get("b", original) // value: 2, exists: true
//	value, exists := Get("d", original) // value: 0, exists: false
func Get[K comparable, V any](key K, m map[K]V) (V, bool) {
	value, exists := m[key]
	return value, exists
}

// GetOrElse retrieves a value from the map, returning a default value if the key doesn't exist.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	value := GetOrElse("b", 0, original) // 2
//	value := GetOrElse("d", 0, original) // 0
func GetOrElse[K comparable, V any](key K, defaultValue V, m map[K]V) V {
	if value, exists := m[key]; exists {
		return value
	}
	return defaultValue
}

// GetOrElseFn retrieves a value from the map, computing a default value if the key doesn't exist.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	value := GetOrElseFn("b", func() int { return 0 }, original) // 2
//	value := GetOrElseFn("d", func() int { return -1 }, original) // -1
func GetOrElseFn[K comparable, V any](key K, defaultFn func() V, m map[K]V) V {
	if value, exists := m[key]; exists {
		return value
	}
	return defaultFn()
}

// Size returns the number of key-value pairs in the map.
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	count := Size(original) // 3
func Size[K comparable, V any](m map[K]V) int {
	return len(m)
}

// IsEmpty checks if the map is empty (has no key-value pairs).
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2}
//	empty := IsEmpty(original) // false
//
//	emptyMap := map[string]int{}
//	isEmpty := IsEmpty(emptyMap) // true
func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}
