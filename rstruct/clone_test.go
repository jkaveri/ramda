package rstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Additional test types for Clone function
type BasicStruct struct {
	Name string
	Age  int
	City string
}

type PointerStruct struct {
	Name     string
	Age      int
	Address  *BasicStruct
	Tags     []string
	Metadata map[string]interface{}
}

type NestedStruct struct {
	User     User
	Address  Address
	Active   bool
	Settings map[string]string
	History  []string
}

type ComplexStruct struct {
	ID       int
	Name     string
	Profile  *PointerStruct
	Friends  []*BasicStruct
	Data     map[string]*BasicStruct
	Metadata interface{}
}

func TestClone(t *testing.T) {
	// Test basic struct
	basic := BasicStruct{Name: "Alice", Age: 25, City: "New York"}

	// Test struct with pointers
	pointer := PointerStruct{
		Name: "Bob",
		Age:  30,
		Address: &BasicStruct{
			Name: "123 Main St",
			Age:  0,
			City: "Los Angeles",
		},
		Tags: []string{"developer", "golang"},
		Metadata: map[string]interface{}{
			"department": "Engineering",
			"level":      3,
		},
	}

	// Test nested struct
	nested := NestedStruct{
		User:     User{Name: "Charlie", Age: 35, Email: "charlie@example.com"},
		Address:  Address{Street: "456 Oak St", City: "Chicago", Country: "USA"},
		Active:   true,
		Settings: map[string]string{"theme": "dark", "lang": "en"},
		History:  []string{"login", "logout", "login"},
	}

	// Test complex struct
	complex := ComplexStruct{
		ID:   1,
		Name: "David",
		Profile: &PointerStruct{
			Name:     "David Profile",
			Age:      40,
			Address:  &BasicStruct{Name: "789 Pine St", Age: 0, City: "Boston"},
			Tags:     []string{"manager", "senior"},
			Metadata: map[string]interface{}{"role": "lead", "experience": 10},
		},
		Friends: []*BasicStruct{
			{Name: "Friend1", Age: 30, City: "NYC"},
			{Name: "Friend2", Age: 35, City: "LA"},
		},
		Data: map[string]*BasicStruct{
			"home": {Name: "Home", Age: 0, City: "Home City"},
			"work": {Name: "Work", Age: 0, City: "Work City"},
		},
		Metadata: "some metadata",
	}

	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "basic struct",
			input:    basic,
			expected: basic,
		},
		{
			name:     "pointer struct",
			input:    pointer,
			expected: pointer,
		},
		{
			name:     "nested struct",
			input:    nested,
			expected: nested,
		},
		{
			name:     "complex struct",
			input:    complex,
			expected: complex,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloned := Clone(tt.input)

			// Test that the clone is not the same instance
			clonedAddr := &cloned
			inputAddr := &tt.input
			assert.False(t, clonedAddr == inputAddr, "Expected cloned struct to be a different instance")

			// Test that the values are equal
			assert.Equal(t, tt.expected, cloned, "Expected cloned struct to have same values")
		})
	}
}

func TestClonePointerInputs(t *testing.T) {
	// Test that pointer inputs work correctly
	basic := BasicStruct{Name: "Alice", Age: 25, City: "New York"}
	basicPtr := &basic

	pointer := PointerStruct{
		Name: "Bob",
		Age:  30,
		Address: &BasicStruct{
			Name: "123 Main St",
			Age:  0,
			City: "Los Angeles",
		},
		Tags: []string{"developer", "golang"},
		Metadata: map[string]interface{}{
			"department": "Engineering",
			"level":      3,
		},
	}
	pointerPtr := &pointer

	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "basic struct pointer",
			input:    basicPtr,
			expected: basicPtr,
		},
		{
			name:     "pointer struct pointer",
			input:    pointerPtr,
			expected: pointerPtr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloned := Clone(tt.input)

			// Test that the clone is not the same instance
			clonedAddr := &cloned
			inputAddr := &tt.input
			assert.False(t, clonedAddr == inputAddr, "Expected cloned struct to be a different instance")

			// Test that the values are equal (both should be pointers)
			assert.Equal(t, tt.expected, cloned, "Expected cloned struct to have same values")
		})
	}
}

