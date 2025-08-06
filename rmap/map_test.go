package rmap

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransformKeys(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test uppercase transformation
	uppercase := TransformKeys(func(k string) string { return strings.ToUpper(k) }, original)

	if len(uppercase) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(uppercase))
	}

	// Check specific entries
	if val, exists := uppercase["A"]; !exists || val != 1 {
		t.Errorf("Expected A: 1, got %v", val)
	}
	if val, exists := uppercase["B"]; !exists || val != 2 {
		t.Errorf("Expected B: 2, got %v", val)
	}
	if val, exists := uppercase["C"]; !exists || val != 3 {
		t.Errorf("Expected C: 3, got %v", val)
	}

	// Test with different key types
	intMap := map[int]string{1: "one", 2: "two", 3: "three"}
	stringKeys := TransformKeys(func(k int) string { return fmt.Sprintf("key_%d", k) }, intMap)

	if len(stringKeys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(stringKeys))
	}

	if val, exists := stringKeys["key_1"]; !exists || val != "one" {
		t.Errorf("Expected key_1: one, got %v", val)
	}
}

func TestTransformValues(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test doubling values
	doubled := TransformValues(func(v int) int { return v * 2 }, original)

	if len(doubled) != 3 {
		t.Errorf("Expected 3 values, got %d", len(doubled))
	}

	// Check specific entries
	if val, exists := doubled["a"]; !exists || val != 2 {
		t.Errorf("Expected a: 2, got %v", val)
	}
	if val, exists := doubled["b"]; !exists || val != 4 {
		t.Errorf("Expected b: 4, got %v", val)
	}
	if val, exists := doubled["c"]; !exists || val != 6 {
		t.Errorf("Expected c: 6, got %v", val)
	}

	// Test with different value types
	stringMap := map[string]int{"a": 1, "b": 2, "c": 3}
	stringValues := TransformValues(func(v int) string { return fmt.Sprintf("value_%d", v) }, stringMap)

	if len(stringValues) != 3 {
		t.Errorf("Expected 3 values, got %d", len(stringValues))
	}

	if val, exists := stringValues["a"]; !exists || val != "value_1" {
		t.Errorf("Expected a: value_1, got %v", val)
	}
}

func TestTransformEntries(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test transforming both keys and values
	transformed := TransformEntries(func(k string, v int) (string, string) {
		return strings.ToUpper(k), fmt.Sprintf("value_%d", v)
	}, original)

	if len(transformed) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(transformed))
	}

	// Check specific entries
	if val, exists := transformed["A"]; !exists || val != "value_1" {
		t.Errorf("Expected A: value_1, got %v", val)
	}
	if val, exists := transformed["B"]; !exists || val != "value_2" {
		t.Errorf("Expected B: value_2, got %v", val)
	}
	if val, exists := transformed["C"]; !exists || val != "value_3" {
		t.Errorf("Expected C: value_3, got %v", val)
	}
}

func TestFilter(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Test filtering even values
	evenValues := Filter(func(k string, v int) bool { return v%2 == 0 }, original)

	if len(evenValues) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(evenValues))
	}

	// Check specific entries
	if val, exists := evenValues["b"]; !exists || val != 2 {
		t.Errorf("Expected b: 2, got %v", val)
	}
	if val, exists := evenValues["d"]; !exists || val != 4 {
		t.Errorf("Expected d: 4, got %v", val)
	}

	// Test filtering by key
	keysStartingWithA := Filter(func(k string, v int) bool { return strings.HasPrefix(k, "a") }, original)

	if len(keysStartingWithA) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(keysStartingWithA))
	}

	if val, exists := keysStartingWithA["a"]; !exists || val != 1 {
		t.Errorf("Expected a: 1, got %v", val)
	}
}

func TestKeys(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := Keys(original)

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Check that all expected keys are present
	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key: %s", key)
		}
	}

	// Test with empty map
	emptyMap := map[string]int{}
	emptyKeys := Keys(emptyMap)

	if len(emptyKeys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(emptyKeys))
	}
}

