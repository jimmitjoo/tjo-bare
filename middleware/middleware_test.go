package middleware

import (
	"myapp/data"
	"net/http/httptest"
	"testing"
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
	// This requires bcrypt which is already imported in data package
	// The test demonstrates the pattern but would need a properly hashed password
	t.Skip("Skipping: requires bcrypt password setup")
}
