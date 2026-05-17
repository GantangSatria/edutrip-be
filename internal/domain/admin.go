package domain

import (
    "time"
    "github.com/google/uuid"
)

type Admin struct {
    ID           uuid.UUID `db:"id"`
    Name         string    `db:"name"`
    Email        string    `db:"email"`
    PasswordHash string    `db:"password_hash"`
    CreatedAt    time.Time `db:"created_at"`
}