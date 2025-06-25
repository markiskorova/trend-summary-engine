package db

import "time"

// User represents a registered user.
type User struct {
    ID           uint   `gorm:"primaryKey"`
    Email        string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
}

// Article represents a saved article URL.
type Article struct {
    ID        uint      `gorm:"primaryKey"`
    URL       string    `gorm:"not null"`
    UserID    uint      `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}