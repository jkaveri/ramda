package main

import (
	"fmt"
	"strconv"

	"github.com/jkaveri/ramda"
)

func main() {
	// Example 1: Using Cast with the new parameter order
	// The function parameter comes first, making it easier to create specialized versions
	toIntWithDefault := func(s string) int {
		return ramda.Cast(strconv.Atoi, 0, s)
	}

	result1 := toIntWithDefault("123")
	result2 := toIntWithDefault("abc")

	fmt.Printf("Cast with new parameter order: %d, %d\n", result1, result2) // 123, 0

	// Example 2: Using Transform with the new parameter order
	double := func(x int) int { return x * 2 }
	toString := func(x int) string { return fmt.Sprintf("Result: %d", x) }

	// The new parameter order makes it easier to create specialized transformers
	curriedDouble := func(val int) int {
		return ramda.Transform(double, val)
	}
	curriedToString := func(val int) string {
		return ramda.Transform(toString, val)
	}

	// Compose the functions - we need to handle the type mismatch
	doubleThenString := func(val int) string {
		doubled := curriedDouble(val)
		return curriedToString(doubled)
	}

	result3 := doubleThenString(5)
	fmt.Printf("Composed transformation: %s\n", result3) // "Result: 10"

	// Example 3: Using TransformWithError with the new parameter order
	parseInt := func(s string) int {
		return ramda.TransformWithError(strconv.Atoi, 0, s)
	}

	result4 := parseInt("456")
	result5 := parseInt("xyz")

	fmt.Printf("TransformWithError with new parameter order: %d, %d\n", result4, result5) // 456, 0

	// Example 4: Building a pipeline with the new parameter order
	toInt := func(s string) int {
		return ramda.Cast(strconv.Atoi, 0, s)
	}

	// Create a pipeline that processes strings through multiple transformations
	processString := func(s string) string {
		asInt := toInt(s)
		doubled := curriedDouble(asInt)
		return curriedToString(doubled)
	}

	// Process a list of strings
	inputs := []string{"1", "2", "abc", "3"}
	for _, input := range inputs {
		result := processString(input)
		fmt.Printf("Pipeline result for '%s': %s\n", input, result)
	}

	// Example 5: Demonstrating the benefit of function-first parameter order
	// We can easily create specialized casting functions
	toIntSafe := func(s string) int {
		return ramda.Cast(strconv.Atoi, 0, s)
	}

	// Create a custom float parser that matches our function signature
	parseFloat := func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}

	toFloatSafe := func(s string) float64 {
		return ramda.Cast(parseFloat, 0.0, s)
	}

	// These can be easily composed
	processNumbers := func(s string) string {
		asInt := toIntSafe(s)
		asFloat := toFloatSafe(fmt.Sprintf("%d", asInt))
		return fmt.Sprintf("Processed: %.2f", asFloat)
	}

	fmt.Printf("Complex pipeline: %s\n", processNumbers("123")) // "Processed: 123.00"

	// Example 6: Pointer casting functions
	fmt.Println("\nPointer casting examples:")

	// FromPtr examples
	var nilIntPtr *int
	var nilStrPtr *string
	val := 42
	str := "hello"
	valPtr := &val
	strPtr := &str

	// Safe pointer dereferencing
	fmt.Printf("FromPtr(nil int): %d\n", ramda.FromPtr(nilIntPtr))    // 0
	fmt.Printf("FromPtr(nil string): %s\n", ramda.FromPtr(nilStrPtr)) // ""
	fmt.Printf("FromPtr(&42): %d\n", ramda.FromPtr(valPtr))           // 42
	fmt.Printf("FromPtr(&hello): %s\n", ramda.FromPtr(strPtr))        // hello

	// FromPtr with Default examples
	fmt.Printf("Default(FromPtr(nil), 100): %d\n", ramda.Default(ramda.FromPtr(nilIntPtr), 100))           // 100
	fmt.Printf("Default(FromPtr(nil), default): %s\n", ramda.Default(ramda.FromPtr(nilStrPtr), "default")) // default
	fmt.Printf("Default(FromPtr(&42), 100): %d\n", ramda.Default(ramda.FromPtr(valPtr), 100))              // 42

	// FromPtr with DefaultFn examples
	fmt.Printf("DefaultFn(fn, FromPtr(nil)): %d\n", ramda.DefaultFn(func() int { return 999 }, ramda.FromPtr(nilIntPtr))) // 999
	fmt.Printf("DefaultFn(fn, FromPtr(&42)): %d\n", ramda.DefaultFn(func() int { return 999 }, ramda.FromPtr(valPtr)))    // 42

	// ToPtr examples (always returns pointer)
	fmt.Printf("ToPtr(42): %v\n", ramda.ToPtr(42))         // &42
	fmt.Printf("ToPtr(0): %v\n", ramda.ToPtr(0))           // &0
	fmt.Printf("ToPtr(hello): %v\n", ramda.ToPtr("hello")) // &hello
	fmt.Printf("ToPtr(empty): %v\n", ramda.ToPtr(""))      // &""
	fmt.Printf("ToPtr(true): %v\n", ramda.ToPtr(true))     // &true
	fmt.Printf("ToPtr(false): %v\n", ramda.ToPtr(false))   // &false

	// NilIfEmpty examples (returns nil for zero values)
	fmt.Printf("NilIfEmpty(42): %v\n", ramda.NilIfEmpty(42))         // &42
	fmt.Printf("NilIfEmpty(0): %v\n", ramda.NilIfEmpty(0))           // nil
	fmt.Printf("NilIfEmpty(hello): %v\n", ramda.NilIfEmpty("hello")) // &hello
	fmt.Printf("NilIfEmpty(empty): %v\n", ramda.NilIfEmpty(""))      // nil
	fmt.Printf("NilIfEmpty(true): %v\n", ramda.NilIfEmpty(true))     // &true
	fmt.Printf("NilIfEmpty(false): %v\n", ramda.NilIfEmpty(false))   // nil

	// Example 7: Building a pipeline with pointer handling
	processOptionalData := func(data *string) string {
		// Safely extract the string value with default
		str := ramda.Default(ramda.FromPtr(data), "")
		// Convert to int if possible
		asInt := ramda.Cast(strconv.Atoi, 0, str)
		// Double the value
		doubled := ramda.Transform(func(x int) int { return x * 2 }, asInt)
		// Convert back to string
		return ramda.Transform(func(x int) string { return fmt.Sprintf("Processed: %d", x) }, doubled)
	}

	// Test with various pointer scenarios
	testData := []*string{
		ramda.NilIfEmpty("123"),
		ramda.NilIfEmpty("abc"),
		nil,
		ramda.NilIfEmpty("456"),
	}

	fmt.Println("\nPipeline with pointer handling:")
	for i, data := range testData {
		result := processOptionalData(data)
		if data == nil {
			fmt.Printf("  Data[%d] (nil): %s\n", i, result)
		} else {
			fmt.Printf("  Data[%d] (%s): %s\n", i, *data, result)
		}
	}
}
