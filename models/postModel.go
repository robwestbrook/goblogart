package models

import "gorm.io/gorm"

// Post model
type Post struct {
	gorm.Model
	Title			string
	Body			string
	Likes			int
	Draft			bool
	Author		string
	UserID		uint	`gorm:"user_id"`
}