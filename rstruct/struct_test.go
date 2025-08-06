package rstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	User
	Address Address
	Active  bool
}

func TestGet(t *testing.T) {
	user := User{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}
	userPtr := &user

	tests := []struct {
		name     string
		input    interface{}
		field    string
		expected interface{}
		found    bool
	}{
		{
			name:     "existing field - name",
			input:    user,
			field:    "Name",
			expected: "Alice",
			found:    true,
		},
		{
			name:     "existing field - age",
			input:    user,
			field:    "Age",
			expected: 25,
			found:    true,
		},
		{
			name:     "case sensitive - Email",
			input:    user,
			field:    "Email",
			expected: "alice@example.com",
			found:    true,
		},
		{
			name:     "non-existing field",
			input:    user,
			field:    "phone",
			expected: nil,
			found:    false,
		},
		{
			name:     "pointer input",
			input:    userPtr,
			field:    "Name",
			expected: "Alice",
			found:    true,
		},
		{
			name:     "nested field",
			input:    user,
			field:    "Address.City",
			expected: "New York",
			found:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := Get(tt.input, tt.field)
			assert.Equal(t, tt.found, found, "Expected found to be %v, got %v", tt.found, found)
			if tt.found {
				assert.Equal(t, tt.expected, value, "Expected %v, got %v", tt.expected, value)
			}
		})
	}
}

func TestGetOrDefault(t *testing.T) {
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}

	tests := []struct {
		name         string
		input        interface{}
		field        string
		defaultValue interface{}
		expected     interface{}
	}{
		{
			name:         "existing field",
			input:        user,
			field:        "Age",
			defaultValue: 0,
			expected:     25,
		},
		{
			name:         "non-existing field",
			input:        user,
			field:        "phone",
			defaultValue: "unknown",
			expected:     "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := GetOrDefault(
				tt.input,
				tt.field,
				tt.defaultValue,
			)
			assert.Equal(
				t,
				tt.expected,
				value,
				"Expected %v, got %v",
				tt.expected,
				value,
			)
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name          string
		input         interface{}
		field         string
		value         interface{}
		expectedAge   int
		expectedName  string
		expectedEmail string
		expectError   bool
	}{
		{
			name:          "setting existing field",
			input:         &User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			field:         "Age",
			value:         30,
			expectedAge:   30,
			expectedName:  "Alice",
			expectedEmail: "alice@example.com",
			expectError:   false,
		},
		{
			name:          "case sensitive",
			input:         &User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			field:         "Email",
			value:         "new@example.com",
			expectedAge:   25,
			expectedName:  "Alice",
			expectedEmail: "new@example.com",
			expectError:   false,
		},
		{
			name:          "non-existing field",
			input:         &User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			field:         "phone",
			value:         "123-456-7890",
			expectedAge:   25,
			expectedName:  "Alice",
			expectedEmail: "alice@example.com",
			expectError:   true,
		},
		{
			name:          "type mismatch",
			input:         &User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			field:         "Age",
			value:         "thirty",
			expectedAge:   25,
			expectedName:  "Alice",
			expectedEmail: "alice@example.com",
			expectError:   true,
		},
		{
			name:          "pointer input",
			input:         &User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			field:         "Name",
			value:         "Bob",
			expectedAge:   25,
			expectedName:  "Bob",
			expectedEmail: "alice@example.com",
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Set(tt.input, tt.field, tt.value)

			if tt.expectError {
				assert.Error(t, err, "Expected error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got: %v", err)

				// Check the modified struct
				if userResult, ok := tt.input.(*User); ok {
					assert.Equal(t, tt.expectedAge, userResult.Age, "Expected Age %d, got %d", tt.expectedAge, userResult.Age)
					assert.Equal(t, tt.expectedName, userResult.Name, "Expected Name %s, got %s", tt.expectedName, userResult.Name)
					assert.Equal(t, tt.expectedEmail, userResult.Email, "Expected Email %s, got %s", tt.expectedEmail, userResult.Email)
				}
			}
		})
	}
}

