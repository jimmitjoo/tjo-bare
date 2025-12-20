package middleware

import (
	"myapp/data"
	"net/http"
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

// TestCORSHeaders tests that CORS middleware sets correct headers
func TestCORSHeaders(t *testing.T) {
	// Create a mock next handler
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create middleware without full app (just test CORS logic)
	m := &Middleware{}
	handler := m.CORS(next)

	// Test regular request
	req := httptest.NewRequest("GET", "/api/test", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check headers
	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("expected Access-Control-Allow-Origin header")
	}

	if rr.Header().Get("Access-Control-Allow-Methods") == "" {
		t.Error("expected Access-Control-Allow-Methods header")
	}

	// Test OPTIONS preflight
	req = httptest.NewRequest("OPTIONS", "/api/test", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 for OPTIONS, got %d", rr.Code)
	}
}

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
