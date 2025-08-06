package rstruct

import (
	"fmt"
	"reflect"
	"strings"
)

// Get retrieves a field value from a struct by field name.
// It returns the value and a boolean indicating if the field was found.
// The field name is case-sensitive and supports dot notation for nested fields.
// For embedded structs, it searches recursively.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//	}
//	user := User{Name: "Alice", Age: 25}
//	value, found := Get(user, "Name")
//	// Result: value = "Alice", found = true
//
//	// Nested field access
//	type Address struct {
//		Street string
//		City   string
//	}
//	type Person struct {
//		Name    string
//		Address Address
//	}
//	person := Person{Name: "Alice", Address: Address{Street: "123 Main St", City: "New York"}}
//	value, found = Get(person, "Address.City")
//	// Result: value = "New York", found = true
func Get[T any](obj T, fieldName string) (any, bool) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, false
	}

	return getNestedFieldRecursive(val, strings.Split(fieldName, "."))
}

// GetOrDefault retrieves a field value from a struct by field name.
// If the field is not found, it returns the provided default value.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//	}
//	user := User{Name: "Alice", Age: 25}
//	value := GetOrDefault(user, "age", 0)
//	// Result: value = 25
//	value = GetOrDefault(user, "email", "unknown")
//	// Result: value = "unknown"
func GetOrDefault[T any](obj T, fieldName string, defaultValue any) any {
	value, found := Get(obj, fieldName)
	if !found {
		return defaultValue
	}
	return value
}

// Set sets a field value in a struct by field name.
// The field name is case-sensitive and supports dot notation for nested fields.
// For embedded structs, it searches recursively.
// Returns an error indicating success or failure of the set operation.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//	}
//	user := User{Name: "Alice", Age: 25}
//	err := Set(&user, "Age", 30)
//	// Result: user.Age = 30
//
//	// Nested field access
//	type Address struct {
//		Street string
//		City   string
//	}
//	type Person struct {
//		Name    string
//		Address Address
//	}
//	person := Person{Name: "Alice", Address: Address{Street: "123 Main St", City: "New York"}}
//	err := Set(&person, "Address.City", "Los Angeles")
//	// Result: person.Address.City = "Los Angeles"
func Set[T any](obj T, fieldName string, value any) error {
	val := reflect.ValueOf(obj)
	isPtr := val.Kind() == reflect.Ptr
	if isPtr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("cannot set field on non-struct type: %v", val.Kind())
	}

	return setNestedFieldRecursive(
		val,
		strings.Split(fieldName, "."),
		value,
	)
}

// getNestedFieldRecursive is a helper function to recursively navigate nested struct fields
func getNestedFieldRecursive(val reflect.Value, parts []string) (any, bool) {
	if len(parts) == 0 {
		return nil, false
	}

	currentField := parts[0]
	typ := val.Type()

	// Search for the current field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Name != currentField {
			continue
		}

		// If this is the last part, return the field value
		if len(parts) == 1 {
			return field.Interface(), true
		}

		// If this is a struct, continue navigating
		if field.Kind() == reflect.Struct {
			return getNestedFieldRecursive(field, parts[1:])
		}

		// If it's not a struct but we have more parts, the path is invalid
		return nil, false
	}

	return nil, false
}

// setNestedFieldRecursive is a helper function to recursively set nested struct fields
func setNestedFieldRecursive(val reflect.Value, parts []string, value any) error {
	if len(parts) == 0 {
		return fmt.Errorf("empty field path")
	}

	currentField := parts[0]
	typ := val.Type()

	// Search for the current field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Name != currentField {
			continue
		}

		// If this is the last part, set the field value
		if len(parts) == 1 {
			valueVal := reflect.ValueOf(value)
			if valueVal.Type().AssignableTo(field.Type()) {
				field.Set(valueVal)
				return nil
			}
			return fmt.Errorf("cannot assign value of type %v to field %s of type %v", valueVal.Type(), currentField, field.Type())
		}

		// If this is a struct, continue navigating
		if field.Kind() == reflect.Struct {
			if field.CanSet() {
				// Recursively set the nested field
				return setNestedFieldRecursive(field, parts[1:], value)
			}
			return fmt.Errorf("field %s is not settable", currentField)
		}

		// If it's not a struct but we have more parts, the path is invalid
		return fmt.Errorf("field %s is not a struct, cannot navigate further", currentField)
	}

	// Field not found
	return fmt.Errorf("field %s not found", currentField)
}

// Has checks if a struct has a field with the given name.
// The field name is case-sensitive and supports dot notation for nested fields.
// For embedded structs, it searches recursively.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//	}
//	user := User{Name: "Alice", Age: 25}
//	hasName := Has(user, "Name")
//	// Result: hasName = true
//	hasEmail := Has(user, "email")
//	// Result: hasEmail = false
//
//	// Nested field check
//	type Address struct {
//		Street string
//		City   string
//	}
//	type Person struct {
//		Name    string
//		Address Address
//	}
//	person := Person{Name: "Alice", Address: Address{Street: "123 Main St", City: "New York"}}
//	hasCity := Has(person, "Address.City")
//	// Result: hasCity = true
func Has[T any](obj T, fieldName string) bool {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return false
	}

	return hasNestedFieldRecursive(val, strings.Split(fieldName, "."))
}