func TestHas(t *testing.T) {
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	userPtr := &user

	tests := []struct {
		name     string
		input    interface{}
		field    string
		expected bool
	}{
		{
			name:     "existing field - name",
			input:    user,
			field:    "Name",
			expected: true,
		},
		{
			name:     "existing field - age",
			input:    user,
			field:    "Age",
			expected: true,
		},
		{
			name:     "existing field - email",
			input:    user,
			field:    "Email",
			expected: true,
		},
		{
			name:     "case sensitive - Name",
			input:    user,
			field:    "Name",
			expected: true,
		},
		{
			name:     "non-existing field",
			input:    user,
			field:    "phone",
			expected: false,
		},
		{
			name:     "pointer input",
			input:    userPtr,
			field:    "Name",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Has(tt.input, tt.field)
			assert.Equal(t, tt.expected, result, "Expected %v, got %v", tt.expected, result)
		})
	}
}

func TestFields(t *testing.T) {
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	userPtr := &user

	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{
			name:     "struct input",
			input:    user,
			expected: []string{"Name", "Age", "Email", "Address"},
		},
		{
			name:     "pointer input",
			input:    userPtr,
			expected: []string{"Name", "Age", "Email", "Address"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := Fields(tt.input)
			assert.Equal(t, len(tt.expected), len(fields), "Expected %d fields, got %d", len(tt.expected), len(fields))

			for i, field := range tt.expected {
				assert.Equal(t, field, fields[i], "Expected field %s, got %s", field, fields[i])
			}
		})
	}
}

func TestToMap(t *testing.T) {
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	userPtr := &user

	expectedMap := map[string]interface{}{
		"Name":  "Alice",
		"Age":   25,
		"Email": "alice@example.com",
	}

	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "struct input",
			input:    user,
			expected: expectedMap,
		},
		{
			name:     "pointer input",
			input:    userPtr,
			expected: expectedMap,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMap(tt.input)
			assert.Equal(t, len(tt.expected), len(result), "Expected %d entries, got %d", len(tt.expected), len(result))

			for key, value := range tt.expected {
				assert.Equal(t, value, result[key], "Expected %v for key %s, got %v", value, key, result[key])
			}
		})
	}
}

func TestFromMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected User
		success  bool
	}{
		{
			name: "valid data",
			input: map[string]interface{}{
				"Name":  "Alice",
				"Age":   25,
				"Email": "alice@example.com",
			},
			expected: User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			success:  true,
		},
		{
			name: "extra fields",
			input: map[string]interface{}{
				"Name":  "Alice",
				"Age":   25,
				"Email": "alice@example.com",
				"Extra": "value",
			},
			expected: User{Name: "Alice", Age: 25, Email: "alice@example.com"},
			success:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, success := FromMap[User](tt.input)
			assert.Equal(t, tt.success, success, "Expected success to be %v, got %v", tt.success, success)

			if tt.success {
				assert.Equal(t, tt.expected.Name, user.Name, "Expected Name '%s', got %s", tt.expected.Name, user.Name)
				assert.Equal(t, tt.expected.Age, user.Age, "Expected Age %d, got %d", tt.expected.Age, user.Age)
				assert.Equal(t, tt.expected.Email, user.Email, "Expected Email '%s', got %s", tt.expected.Email, user.Email)
			}
		})
	}
}

func TestPick(t *testing.T) {
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}

	tests := []struct {
		name     string
		input    User
		fields   []string
		expected User
	}{
		{
			name:     "picking existing fields",
			input:    user,
			fields:   []string{"Name", "Age"},
			expected: User{Name: "Alice", Age: 25, Email: ""},
		},
		{
			name:     "case sensitive",
			input:    user,
			fields:   []string{"Name", "Email"},
			expected: User{Name: "Alice", Age: 0, Email: "alice@example.com"},
		},
		{
			name:     "with non-existing fields",
			input:    user,
			fields:   []string{"Name", "Phone"},
			expected: User{Name: "Alice", Age: 0, Email: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			picked := Pick(tt.input, tt.fields)
			assert.Equal(t, tt.expected.Name, picked.Name, "Expected Name '%s', got %s", tt.expected.Name, picked.Name)
			assert.Equal(t, tt.expected.Age, picked.Age, "Expected Age %d, got %d", tt.expected.Age, picked.Age)
			assert.Equal(t, tt.expected.Email, picked.Email, "Expected Email '%s', got %s", tt.expected.Email, picked.Email)
		})
	}
}

