package ramda

import (
	"strconv"
	"testing"
)

func TestCast(t *testing.T) {
	// Test successful conversion
	result := Cast(strconv.Atoi, 0, "123")
	if result != 123 {
		t.Errorf("Expected 123, got %d", result)
	}

	// Test failed conversion
	result = Cast(strconv.Atoi, 42, "abc")
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

func TestCastFn(t *testing.T) {
	// Test successful conversion
	result := CastFn(strconv.Atoi, func() int { return 0 }, "123")
	if result != 123 {
		t.Errorf("Expected 123, got %d", result)
	}

	// Test failed conversion
	result = CastFn(strconv.Atoi, func() int { return 42 }, "abc")
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{123, "123"},
		{3.14, "3.14"},
		{true, "true"},
		{"hello", "hello"},
		{nil, "<nil>"},
	}

	for _, test := range tests {
		result := ToString(test.input)
		if result != test.expected {
			t.Errorf("ToString(%v) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"0", 0},
		{"-42", -42},
		{"abc", 0}, // should return 0 for invalid input
		{"", 0},    // should return 0 for empty string
	}

	for _, test := range tests {
		result := ToInt(test.input)
		if result != test.expected {
			t.Errorf("ToInt(%s) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"123", 123},
		{"0", 0},
		{"-42", -42},
		{"abc", 0},
		{"", 0},
	}

	for _, test := range tests {
		result := ToInt64(test.input)
		if result != test.expected {
			t.Errorf("ToInt64(%s) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"123.45", 123.45},
		{"0", 0.0},
		{"-42.5", -42.5},
		{"abc", 0.0},
		{"", 0.0},
	}

	for _, test := range tests {
		result := ToFloat64(test.input)
		if result != test.expected {
			t.Errorf("ToFloat64(%s) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"TRUE", true},
		{"FALSE", false},
		{"1", true},
		{"0", false},
		{"abc", false},
		{"", false},
	}

	for _, test := range tests {
		result := ToBool(test.input)
		if result != test.expected {
			t.Errorf("ToBool(%s) = %t, expected %t", test.input, result, test.expected)
		}
	}
}

func TestFromInt(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{123, "123"},
		{0, "0"},
		{-42, "-42"},
	}

	for _, test := range tests {
		result := FromInt(test.input)
		if result != test.expected {
			t.Errorf("FromInt(%d) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestFromInt64(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{123, "123"},
		{0, "0"},
		{-42, "-42"},
	}

	for _, test := range tests {
		result := FromInt64(test.input)
		if result != test.expected {
			t.Errorf("FromInt64(%d) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestFromFloat64(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{123.45, "123.45"},
		{0.0, "0"},
		{-42.5, "-42.5"},
	}

	for _, test := range tests {
		result := FromFloat64(test.input)
		if result != test.expected {
			t.Errorf("FromFloat64(%f) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestFromBool(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, test := range tests {
		result := FromBool(test.input)
		if result != test.expected {
			t.Errorf("FromBool(%t) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestAs(t *testing.T) {
	// Test successful type assertion
	var val interface{} = "hello"
	result := As[string](val)
	if result != "hello" {
		t.Errorf("As[string](%v) = %s, expected hello", val, result)
	}

	// Test failed type assertion
	var val2 interface{} = 123
	result2 := As[string](val2)
	if result2 != "" {
		t.Errorf("As[string](%v) = %s, expected empty string", val2, result2)
	}
}

func TestAsWithDefault(t *testing.T) {
	// Test successful type assertion
	var val interface{} = "hello"
	result := AsWithDefault(val, "default")
	if result != "hello" {
		t.Errorf("AsWithDefault(%v, default) = %s, expected hello", val, result)
	}

	// Test failed type assertion
	var val2 interface{} = 123
	result2 := AsWithDefault(val2, "default")
	if result2 != "default" {
		t.Errorf("AsWithDefault(%v, default) = %s, expected default", val2, result2)
	}
}

func TestAsWithDefaultFn(t *testing.T) {
	// Test successful type assertion
	var val interface{} = "hello"
	result := AsWithDefaultFn(func() string { return "default" }, val)
	if result != "hello" {
		t.Errorf("AsWithDefaultFn(defaultFn, %v) = %s, expected hello", val, result)
	}

	// Test failed type assertion
	var val2 interface{} = 123
	result2 := AsWithDefaultFn(func() string { return "default" }, val2)
	if result2 != "default" {
		t.Errorf("AsWithDefaultFn(defaultFn, %v) = %s, expected default", val2, result2)
	}
}

func TestTransform(t *testing.T) {
	// Test simple transformation
	result := Transform(func(x int) int { return x * 2 }, 5)
	if result != 10 {
		t.Errorf("Transform(x*2, 5) = %d, expected 10", result)
	}

	// Test type transformation
	result2 := Transform(func(s string) int { return len(s) }, "hello")
	if result2 != 5 {
		t.Errorf("Transform(len, hello) = %d, expected 5", result2)
	}
}

func TestTransformWithError(t *testing.T) {
	// Test successful transformation
	result := TransformWithError(strconv.Atoi, 0, "123")
	if result != 123 {
		t.Errorf("TransformWithError(Atoi, 0, 123) = %d, expected 123", result)
	}

	// Test failed transformation
	result2 := TransformWithError(strconv.Atoi, 42, "abc")
	if result2 != 42 {
		t.Errorf("TransformWithError(Atoi, 42, abc) = %d, expected 42", result2)
	}
}

func TestTransformWithErrorFn(t *testing.T) {
	// Test successful transformation
	result := TransformWithErrorFn(strconv.Atoi, func() int { return 0 }, "123")
	if result != 123 {
		t.Errorf("TransformWithErrorFn(Atoi, defaultFn, 123) = %d, expected 123", result)
	}

	// Test failed transformation
	result2 := TransformWithErrorFn(strconv.Atoi, func() int { return 42 }, "abc")
	if result2 != 42 {
		t.Errorf("TransformWithErrorFn(Atoi, defaultFn, abc) = %d, expected 42", result2)
	}
}

func TestFromPtr(t *testing.T) {
	// Test with non-nil pointer
	val := 42
	ptr := &val
	result := FromPtr(ptr)
	if result != 42 {
		t.Errorf("FromPtr(%v) = %d, expected 42", ptr, result)
	}

	// Test with nil pointer
	var nilPtr *int
	result2 := FromPtr(nilPtr)
	if result2 != 0 {
		t.Errorf("FromPtr(nil) = %d, expected 0", result2)
	}

	// Test with string pointer
	str := "hello"
	strPtr := &str
	result3 := FromPtr(strPtr)
	if result3 != "hello" {
		t.Errorf("FromPtr(%v) = %s, expected hello", strPtr, result3)
	}

	// Test with nil string pointer
	var nilStrPtr *string
	result4 := FromPtr(nilStrPtr)
	if result4 != "" {
		t.Errorf("FromPtr(nil string) = %s, expected empty string", result4)
	}
}

func TestToPtr(t *testing.T) {
	// Test with non-zero value
	result := ToPtr(42)
	if result == nil {
		t.Error("ToPtr(42) returned nil, expected pointer to 42")
	}
	if *result != 42 {
		t.Errorf("ToPtr(42) = %d, expected 42", *result)
	}

	// Test with zero value
	result2 := ToPtr(0)
	if result2 == nil {
		t.Error("ToPtr(0) returned nil, expected pointer to 0")
	}
	if *result2 != 0 {
		t.Errorf("ToPtr(0) = %d, expected 0", *result2)
	}

	// Test with string
	result3 := ToPtr("hello")
	if result3 == nil {
		t.Error("ToPtr(hello) returned nil, expected pointer to hello")
	}
	if *result3 != "hello" {
		t.Errorf("ToPtr(hello) = %s, expected hello", *result3)
	}

	// Test with empty string
	result4 := ToPtr("")
	if result4 == nil {
		t.Error("ToPtr(empty) returned nil, expected pointer to empty string")
	}
	if *result4 != "" {
		t.Errorf("ToPtr(empty) = %s, expected empty string", *result4)
	}

	// Test with boolean
	result5 := ToPtr(true)
	if result5 == nil {
		t.Error("ToPtr(true) returned nil, expected pointer to true")
	}
	if *result5 != true {
		t.Errorf("ToPtr(true) = %t, expected true", *result5)
	}

	// Test with false
	result6 := ToPtr(false)
	if result6 == nil {
		t.Error("ToPtr(false) returned nil, expected pointer to false")
	}
	if *result6 != false {
		t.Errorf("ToPtr(false) = %t, expected false", *result6)
	}
}

func TestDefault(t *testing.T) {
	// Test with non-zero value
	result := Default(42, 100)
	if result != 42 {
		t.Errorf("Default(42, 100) = %d, expected 42", result)
	}

	// Test with zero value
	result2 := Default(0, 100)
	if result2 != 100 {
		t.Errorf("Default(0, 100) = %d, expected 100", result2)
	}

	// Test with string
	result3 := Default("hello", "default")
	if result3 != "hello" {
		t.Errorf("Default(hello, default) = %s, expected hello", result3)
	}

	// Test with empty string
	result4 := Default("", "default")
	if result4 != "default" {
		t.Errorf("Default(empty, default) = %s, expected default", result4)
	}
}

func TestDefaultFn(t *testing.T) {
	// Test with non-zero value
	result := DefaultFn(func() int { return 100 }, 42)
	if result != 42 {
		t.Errorf("DefaultFn(defaultFn, 42) = %d, expected 42", result)
	}

	// Test with zero value
	result2 := DefaultFn(func() int { return 100 }, 0)
	if result2 != 100 {
		t.Errorf("DefaultFn(defaultFn, 0) = %d, expected 100", result2)
	}

	// Test with string
	result3 := DefaultFn(func() string { return "default" }, "hello")
	if result3 != "hello" {
		t.Errorf("DefaultFn(defaultFn, hello) = %s, expected hello", result3)
	}

	// Test with empty string
	result4 := DefaultFn(func() string { return "default" }, "")
	if result4 != "default" {
		t.Errorf("DefaultFn(defaultFn, empty) = %s, expected default", result4)
	}
}

func TestNilIfEmpty(t *testing.T) {
	// Test with non-zero value
	result := NilIfEmpty(42)
	if result == nil {
		t.Error("NilIfEmpty(42) returned nil, expected pointer to 42")
	}
	if *result != 42 {
		t.Errorf("NilIfEmpty(42) = %d, expected 42", *result)
	}

	// Test with zero value
	result2 := NilIfEmpty(0)
	if result2 != nil {
		t.Errorf("NilIfEmpty(0) = %v, expected nil", result2)
	}

	// Test with string
	result3 := NilIfEmpty("hello")
	if result3 == nil {
		t.Error("NilIfEmpty(hello) returned nil, expected pointer to hello")
	}
	if *result3 != "hello" {
		t.Errorf("NilIfEmpty(hello) = %s, expected hello", *result3)
	}

	// Test with empty string
	result4 := NilIfEmpty("")
	if result4 != nil {
		t.Errorf("NilIfEmpty(empty) = %v, expected nil", result4)
	}

	// Test with boolean
	result5 := NilIfEmpty(true)
	if result5 == nil {
		t.Error("NilIfEmpty(true) returned nil, expected pointer to true")
	}
	if *result5 != true {
		t.Errorf("NilIfEmpty(true) = %t, expected true", *result5)
	}

	// Test with false
	result6 := NilIfEmpty(false)
	if result6 != nil {
		t.Errorf("NilIfEmpty(false) = %v, expected nil", result6)
	}
}
