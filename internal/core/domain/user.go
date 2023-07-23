package domain

import (
	"net/mail"
	"strings"
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	UserName  string    `json:"username,omitempty,string" db:"username,omitempty,string"`
	RoleID    string    `json:"-" db:"role_id,omitempty,string"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// NewUser crea un nuevo usuario
func NewUser(username, password, email string) (*User, error) {
	user := &User{
		Email:     email,
		UserName:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Validate valida al usuario
func (user *User) Validate() error {
	if user.UserName == "" || user.Password == "" || user.Email == "" {
		return ErrEmptyUserField
	}

	if strings.ContainsAny(user.UserName, " \t\r\n") || strings.ContainsAny(user.Password, " \t\r\n") {
		return ErrFieldWithSpaces
	}

	if len(user.Password) < 6 {
		return ErrShortPassword
	}

	if len(user.Password) > 72 {
		return ErrLongPassword
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
