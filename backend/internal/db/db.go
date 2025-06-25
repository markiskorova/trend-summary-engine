package db

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    dsn := os.Getenv("DATABASE_URL")
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to DB:", err)
    }

    err = DB.AutoMigrate(&User{}, &Article{})
    if err != nil {
        log.Fatal("failed to auto-migrate models:", err)
    }
}