func TestValues(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	values := Values(original)

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Check that all expected values are present
	expectedValues := map[int]bool{1: true, 2: true, 3: true}
	for _, value := range values {
		if !expectedValues[value] {
			t.Errorf("Unexpected value: %d", value)
		}
	}

	// Test with empty map
	emptyMap := map[string]int{}
	emptyValues := Values(emptyMap)

	if len(emptyValues) != 0 {
		t.Errorf("Expected 0 values, got %d", len(emptyValues))
	}
}

func TestEntries(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	entries := Entries(original)

	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}

	// Check that all expected entries are present
	expectedEntries := map[string]int{"a": 1, "b": 2, "c": 3}
	for _, entry := range entries {
		if expectedValue, exists := expectedEntries[entry.Key]; !exists || expectedValue != entry.Value {
			t.Errorf("Unexpected entry: %s: %d", entry.Key, entry.Value)
		}
	}

	// Test with empty map
	emptyMap := map[string]int{}
	emptyEntries := Entries(emptyMap)

	if len(emptyEntries) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(emptyEntries))
	}
}

func TestFromEntries(t *testing.T) {
	entries := []struct {
		Key   string
		Value int
	}{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
	}

	result := FromEntries(entries)

	if len(result) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(result))
	}

	// Check specific entries
	if val, exists := result["a"]; !exists || val != 1 {
		t.Errorf("Expected a: 1, got %v", val)
	}
	if val, exists := result["b"]; !exists || val != 2 {
		t.Errorf("Expected b: 2, got %v", val)
	}
	if val, exists := result["c"]; !exists || val != 3 {
		t.Errorf("Expected c: 3, got %v", val)
	}

	// Test with empty slice
	emptyEntries := []struct {
		Key   string
		Value int
	}{}
	emptyResult := FromEntries(emptyEntries)

	if len(emptyResult) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(emptyResult))
	}
}

func TestMerge(t *testing.T) {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}
	map3 := map[string]int{"d": 5}

	merged := Merge(map1, map2, map3)

	if len(merged) != 4 {
		t.Errorf("Expected 4 entries, got %d", len(merged))
	}

	// Check specific entries
	if val, exists := merged["a"]; !exists || val != 1 {
		t.Errorf("Expected a: 1, got %v", val)
	}
	if val, exists := merged["b"]; !exists || val != 3 {
		t.Errorf("Expected b: 3 (overwritten), got %v", val)
	}
	if val, exists := merged["c"]; !exists || val != 4 {
		t.Errorf("Expected c: 4, got %v", val)
	}
	if val, exists := merged["d"]; !exists || val != 5 {
		t.Errorf("Expected d: 5, got %v", val)
	}

	// Test with single map
	singleMap := map[string]int{"a": 1}
	result := Merge(singleMap)

	if len(result) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(result))
	}

	if val, exists := result["a"]; !exists || val != 1 {
		t.Errorf("Expected a: 1, got %v", val)
	}

	// Test with no maps
	emptyResult := Merge[string, int]()

	if len(emptyResult) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(emptyResult))
	}
}

func TestPick(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	keys := []string{"a", "c", "e"}

	picked := Pick(keys, original)

	if len(picked) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(picked))
	}

	// Check specific entries
	if val, exists := picked["a"]; !exists || val != 1 {
		t.Errorf("Expected a: 1, got %v", val)
	}
	if val, exists := picked["c"]; !exists || val != 3 {
		t.Errorf("Expected c: 3, got %v", val)
	}

	// Check that non-existent key is not included
	if _, exists := picked["e"]; exists {
		t.Error("Expected e to not exist in picked map")
	}

	// Test with empty keys slice
	emptyKeys := []string{}
	emptyPicked := Pick(emptyKeys, original)

	if len(emptyPicked) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(emptyPicked))
	}
}