func TestCloneIndependence(t *testing.T) {
	// Test that modifying the original doesn't affect the clone
	original := PointerStruct{
		Name: "Alice",
		Age:  25,
		Address: &BasicStruct{
			Name: "123 Main St",
			Age:  0,
			City: "New York",
		},
		Tags: []string{"developer", "golang"},
		Metadata: map[string]interface{}{
			"department": "Engineering",
			"level":      3,
		},
	}

	cloned := Clone(original)

	// Verify that pointer fields point to different memory addresses
	assert.False(t, original.Address == cloned.Address, "Address pointers should be different")
	assert.False(t, &original.Tags == &cloned.Tags, "Tags slice should be different")
	assert.False(t, &original.Metadata == &cloned.Metadata, "Metadata map should be different")

	// Verify that the values are the same initially
	assert.Equal(t, original.Address.Name, cloned.Address.Name, "Address values should be the same")
	assert.Equal(t, original.Address.City, cloned.Address.City, "Address values should be the same")
	assert.Equal(t, original.Tags[0], cloned.Tags[0], "Tags values should be the same")
	assert.Equal(t, original.Metadata["level"], cloned.Metadata["level"], "Metadata values should be the same")

	// Modify the original
	original.Name = "Bob"
	original.Age = 30
	original.Address.City = "Los Angeles"
	original.Tags[0] = "manager"
	original.Metadata["level"] = 5

	// Verify the clone is independent
	assert.Equal(t, "Alice", cloned.Name, "Clone should be independent of original")
	assert.Equal(t, 25, cloned.Age, "Clone should be independent of original")
	assert.Equal(t, "New York", cloned.Address.City, "Clone should be independent of original")
	assert.Equal(t, "developer", cloned.Tags[0], "Clone should be independent of original")
	assert.Equal(t, 3, cloned.Metadata["level"], "Clone should be independent of original")
}

func TestCloneEdgeCases(t *testing.T) {
	// Test empty struct
	empty := BasicStruct{}
	clonedEmpty := Clone(empty)
	assert.Equal(t, empty, clonedEmpty, "Empty struct should be cloned correctly")

	// Test struct with nil pointer
	nilPointer := PointerStruct{
		Name:     "Test",
		Age:      25,
		Address:  nil,
		Tags:     []string{"test"},
		Metadata: map[string]interface{}{"test": "value"},
	}
	clonedNilPointer := Clone(nilPointer)
	assert.Equal(t, nilPointer, clonedNilPointer, "Struct with nil pointer should be cloned correctly")

	// Test struct with empty slice
	emptySlice := PointerStruct{
		Name:     "Test",
		Age:      25,
		Address:  &BasicStruct{Name: "Test", Age: 0, City: "Test"},
		Tags:     []string{},
		Metadata: map[string]interface{}{},
	}
	clonedEmptySlice := Clone(emptySlice)
	assert.Equal(t, emptySlice, clonedEmptySlice, "Struct with empty slice should be cloned correctly")

	// Test struct with empty map
	emptyMap := PointerStruct{
		Name:     "Test",
		Age:      25,
		Address:  &BasicStruct{Name: "Test", Age: 0, City: "Test"},
		Tags:     []string{"test"},
		Metadata: map[string]interface{}{},
	}
	clonedEmptyMap := Clone(emptyMap)
	assert.Equal(t, emptyMap, clonedEmptyMap, "Struct with empty map should be cloned correctly")
}

