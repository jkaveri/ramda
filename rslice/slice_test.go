package rslice

import (
	"testing"
)

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(func(n int) int { return n * 2 }, numbers)
	expected := []int{2, 4, 6, 8, 10}

	if len(doubled) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(doubled))
	}

	for i, v := range expected {
		if doubled[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, doubled[i])
		}
	}

	// Test with different types
	strings := []string{"hello", "world", "go"}
	lengths := Map(func(s string) int { return len(s) }, strings)
	expectedLengths := []int{5, 5, 2}

	for i, v := range expectedLengths {
		if lengths[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, lengths[i])
		}
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	even := Filter(func(n int) bool { return n%2 == 0 }, numbers)
	expected := []int{2, 4, 6}

	if len(even) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(even))
	}

	for i, v := range expected {
		if even[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, even[i])
		}
	}
}

func TestReduce(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(func(acc, n int) int { return acc + n }, 0, numbers)
	expected := 15

	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}

	// Test with different types
	strings := []string{"hello", "world", "go"}
	concatenated := Reduce(func(acc, s string) string { return acc + " " + s }, "", strings)
	expectedStr := " hello world go"

	if concatenated != expectedStr {
		t.Errorf("Expected %s, got %s", expectedStr, concatenated)
	}
}

func TestFind(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	found, exists := Find(func(n int) bool { return n > 3 }, numbers)

	if !exists {
		t.Error("Expected to find element, but didn't")
	}

	if found != 4 {
		t.Errorf("Expected 4, got %d", found)
	}

	// Test not found
	_, exists = Find(func(n int) bool { return n > 10 }, numbers)
	if exists {
		t.Error("Expected not to find element, but did")
	}
}

func TestAny(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	hasEven := Any(func(n int) bool { return n%2 == 0 }, numbers)

	if !hasEven {
		t.Error("Expected to find even number, but didn't")
	}

	// Test no match
	hasNegative := Any(func(n int) bool { return n < 0 }, numbers)
	if hasNegative {
		t.Error("Expected not to find negative number, but did")
	}
}

func TestAll(t *testing.T) {
	evenNumbers := []int{2, 4, 6, 8, 10}
	allEven := All(func(n int) bool { return n%2 == 0 }, evenNumbers)

	if !allEven {
		t.Error("Expected all numbers to be even, but they weren't")
	}

	// Test not all match
	mixedNumbers := []int{2, 4, 5, 6, 8}
	allEvenMixed := All(func(n int) bool { return n%2 == 0 }, mixedNumbers)
	if allEvenMixed {
		t.Error("Expected not all numbers to be even, but they were")
	}
}

func TestTake(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	firstThree := Take(3, numbers)
	expected := []int{1, 2, 3}

	if len(firstThree) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(firstThree))
	}

	for i, v := range expected {
		if firstThree[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, firstThree[i])
		}
	}

	// Test take more than available
	all := Take(10, numbers)
	if len(all) != len(numbers) {
		t.Errorf("Expected length %d, got %d", len(numbers), len(all))
	}

	// Test take zero
	empty := Take(0, numbers)
	if len(empty) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(empty))
	}
}

func TestDrop(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	withoutFirstTwo := Drop(2, numbers)
	expected := []int{3, 4, 5}

	if len(withoutFirstTwo) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(withoutFirstTwo))
	}

	for i, v := range expected {
		if withoutFirstTwo[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, withoutFirstTwo[i])
		}
	}

	// Test drop more than available
	empty := Drop(10, numbers)
	if len(empty) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(empty))
	}

	// Test drop zero
	all := Drop(0, numbers)
	if len(all) != len(numbers) {
		t.Errorf("Expected length %d, got %d", len(numbers), len(all))
	}
}

func TestUnique(t *testing.T) {
	numbers := []int{1, 2, 2, 3, 3, 3, 4}
	unique := Unique(numbers)
	expected := []int{1, 2, 3, 4}

	if len(unique) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(unique))
	}

	for i, v := range expected {
		if unique[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, unique[i])
		}
	}
}

func TestFlatten(t *testing.T) {
	matrix := [][]int{{1, 2}, {3, 4}, {5}}
	flattened := Flatten(matrix)
	expected := []int{1, 2, 3, 4, 5}

	if len(flattened) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(flattened))
	}

	for i, v := range expected {
		if flattened[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, flattened[i])
		}
	}
}