func TestOmit(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	keysToOmit := []string{"b", "d"}

	omitted := Omit(keysToOmit, original)

	if len(omitted) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(omitted))
	}

	// Check specific entries
	if val, exists := omitted["a"]; !exists || val != 1 {
		t.Errorf("Expected a: 1, got %v", val)
	}
	if val, exists := omitted["c"]; !exists || val != 3 {
		t.Errorf("Expected c: 3, got %v", val)
	}

	// Check that omitted keys are not included
	if _, exists := omitted["b"]; exists {
		t.Error("Expected b to not exist in omitted map")
	}
	if _, exists := omitted["d"]; exists {
		t.Error("Expected d to not exist in omitted map")
	}

	// Test with empty keys slice
	emptyKeys := []string{}
	noOmission := Omit(emptyKeys, original)

	if len(noOmission) != 4 {
		t.Errorf("Expected 4 entries, got %d", len(noOmission))
	}
}

func TestHas(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test existing keys
	if !Has("a", original) {
		t.Error("Expected a to exist in map")
	}
	if !Has("b", original) {
		t.Error("Expected b to exist in map")
	}
	if !Has("c", original) {
		t.Error("Expected c to exist in map")
	}

	// Test non-existing keys
	if Has("d", original) {
		t.Error("Expected d to not exist in map")
	}
	if Has("x", original) {
		t.Error("Expected x to not exist in map")
	}
}

func TestGet(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test existing keys
	if value, exists := Get("a", original); !exists || value != 1 {
		t.Errorf("Expected a: 1, exists: true, got %v, exists: %v", value, exists)
	}
	if value, exists := Get("b", original); !exists || value != 2 {
		t.Errorf("Expected b: 2, exists: true, got %v, exists: %v", value, exists)
	}
	if value, exists := Get("c", original); !exists || value != 3 {
		t.Errorf("Expected c: 3, exists: true, got %v, exists: %v", value, exists)
	}

	// Test non-existing keys
	if value, exists := Get("d", original); exists {
		t.Errorf("Expected d to not exist, got %v, exists: %v", value, exists)
	}
	if value, exists := Get("x", original); exists {
		t.Errorf("Expected x to not exist, got %v, exists: %v", value, exists)
	}
}

func TestGetOrElse(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test existing keys
	if value := GetOrElse("a", 0, original); value != 1 {
		t.Errorf("Expected 1, got %v", value)
	}
	if value := GetOrElse("b", 0, original); value != 2 {
		t.Errorf("Expected 2, got %v", value)
	}
	if value := GetOrElse("c", 0, original); value != 3 {
		t.Errorf("Expected 3, got %v", value)
	}

	// Test non-existing keys
	if value := GetOrElse("d", 0, original); value != 0 {
		t.Errorf("Expected 0, got %v", value)
	}
	if value := GetOrElse("x", -1, original); value != -1 {
		t.Errorf("Expected -1, got %v", value)
	}
}

func TestGetOrElseFn(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test existing keys
	if value := GetOrElseFn("a", func() int { return 0 }, original); value != 1 {
		t.Errorf("Expected 1, got %v", value)
	}
	if value := GetOrElseFn("b", func() int { return 0 }, original); value != 2 {
		t.Errorf("Expected 2, got %v", value)
	}
	if value := GetOrElseFn("c", func() int { return 0 }, original); value != 3 {
		t.Errorf("Expected 3, got %v", value)
	}

	// Test non-existing keys
	if value := GetOrElseFn("d", func() int { return 0 }, original); value != 0 {
		t.Errorf("Expected 0, got %v", value)
	}
	if value := GetOrElseFn("x", func() int { return -1 }, original); value != -1 {
		t.Errorf("Expected -1, got %v", value)
	}
}

func TestSize(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	if size := Size(original); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// Test empty map
	emptyMap := map[string]int{}
	if size := Size(emptyMap); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}
}

func TestIsEmpty(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	if IsEmpty(original) {
		t.Error("Expected non-empty map to return false")
	}

	// Test empty map
	emptyMap := map[string]int{}
	if !IsEmpty(emptyMap) {
		t.Error("Expected empty map to return true")
	}
}
