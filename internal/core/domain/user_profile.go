package domain

import (
	"database/sql"
	"time"
)

// UserProfileQueryItem struct
type UserProfileQueryItem struct {
	ID        sql.NullInt16  `json:"id" db:"id"`
	UserName  sql.NullString `json:"username" db:"username" bson:"username"`
	FullName  sql.NullString `json:"full_name" db:"full_name" bson:"full_name"`
	Role      sql.NullString `json:"role" db:"role,omitempty,string" bson:"role,omitempty,string"`
	Bio       sql.NullString `json:"bio" db:"bio" bson:"bio"`
	Picture   sql.NullString `json:"picture" db:"picture" bson:"picture"`
	CreatedAt sql.NullTime   `json:"-" db:"created_at" bson:"created_at"`
	UpdatedAt sql.NullTime   `json:"-" db:"updated_at" bson:"updated_at"`
	DeletedAt sql.NullTime   `json:"-" db:"deleted_at" bson:"deleted_at"`
}

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
