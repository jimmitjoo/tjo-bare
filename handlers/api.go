package handlers

import (
	"encoding/json"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// APIResponse is a standard response wrapper for API endpoints
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// writeJSON writes a JSON response with the given status code
func (h *Handlers) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// errorJSON writes an error response
func (h *Handlers) errorJSON(w http.ResponseWriter, status int, message string) {
	h.writeJSON(w, status, APIResponse{
		Success: false,
		Error:   message,
	})
}

// validationErrorJSON writes a validation error response with field-level errors
func (h *Handlers) validationErrorJSON(w http.ResponseWriter, errors map[string]string) {
	h.writeJSON(w, http.StatusBadRequest, APIResponse{
		Success: false,
		Error:   "Validation failed",
		Data:    errors,
	})
}

// successJSON writes a success response with data
func (h *Handlers) successJSON(w http.ResponseWriter, status int, data interface{}) {
	h.writeJSON(w, status, APIResponse{
		Success: true,
		Data:    data,
	})
}

// APIGetUsers returns paginated users as JSON
// GET /api/users?page=1&per_page=10
func (h *Handlers) APIGetUsers(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page := 1
	perPage := 10

	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if pp := r.URL.Query().Get("per_page"); pp != "" {
		if parsed, err := strconv.Atoi(pp); err == nil && parsed > 0 {
			perPage = parsed
		}
	}

	result, err := h.Models.GetUsersPaginated(page, perPage)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching users: %v", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	h.successJSON(w, http.StatusOK, result)
}

// APIGetUser returns a single user by ID
// GET /api/users/{id}
func (h *Handlers) APIGetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.errorJSON(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.Models.GetUserByID(id)
	if err != nil {
		h.errorJSON(w, http.StatusNotFound, "User not found")
		return
	}

	h.successJSON(w, http.StatusOK, user)
}

// CreateUserRequest is the request body for creating a user
type CreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// APICreateUser creates a new user
// POST /api/users
func (h *Handlers) APICreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorJSON(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate input
	v := h.App.Validator(nil)
	v.Check(req.Email != "", "email", "Email is required")
	v.Check(req.Password != "", "password", "Password is required")
	v.IsEmail("email", req.Email)
	v.MinLength("password", req.Password, 8)
	v.MaxLength("first_name", req.FirstName, 100)
	v.MaxLength("last_name", req.LastName, 100)

	if !v.Valid() {
		h.validationErrorJSON(w, v.Errors)
		return
	}

	// Check if email already exists
	_, err := h.Models.GetUserByEmail(req.Email)
	if err == nil {
		h.errorJSON(w, http.StatusConflict, "Email already exists")
		return
	}

	user := &data.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	id, err := h.Models.CreateUser(user)
	if err != nil {
		h.App.ErrorLog.Printf("Error creating user: %v", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	user.ID = id
	user.Password = "" // Don't return password
	h.successJSON(w, http.StatusCreated, user)
}

// APIDeleteUser deletes a user by ID
// DELETE /api/users/{id}
// Authorization: Users can only delete their own account.
// To add admin functionality, extend the User model with a Role field
// and check for admin role here.
func (h *Handlers) APIDeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.errorJSON(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Get the current user's ID from the session
	currentUserID := h.App.Session.GetInt(r.Context(), "user_id")
	if currentUserID == 0 {
		h.errorJSON(w, http.StatusUnauthorized, "Authentication required")
		return
	}

	// Authorization check: users can only delete their own account
	// To allow admins to delete any user, add a role check here:
	// isAdmin := h.App.Session.GetBool(r.Context(), "is_admin")
	// if currentUserID != id && !isAdmin {
	if currentUserID != id {
		h.errorJSON(w, http.StatusForbidden, "You can only delete your own account")
		return
	}

	if err := h.Models.DeleteUser(id); err != nil {
		h.App.ErrorLog.Printf("Error deleting user: %v", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	h.writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}

// LoginRequest is the request body for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// APILogin handles user authentication
// POST /api/login
func (h *Handlers) APILogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorJSON(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate input before DB lookup
	v := h.App.Validator(nil)
	v.Check(req.Email != "", "email", "Email is required")
	v.Check(req.Password != "", "password", "Password is required")
	v.IsEmail("email", req.Email)

	if !v.Valid() {
		h.validationErrorJSON(w, v.Errors)
		return
	}

	user, err := h.Models.Authenticate(req.Email, req.Password)
	if err != nil {
		h.errorJSON(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Set user in session
	h.App.Session.Put(r.Context(), "user_id", user.ID)

	user.Password = "" // Don't return password
	h.successJSON(w, http.StatusOK, user)
}

// APILogout handles user logout
// POST /api/logout
func (h *Handlers) APILogout(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Remove(r.Context(), "user_id")
	h.App.Session.Destroy(r.Context())

	h.writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "Logged out successfully",
	})
}

// APIHealthCheck returns the health status
// GET /api/health
func (h *Handlers) APIHealthCheck(w http.ResponseWriter, r *http.Request) {
	h.successJSON(w, http.StatusOK, map[string]string{
		"status":  "healthy",
		"version": h.App.Version,
	})
}