func TestZip(t *testing.T) {
	numbers := []int{1, 2, 3}
	letters := []string{"a", "b", "c"}
	zipped := Zip(numbers, letters)

	if len(zipped) != 3 {
		t.Errorf("Expected length 3, got %d", len(zipped))
	}

	expected := []struct {
		First  int
		Second string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	for i, v := range expected {
		if zipped[i].First != v.First {
			t.Errorf("Expected First %d at index %d, got %d", v.First, i, zipped[i].First)
		}
		if zipped[i].Second != v.Second {
			t.Errorf("Expected Second %s at index %d, got %s", v.Second, i, zipped[i].Second)
		}
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	reversed := Reverse(numbers)
	expected := []int{5, 4, 3, 2, 1}

	if len(reversed) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(reversed))
	}

	for i, v := range expected {
		if reversed[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, reversed[i])
		}
	}
}

func TestGroupBy(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	grouped := GroupBy(func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}, numbers)

	even, exists := grouped["even"]
	if !exists {
		t.Error("Expected even group to exist")
	}
	if len(even) != 3 {
		t.Errorf("Expected 3 even numbers, got %d", len(even))
	}

	odd, exists := grouped["odd"]
	if !exists {
		t.Error("Expected odd group to exist")
	}
	if len(odd) != 3 {
		t.Errorf("Expected 3 odd numbers, got %d", len(odd))
	}
}

func TestSortBy(t *testing.T) {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sorted := SortBy(func(a, b int) bool { return a < b }, numbers)
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}

	if len(sorted) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(sorted))
	}

	for i, v := range expected {
		if sorted[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, sorted[i])
		}
	}
}

func TestIndexBy(t *testing.T) {
	type User struct {
		ID   string
		Name string
	}

	users := []User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Charlie"},
	}

	indexed := IndexBy(func(u User) string { return u.ID }, users)

	// Test that all users are indexed
	if len(indexed) != 3 {
		t.Errorf("Expected 3 users in map, got %d", len(indexed))
	}

	// Test specific entries
	alice, exists := indexed["1"]
	if !exists {
		t.Error("Expected Alice to be indexed")
	}
	if alice.Name != "Alice" {
		t.Errorf("Expected Alice, got %s", alice.Name)
	}

	bob, exists := indexed["2"]
	if !exists {
		t.Error("Expected Bob to be indexed")
	}
	if bob.Name != "Bob" {
		t.Errorf("Expected Bob, got %s", bob.Name)
	}

	charlie, exists := indexed["3"]
	if !exists {
		t.Error("Expected Charlie to be indexed")
	}
	if charlie.Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", charlie.Name)
	}

	// Test overwrite behavior
	duplicateUsers := []User{
		{ID: "1", Name: "Alice"},
		{ID: "1", Name: "AliceUpdated"},
		{ID: "2", Name: "Bob"},
	}

	indexedWithDuplicates := IndexBy(func(u User) string { return u.ID }, duplicateUsers)

	if len(indexedWithDuplicates) != 2 {
		t.Errorf("Expected 2 unique keys, got %d", len(indexedWithDuplicates))
	}

	// Should have the last value for duplicate keys
	updatedAlice, exists := indexedWithDuplicates["1"]
	if !exists {
		t.Error("Expected Alice to be indexed")
	}
	if updatedAlice.Name != "AliceUpdated" {
		t.Errorf("Expected AliceUpdated, got %s", updatedAlice.Name)
	}
}

func TestToSet(t *testing.T) {
	// Test with integers
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	numberSet := ToSet(numbers)

	// Test that all unique values are in the set
	expectedNumbers := []int{1, 2, 3, 4}
	if len(numberSet) != len(expectedNumbers) {
		t.Errorf("Expected %d unique numbers, got %d", len(expectedNumbers), len(numberSet))
	}

	for _, num := range expectedNumbers {
		if _, exists := numberSet[num]; !exists {
			t.Errorf("Expected %d to be in set, but it wasn't", num)
		}
	}

	// Test that duplicates are not in the set
	if len(numberSet) != 4 {
		t.Errorf("Expected 4 unique numbers, got %d", len(numberSet))
	}

	// Test with strings
	words := []string{"hello", "world", "hello", "go", "world", "programming"}
	wordSet := ToSet(words)

	expectedWords := []string{"hello", "world", "go", "programming"}
	if len(wordSet) != len(expectedWords) {
		t.Errorf("Expected %d unique words, got %d", len(expectedWords), len(wordSet))
	}

	for _, word := range expectedWords {
		if _, exists := wordSet[word]; !exists {
			t.Errorf("Expected %s to be in set, but it wasn't", word)
		}
	}

	// Test with empty slice
	emptySlice := []int{}
	emptySet := ToSet(emptySlice)
	if len(emptySet) != 0 {
		t.Errorf("Expected empty set, got %d elements", len(emptySet))
	}

	// Test membership checking
	if _, exists := numberSet[5]; exists {
		t.Error("Expected 5 to not be in set, but it was")
	}

	if _, exists := numberSet[1]; !exists {
		t.Error("Expected 1 to be in set, but it wasn't")
	}
}
