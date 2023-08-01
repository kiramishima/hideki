package domain

import (
	"time"
)

// UserProfile struct
type UserProfile struct {
	ID        int16     `json:"id" db:"user_id"`
	UserName  string    `json:"username" db:"username"`
	FullName  string    `json:"full_name" db:"full_name"`
	Role      string    `json:"role" db:"role,omitempty,string"`
	Bio       string    `json:"bio" db:"bio"`
	Picture   string    `json:"picture" db:"picture"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	DeletedAt time.Time `json:"-" db:"deleted_at"`
}