func TestOmit(t *testing.T) {
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}

	tests := []struct {
		name     string
		input    User
		fields   []string
		expected User
	}{
		{
			name:     "omitting existing fields",
			input:    user,
			fields:   []string{"Email"},
			expected: User{Name: "Alice", Age: 25, Email: ""},
		},
		{
			name:     "case sensitive",
			input:    user,
			fields:   []string{"Email", "Age"},
			expected: User{Name: "Alice", Age: 0, Email: ""},
		},
		{
			name:     "with non-existing fields",
			input:    user,
			fields:   []string{"Phone"},
			expected: User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			omitted := Omit(tt.input, tt.fields)
			assert.Equal(t, tt.expected.Name, omitted.Name, "Expected Name '%s', got %s", tt.expected.Name, omitted.Name)
			assert.Equal(t, tt.expected.Age, omitted.Age, "Expected Age %d, got %d", tt.expected.Age, omitted.Age)
			assert.Equal(t, tt.expected.Email, omitted.Email, "Expected Email '%s', got %s", tt.expected.Email, omitted.Email)
		})
	}
}

func TestMerge(t *testing.T) {
	user1 := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	user2 := User{Name: "Bob", Age: 30, Email: ""}
	user3 := User{Name: "", Age: 0, Email: "charlie@example.com"}

	tests := []struct {
		name     string
		input    []User
		expected User
	}{
		{
			name:     "merging two structs",
			input:    []User{user1, user2},
			expected: User{Name: "Bob", Age: 30, Email: "alice@example.com"},
		},
		{
			name:     "merging three structs",
			input:    []User{user1, user2, user3},
			expected: User{Name: "Bob", Age: 30, Email: "charlie@example.com"},
		},
		{
			name:     "single struct",
			input:    []User{user1},
			expected: User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		},
		{
			name:     "empty slice",
			input:    []User{},
			expected: User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := Merge(tt.input...)
			assert.Equal(t, tt.expected.Name, merged.Name, "Expected Name '%s', got %s", tt.expected.Name, merged.Name)
			assert.Equal(t, tt.expected.Age, merged.Age, "Expected Age %d, got %d", tt.expected.Age, merged.Age)
			assert.Equal(t, tt.expected.Email, merged.Email, "Expected Email '%s', got %s", tt.expected.Email, merged.Email)
		})
	}
}

// Test types for nested field functionality
type Company struct {
	Name    string
	Address Address
}

type Employee struct {
	User
	Company Company
	Salary  int
}

type Department struct {
	Name      string
	Manager   Employee
	Employees []Employee
}

func TestGetNested(t *testing.T) {
	person := Person{
		User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
		Active:  true,
	}
	personPtr := &person

	employee := Employee{
		User:    User{Name: "Bob", Age: 30, Email: "bob@company.com"},
		Company: Company{Name: "Tech Corp", Address: Address{Street: "456 Tech Ave", City: "San Francisco", Country: "USA"}},
		Salary:  75000,
	}

	tests := []struct {
		name     string
		input    interface{}
		field    string
		expected interface{}
		found    bool
	}{
		{
			name:     "single level nested field",
			input:    person,
			field:    "Address.City",
			expected: "New York",
			found:    true,
		},
		{
			name:     "case sensitive",
			input:    person,
			field:    "Address.Street",
			expected: "123 Main St",
			found:    true,
		},
		{
			name:     "embedded struct field",
			input:    person,
			field:    "User.Name",
			expected: "Alice",
			found:    true,
		},
		{
			name:     "non-existing nested field",
			input:    person,
			field:    "Address.phone",
			expected: nil,
			found:    false,
		},
		{
			name:     "invalid path",
			input:    person,
			field:    "User.Name.invalid",
			expected: nil,
			found:    false,
		},
		{
			name:     "pointer input",
			input:    personPtr,
			field:    "Address.Country",
			expected: "USA",
			found:    true,
		},
		{
			name:     "complex nested structure",
			input:    employee,
			field:    "Company.Address.City",
			expected: "San Francisco",
			found:    true,
		},
		{
			name:     "user email in complex structure",
			input:    employee,
			field:    "User.Email",
			expected: "bob@company.com",
			found:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := Get(tt.input, tt.field)
			assert.Equal(t, tt.found, found, "Expected found to be %v, got %v", tt.found, found)
			if tt.found {
				assert.Equal(t, tt.expected, value, "Expected %v, got %v", tt.expected, value)
			}
		})
	}
}

