package middleware

import (
	"myapp/data"
	"net/http/httptest"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestGetUserFromContext tests retrieving user from context
func TestGetUserFromContext(t *testing.T) {
	// Test with no user in context
	req := httptest.NewRequest("GET", "/", nil)
	user := GetUserFromContext(req)
	if user != nil {
		t.Error("expected nil user when not in context")
	}
}

// Note: CORS is handled by the gemquick framework's api.CORS() middleware.
// See gemquick/api/middleware.go for CORS implementation and tests.

// TestContextKey tests context key uniqueness
func TestContextKey(t *testing.T) {
	// Ensure context key is properly typed
	var key contextKey = "test"
	if string(key) != "test" {
		t.Error("context key conversion failed")
	}

	// Ensure UserContextKey is defined
	if UserContextKey == "" {
		t.Error("UserContextKey should not be empty")
	}
}

// TestUserModel tests User struct methods
func TestUserFullName(t *testing.T) {
	user := &data.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	fullName := user.FullName()
	if fullName != "John Doe" {
		t.Errorf("expected 'John Doe', got '%s'", fullName)
	}
}

// TestCheckPassword tests password verification
func TestCheckPassword(t *testing.T) {
	hashed, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	user := &data.User{Password: string(hashed)}

	if !user.CheckPassword("password123") {
		t.Error("correct password should match")
	}
	if user.CheckPassword("wrongpassword") {
		t.Error("wrong password should not match")
	}
	if user.CheckPassword("") {
		t.Error("empty password should not match")
	}
}