func TestCloneDeepCopyIndependence(t *testing.T) {
	// Test complex nested structure to verify deep copy independence
	original := ComplexStruct{
		ID:   1,
		Name: "David",
		Profile: &PointerStruct{
			Name: "David Profile",
			Age:  40,
			Address: &BasicStruct{
				Name: "789 Pine St",
				Age:  0,
				City: "Boston",
			},
			Tags: []string{"manager", "senior"},
			Metadata: map[string]interface{}{
				"role":       "lead",
				"experience": 10,
				"department": "Engineering",
			},
		},
		Friends: []*BasicStruct{
			{Name: "Friend1", Age: 30, City: "NYC"},
			{Name: "Friend2", Age: 35, City: "LA"},
		},
		Data: map[string]*BasicStruct{
			"home": {Name: "Home", Age: 0, City: "Home City"},
			"work": {Name: "Work", Age: 0, City: "Work City"},
		},
		Metadata: "some metadata",
	}

	cloned := Clone(original)

	// Verify that all pointer fields point to different memory addresses
	assert.False(t, original.Profile == cloned.Profile, "Profile pointers should be different")
	assert.False(t, &original.Friends == &cloned.Friends, "Friends slice should be different")
	assert.False(t, &original.Data == &cloned.Data, "Data map should be different")

	// Verify nested pointer independence
	assert.False(t, original.Profile.Address == cloned.Profile.Address, "Profile.Address pointers should be different")
	assert.False(t, &original.Profile.Tags == &cloned.Profile.Tags, "Profile.Tags slice should be different")
	assert.False(t, &original.Profile.Metadata == &cloned.Profile.Metadata, "Profile.Metadata map should be different")

	// Verify slice element independence
	assert.False(t, original.Friends[0] == cloned.Friends[0], "Friends[0] pointers should be different")
	assert.False(t, original.Friends[1] == cloned.Friends[1], "Friends[1] pointers should be different")

	// Verify map value independence
	assert.False(t, original.Data["home"] == cloned.Data["home"], "Data['home'] pointers should be different")
	assert.False(t, original.Data["work"] == cloned.Data["work"], "Data['work'] pointers should be different")

	// Verify that values are the same initially
	assert.Equal(t, original.Profile.Name, cloned.Profile.Name, "Profile values should be the same")
	assert.Equal(t, original.Profile.Address.City, cloned.Profile.Address.City, "Profile.Address values should be the same")
	assert.Equal(t, original.Friends[0].Name, cloned.Friends[0].Name, "Friends[0] values should be the same")
	assert.Equal(t, original.Data["home"].City, cloned.Data["home"].City, "Data['home'] values should be the same")

	// Modify the original deeply
	original.Profile.Name = "Modified Profile"
	original.Profile.Address.City = "Modified City"
	original.Friends[0].Name = "Modified Friend"
	original.Data["home"].City = "Modified Home City"
	original.Profile.Tags[0] = "modified"
	original.Profile.Metadata["role"] = "modified"

	// Verify the clone is completely independent
	assert.Equal(t, "David Profile", cloned.Profile.Name, "Clone should be independent of original")
	assert.Equal(t, "Boston", cloned.Profile.Address.City, "Clone should be independent of original")
	assert.Equal(t, "Friend1", cloned.Friends[0].Name, "Clone should be independent of original")
	assert.Equal(t, "Home City", cloned.Data["home"].City, "Clone should be independent of original")
	assert.Equal(t, "manager", cloned.Profile.Tags[0], "Clone should be independent of original")
	assert.Equal(t, "lead", cloned.Profile.Metadata["role"], "Clone should be independent of original")
}

func TestCloneNonStructTypes(t *testing.T) {
	// Test that non-struct types return zero value
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{"string", "hello", nil},
		{"int", 42, nil},
		{"slice", []int{1, 2, 3}, nil},
		{"map", map[string]int{"a": 1, "b": 2}, nil},
		{"bool", true, nil},
		{"float", 3.14, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloned := Clone(tt.input)
			// For non-struct types, Clone should return nil (zero value for interface{})
			assert.Equal(t, tt.expected, cloned, "Non-struct type should return nil")
		})
	}
}
