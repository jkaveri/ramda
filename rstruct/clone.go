package rstruct

import "reflect"

// Clone creates a deep copy of a struct.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//		Email string
//	}
//	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
//	cloned := Clone(user)
//	// Result: cloned = User{Name: "Alice", Age: 25, Email: "alice@example.com"}
func Clone[T any](obj T) T {
	val := reflect.ValueOf(obj)
	isPtr := val.Kind() == reflect.Ptr
	if isPtr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		var zero T
		return zero
	}

	// Create a deep copy
	newVal := deepCopyValue(val)

	if isPtr {
		// Return pointer to the new struct
		ptr := reflect.New(newVal.Type())
		ptr.Elem().Set(newVal)
		return ptr.Interface().(T)
	}

	return newVal.Interface().(T)
}

// deepCopyValue performs a deep copy of a reflect.Value
func deepCopyValue(val reflect.Value) reflect.Value {
	if !val.IsValid() {
		return reflect.Value{}
	}

	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return reflect.New(val.Type()).Elem()
		}
		elem := deepCopyValue(val.Elem())
		ptr := reflect.New(elem.Type())
		ptr.Elem().Set(elem)
		return ptr

	case reflect.Struct:
		// Always recursively copy structs to ensure deep copy
		newStruct := reflect.New(val.Type()).Elem()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			newStruct.Field(i).Set(deepCopyValue(field))
		}
		return newStruct

	case reflect.Slice:
		if val.IsNil() {
			return reflect.New(val.Type()).Elem()
		}
		newSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			newSlice.Index(i).Set(deepCopyValue(val.Index(i)))
		}
		return newSlice

	case reflect.Map:
		if val.IsNil() {
			return reflect.New(val.Type()).Elem()
		}
		newMap := reflect.MakeMap(val.Type())
		for _, key := range val.MapKeys() {
			newMap.SetMapIndex(deepCopyValue(key), deepCopyValue(val.MapIndex(key)))
		}
		return newMap

	case reflect.Array:
		newArray := reflect.New(val.Type()).Elem()
		for i := 0; i < val.Len(); i++ {
			newArray.Index(i).Set(deepCopyValue(val.Index(i)))
		}
		return newArray

	case reflect.Interface:
		if val.IsNil() {
			return reflect.New(val.Type()).Elem()
		}
		elem := deepCopyValue(val.Elem())
		newInterface := reflect.New(val.Type()).Elem()
		newInterface.Set(elem)
		return newInterface

	default:
		// For basic types (int, string, bool, etc.), just copy directly
		newVal := reflect.New(val.Type()).Elem()
		newVal.Set(val)
		return newVal
	}
}
