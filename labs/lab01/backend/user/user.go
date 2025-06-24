package user

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidEmail is returned when the email format is invalid
	ErrInvalidEmail = errors.New("invalid email format")
	// ErrInvalidAge is returned when the age is invalid
	ErrInvalidAge = errors.New("invalid age: must be between 0 and 150")
	// ErrEmptyName is returned when the name is empty
	ErrEmptyName = errors.New("name cannot be empty")
)

// User represents a user in the system
type User struct {
	Name  string
	Age   int
	Email string
}

// NewUser creates a new user with validation
func NewUser(name string, age int, email string) (*User, error) {
	// TODO: Implement user creation with validation
	if email == "" || !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	} else if age < 0 || age > 150 {
		return nil, ErrInvalidAge
	} else if name == "" {
		return nil, ErrEmptyName
	}
	user := new(User)
	user.Age = age
	user.Email = email
	user.Name = name
	return user, nil
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	// TODO: Implement user validation
	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	} else if u.Age < 0 || u.Age > 150 {
		return ErrInvalidAge
	} else if u.Name == "" {
		return ErrEmptyName
	}
	return nil
}

// String returns a string representation of the user
func (u *User) String() string {
	// TODO: Implement string representation
	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return "You have provided an invalid email. Please, try more"
	} else if u.Age < 0 || u.Age > 150 {
		return "You have provided an invalid age. It has to be between 0 and 150 years. Please, try more"
	} else if u.Name == "" {
		return "You have provided an empty name. Please, try more"
	}
	return fmt.Sprintf("{\n\tName: %s\n\tAge: %d\n\tEmail: %s\n}",
		u.Name, u.Age, u.Email)
}

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	// TODO: Implement email validation
	if email != "" && strings.Contains(email, "@") {
		return true
	}
	return false
}