func TestSetNested(t *testing.T) {
	tests := []struct {
		name            string
		input           interface{}
		field           string
		value           interface{}
		expectedCity    string
		expectedStreet  string
		expectedAge     int
		expectedCountry string
		expectError     bool
	}{
		{
			name: "setting single level nested field",
			input: &Person{
				User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
				Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
				Active:  true,
			},
			field:           "Address.City",
			value:           "Los Angeles",
			expectedCity:    "Los Angeles",
			expectedStreet:  "123 Main St",
			expectedAge:     25,
			expectedCountry: "USA",
			expectError:     false,
		},
		{
			name: "case sensitive",
			input: &Person{
				User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
				Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
				Active:  true,
			},
			field:           "Address.Street",
			value:           "789 New Ave",
			expectedCity:    "New York",
			expectedStreet:  "789 New Ave",
			expectedAge:     25,
			expectedCountry: "USA",
			expectError:     false,
		},
		{
			name: "setting embedded struct field",
			input: &Person{
				User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
				Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
				Active:  true,
			},
			field:           "User.Age",
			value:           30,
			expectedCity:    "New York",
			expectedStreet:  "123 Main St",
			expectedAge:     30,
			expectedCountry: "USA",
			expectError:     false,
		},
		{
			name: "non-existing nested field",
			input: &Person{
				User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
				Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
				Active:  true,
			},
			field:           "Address.phone",
			value:           "123-456-7890",
			expectedCity:    "New York",
			expectedStreet:  "123 Main St",
			expectedAge:     25,
			expectedCountry: "USA",
			expectError:     true,
		},
		{
			name: "type mismatch",
			input: &Person{
				User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
				Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
				Active:  true,
			},
			field:           "Address.City",
			value:           42,
			expectedCity:    "New York",
			expectedStreet:  "123 Main St",
			expectedAge:     25,
			expectedCountry: "USA",
			expectError:     true,
		},
		{
			name: "pointer input",
			input: &Person{
				User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
				Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
				Active:  true,
			},
			field:           "Address.Country",
			value:           "Canada",
			expectedCity:    "New York",
			expectedStreet:  "123 Main St",
			expectedAge:     25,
			expectedCountry: "Canada",
			expectError:     false,
		},
		{
			name: "complex nested structure",
			input: &Employee{
				User:    User{Name: "Bob", Age: 30, Email: "bob@company.com"},
				Company: Company{Name: "Tech Corp", Address: Address{Street: "456 Tech Ave", City: "San Francisco", Country: "USA"}},
				Salary:  75000,
			},
			field:           "Company.Address.City",
			value:           "Seattle",
			expectedCity:    "Seattle",
			expectedStreet:  "456 Tech Ave",
			expectedAge:     30,
			expectedCountry: "USA",
			expectError:     false,
		},
		{
			name: "user email in complex structure",
			input: &Employee{
				User:    User{Name: "Bob", Age: 30, Email: "bob@company.com"},
				Company: Company{Name: "Tech Corp", Address: Address{Street: "456 Tech Ave", City: "San Francisco", Country: "USA"}},
				Salary:  75000,
			},
			field:           "User.Email",
			value:           "bob.new@company.com",
			expectedCity:    "San Francisco",
			expectedStreet:  "456 Tech Ave",
			expectedAge:     30,
			expectedCountry: "USA",
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Set(tt.input, tt.field, tt.value)

			if tt.expectError {
				assert.Error(t, err, "Expected error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got: %v", err)

				// Check the modified struct
				if personResult, ok := tt.input.(*Person); ok {
					assert.Equal(t, tt.expectedCity, personResult.Address.City, "Expected City to be '%s', got %s", tt.expectedCity, personResult.Address.City)
					assert.Equal(t, tt.expectedStreet, personResult.Address.Street, "Expected Street to be '%s', got %s", tt.expectedStreet, personResult.Address.Street)
					assert.Equal(t, tt.expectedAge, personResult.User.Age, "Expected Age to be %d, got %d", tt.expectedAge, personResult.User.Age)
					assert.Equal(t, tt.expectedCountry, personResult.Address.Country, "Expected Country to be '%s', got %s", tt.expectedCountry, personResult.Address.Country)
				} else if employeeResult, ok := tt.input.(*Employee); ok {
					assert.Equal(t, tt.expectedCity, employeeResult.Company.Address.City, "Expected City to be '%s', got %s", tt.expectedCity, employeeResult.Company.Address.City)
					assert.Equal(t, tt.expectedStreet, employeeResult.Company.Address.Street, "Expected Street to be '%s', got %s", tt.expectedStreet, employeeResult.Company.Address.Street)
					assert.Equal(t, tt.expectedAge, employeeResult.User.Age, "Expected Age to be %d, got %d", tt.expectedAge, employeeResult.User.Age)
					assert.Equal(t, tt.expectedCountry, employeeResult.Company.Address.Country, "Expected Country to be '%s', got %s", tt.expectedCountry, employeeResult.Company.Address.Country)
				}
			}
		})
	}
}

