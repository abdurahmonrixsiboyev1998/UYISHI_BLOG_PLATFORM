package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "log"
)

func Connect() (*gorm.DB, error) {
    dsn := os.Getenv("POSTGRES_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
        return nil, err
    }
    log.Println("Connected to the database successfully")
    return db, nil
}
