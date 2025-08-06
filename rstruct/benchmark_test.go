package rstruct

import (
	"testing"
)

type BenchmarkUser struct {
	Name  string
	Age   int
	Email string
	Phone string
	City  string
	State string
	Zip   string
}

func BenchmarkGet(b *testing.B) {
	user := BenchmarkUser{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
		Phone: "123-456-7890",
		City:  "New York",
		State: "NY",
		Zip:   "10001",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Get(user, "Name")
		Get(user, "Age")
		Get(user, "Email")
		Get(user, "Phone")
		Get(user, "City")
		Get(user, "State")
		Get(user, "Zip")
	}
}

func BenchmarkSet(b *testing.B) {
	user := BenchmarkUser{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
		Phone: "123-456-7890",
		City:  "New York",
		State: "NY",
		Zip:   "10001",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Set(&user, "Name", "Bob")
		Set(&user, "Age", 30)
		Set(&user, "Email", "bob@example.com")
		Set(&user, "Phone", "987-654-3210")
		Set(&user, "City", "Los Angeles")
		Set(&user, "State", "CA")
		Set(&user, "Zip", "90210")
	}
}

func BenchmarkHas(b *testing.B) {
	user := BenchmarkUser{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
		Phone: "123-456-7890",
		City:  "New York",
		State: "NY",
		Zip:   "10001",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Has(user, "Name")
		Has(user, "Age")
		Has(user, "Email")
		Has(user, "Phone")
		Has(user, "City")
		Has(user, "State")
		Has(user, "Zip")
		Has(user, "nonexistent")
	}
}
