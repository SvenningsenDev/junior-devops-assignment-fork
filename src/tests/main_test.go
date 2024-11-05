package main

import (
	"encoding/json"
	"os"
	"testing"
)

type Users []string

func (u *Users) FromJson(in string) error {
	return json.Unmarshal([]byte(in), u)
}

func (u Users) ValidUser(in string) bool {
	for _, usr := range u {
		if usr == in {
			return true
		}
	}
	return false
}

func TestValidUser(t *testing.T) {
	// Read the actual users.json file
	b, err := os.ReadFile("../secrets/users.json")
	if err != nil {
		t.Fatalf("Failed to read users.json: %v", err)
	}

	var users Users
	if err := users.FromJson(string(b)); err != nil {
		t.Fatalf("Failed to parse users.json: %v", err)
	}

	// Test valid users from the actual file
	if !users.ValidUser("username") {
		t.Errorf("Expected 'username' to be a valid user")
	}
	if !users.ValidUser("admin") {
		t.Errorf("Expected 'admin' to be a valid user")
	}

	// Test invalid user
	if users.ValidUser("nonexistent") {
		t.Errorf("Expected 'nonexistent' to be an invalid user")
	}
}
