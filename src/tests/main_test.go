package main

import "testing"

type Users []string

func (u Users) ValidUser(in string) bool {
	for _, usr := range u {
		if usr == in {
			return true
		}
	}
	return false
}

func TestValidUser(t *testing.T) {
	users := Users{"username", "admin"}

	// Test valid users
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