// hasNestedFieldRecursive is a helper function to recursively check for nested struct fields
func hasNestedFieldRecursive(val reflect.Value, parts []string) bool {
	if len(parts) == 0 {
		return false
	}

	currentField := parts[0]
	typ := val.Type()

	// Search for the current field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Name != currentField {
			continue
		}

		// If this is the last part, the field exists
		if len(parts) == 1 {
			return true
		}

		// If this is a struct, continue navigating
		if field.Kind() == reflect.Struct {
			return hasNestedFieldRecursive(field, parts[1:])
		}

		// If it's not a struct but we have more parts, the path is invalid
		return false
	}

	return false
}

// Fields returns all field names of a struct as a slice of strings.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
//	fields := Fields(user)
//	// Result: []string{"Name", "Age", "Email"}
func Fields[T any](obj T) []string {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return []string{}
	}

	typ := val.Type()
	fields := make([]string, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		fields = append(fields, fieldType.Name)
	}

	return fields
}

// ToMap converts a struct to a map[string]any where keys are field names.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
//	result := ToMap(user)
//	// Result: map[string]any{
//	//   "Name": "Alice",
//	//   "Age": 25,
//	//   "Email": "alice@example.com"
//	// }
func ToMap[T any](obj T) map[string]any {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return map[string]any{}
	}

	typ := val.Type()
	result := make(map[string]any, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if field.IsValid() && !field.IsZero() && field.CanInterface() {
			result[fieldType.Name] = field.Interface()
		}
	}

	return result
}

// FromMap creates a struct from a map[string]any.
// It returns a new struct instance and a boolean indicating success.
// Only fields that exist in the struct will be set.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	data := map[string]any{
//		"Name": "Alice",
//		"Age": 25,
//		"Email": "alice@example.com",
//	}
//	user, success := FromMap[User](data)
//	// Result: user = User{Name: "Alice", Age: 25, Email: "alice@example.com"}, success = true
func FromMap[T any](data map[string]any) (T, bool) {
	var obj T
	val := reflect.ValueOf(&obj).Elem()

	if val.Kind() != reflect.Struct {
		var zero T
		return zero, false
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if value, exists := data[fieldType.Name]; exists {
			valueVal := reflect.ValueOf(value)
			if valueVal.Type().AssignableTo(field.Type()) {
				field.Set(valueVal)
			}
		}
	}

	return obj, true
}

// Pick creates a new struct with only the specified fields.
// Fields that don't exist in the original struct are ignored.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
//	picked := Pick(user, []string{"Name", "Age"})
//	// Result: User{Name: "Alice", Age: 25, Email: ""}
func Pick[T any](obj T, fields []string) T {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		var zero T
		return zero
	}

	// Create a new struct of the same type
	newVal := reflect.New(val.Type()).Elem()
	typ := val.Type()

	// Create a set of field names to pick (case-sensitive)
	fieldSet := make(map[string]bool)
	for _, field := range fields {
		fieldSet[field] = true
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldSet[fieldType.Name] {
			newVal.Field(i).Set(field)
		}
	}

	return newVal.Interface().(T)
}

// Omit creates a new struct with all fields except the specified ones.
// Fields that don't exist in the original struct are ignored.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
//	omitted := Omit(user, []string{"Email"})
//	// Result: User{Name: "Alice", Age: 25, Email: ""}
func Omit[T any](obj T, fields []string) T {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		var zero T
		return zero
	}

	// Create a new struct of the same type
	newVal := reflect.New(val.Type()).Elem()
	typ := val.Type()

	// Create a set of field names to omit (case-sensitive)
	fieldSet := make(map[string]bool)
	for _, field := range fields {
		fieldSet[field] = true
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if !fieldSet[fieldType.Name] {
			newVal.Field(i).Set(field)
		}
	}

	return newVal.Interface().(T)
}

// Merge combines multiple structs into one.
// Later structs override earlier ones for fields with the same name.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	user1 := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
//	user2 := User{Name: "Bob", Age: 30, Email: ""}
//	merged := Merge(user1, user2)
//	// Result: User{Name: "Bob", Age: 30, Email: "alice@example.com"}
func Merge[T any](structs ...T) T {
	if len(structs) == 0 {
		var zero T
		return zero
	}

	if len(structs) == 1 {
		return structs[0]
	}

	// Start with the first struct
	result := structs[0]

	// Merge each subsequent struct
	for i := 1; i < len(structs); i++ {
		val := reflect.ValueOf(structs[i])
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if val.Kind() != reflect.Struct {
			continue
		}

		resultVal := reflect.ValueOf(&result).Elem()
		typ := val.Type()

		for j := 0; j < val.NumField(); j++ {
			field := val.Field(j)
			fieldType := typ.Field(j)

			// Only set non-zero values
			if field.IsValid() && !field.IsZero() {
				resultVal.FieldByName(fieldType.Name).Set(field)
			}
		}
	}

	return result
}
