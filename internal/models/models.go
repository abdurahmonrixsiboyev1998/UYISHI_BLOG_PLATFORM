package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id'`
	Username     string `gorm:"unique" json:"username"`
	Email        string `gorm:"unique" json:"email"`
	PasswordHash string `json:"-"`
	Posts        []Post `json:"posts"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
