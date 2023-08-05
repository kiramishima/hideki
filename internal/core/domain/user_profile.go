package domain

import (
	"time"
)

// UserProfile struct
type UserProfile struct {
	ID        int16     `json:"id" db:"user_id"`
	UserName  string    `json:"username" db:"username" bson:"username"`
	FullName  string    `json:"full_name" db:"full_name" bson:"full_name"`
	Role      string    `json:"role" db:"role,omitempty,string" bson:"role,omitempty,string"`
	Bio       string    `json:"bio" db:"bio" bson:"bio"`
	Picture   string    `json:"picture" db:"picture" bson:"picture"`
	CreatedAt time.Time `json:"-" db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"-" db:"deleted_at" bson:"deleted_at"`
}
