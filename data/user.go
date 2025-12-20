package data

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
// This is an example model demonstrating CRUD operations.
type User struct {
	ID        int       `db:"id,omitempty" json:"id"`
	Email     string    `db:"email" json:"email"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Password  string    `db:"password" json:"-"` // Never expose password in JSON
	Active    bool      `db:"active" json:"active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// UserTable is the database table name for users
const UserTable = "users"

// ErrUserNotFound is returned when a user cannot be found
var ErrUserNotFound = errors.New("user not found")

// ErrDuplicateEmail is returned when trying to create a user with an existing email
var ErrDuplicateEmail = errors.New("email already exists")

// PaginatedUsers holds paginated user data
type PaginatedUsers struct {
	Users      []User `json:"users"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      uint64 `json:"total"`
	TotalPages int    `json:"total_pages"`
}

// GetAllUsers returns all users from the database
// Note: For large datasets, use GetUsersPaginated instead
func (m *Models) GetAllUsers() ([]User, error) {
	collection := upper.Collection(UserTable)
	var users []User

	res := collection.Find()
	err := res.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUsersPaginated returns a paginated list of users
func (m *Models) GetUsersPaginated(page, perPage int) (*PaginatedUsers, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100 // Cap at 100 to prevent excessive queries
	}

	collection := upper.Collection(UserTable)
	var users []User

	res := collection.Find()

	// Get total count
	total, err := res.Count()
	if err != nil {
		return nil, err
	}

	// Calculate pagination
	offset := (page - 1) * perPage
	totalPages := int((total + uint64(perPage) - 1) / uint64(perPage))

	// Get paginated results
	err = res.Offset(offset).Limit(perPage).All(&users)
	if err != nil {
		return nil, err
	}

	return &PaginatedUsers{
		Users:      users,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

// GetUserByID retrieves a user by their ID
func (m *Models) GetUserByID(id int) (*User, error) {
	collection := upper.Collection(UserTable)
	var user User

	res := collection.Find("id", id)
	err := res.One(&user)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

// GetUserByEmail retrieves a user by their email address
func (m *Models) GetUserByEmail(email string) (*User, error) {
	collection := upper.Collection(UserTable)
	var user User

	res := collection.Find("email", email)
	err := res.One(&user)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

// CreateUser creates a new user with hashed password
func (m *Models) CreateUser(u *User) (int, error) {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	u.Password = string(hashedPassword)
	u.Active = true
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	collection := upper.Collection(UserTable)
	res, err := collection.Insert(u)
	if err != nil {
		return 0, err
	}

	return getInsertID(res.ID()), nil
}

// UpdateUser updates an existing user
func (m *Models) UpdateUser(u *User) error {
	u.UpdatedAt = time.Now()

	collection := upper.Collection(UserTable)
	res := collection.Find("id", u.ID)

	return res.Update(u)
}

// DeleteUser removes a user from the database
func (m *Models) DeleteUser(id int) error {
	collection := upper.Collection(UserTable)
	res := collection.Find("id", id)

	return res.Delete()
}

// CheckPassword verifies that the provided password matches the user's hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// Authenticate verifies email and password and returns the user if valid
func (m *Models) Authenticate(email, password string) (*User, error) {
	user, err := m.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	if !user.Active {
		return nil, errors.New("account is not active")
	}

	return user, nil
}

// GetUserByIDWithContext retrieves a user with context support for timeouts
func (m *Models) GetUserByIDWithContext(ctx context.Context, id int) (*User, error) {
	// Check if context is already cancelled
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return m.GetUserByID(id)
}