func TestHasNested(t *testing.T) {
	person := Person{
		User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
		Active:  true,
	}
	personPtr := &person

	employee := Employee{
		User: User{Name: "Bob", Age: 30, Email: "bob@company.com"},
		Company: Company{
			Name: "Tech Corp",
			Address: Address{
				Street:  "456 Tech Ave",
				City:    "San Francisco",
				Country: "USA",
			},
		},
		Salary: 75000,
	}

	tests := []struct {
		name     string
		input    interface{}
		field    string
		expected bool
	}{
		{
			name:     "existing nested field - address.city",
			input:    person,
			field:    "Address.City",
			expected: true,
		},
		{
			name:     "existing nested field - address.street",
			input:    person,
			field:    "Address.Street",
			expected: true,
		},
		{
			name:     "existing nested field - user.name",
			input:    person,
			field:    "User.Name",
			expected: true,
		},
		{
			name:     "existing nested field - user.age",
			input:    person,
			field:    "User.Age",
			expected: true,
		},
		{
			name:     "case sensitive - Address.Country",
			input:    person,
			field:    "Address.Country",
			expected: true,
		},
		{
			name:     "case sensitive - User.Email",
			input:    person,
			field:    "User.Email",
			expected: true,
		},
		{
			name:     "non-existing nested field - Address.Phone",
			input:    person,
			field:    "Address.Phone",
			expected: false,
		},
		{
			name:     "non-existing nested field - User.Phone",
			input:    person,
			field:    "User.Phone",
			expected: false,
		},
		{
			name:     "invalid field path",
			input:    person,
			field:    "invalid.field",
			expected: false,
		},
		{
			name:     "pointer input",
			input:    personPtr,
			field:    "Address.City",
			expected: true,
		},
		{
			name:     "complex nested structure - company.address.city",
			input:    employee,
			field:    "Company.Address.City",
			expected: true,
		},
		{
			name:     "complex nested structure - company.name",
			input:    employee,
			field:    "Company.Name",
			expected: true,
		},
		{
			name:     "complex nested structure - user.email",
			input:    employee,
			field:    "User.Email",
			expected: true,
		},
		{
			name:     "complex nested structure - salary",
			input:    employee,
			field:    "Salary",
			expected: true,
		},
		{
			name:     "complex nested structure - non-existing company.address.phone",
			input:    employee,
			field:    "Company.Address.Phone",
			expected: false,
		},
		{
			name:     "complex nested structure - non-existing user.phone",
			input:    employee,
			field:    "User.Phone",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Has(tt.input, tt.field)
			assert.Equal(t, tt.expected, result, "Expected %v, got %v", tt.expected, result)
		})
	}
}
