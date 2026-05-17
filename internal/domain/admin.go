package domain

import (
    "time"
    "github.com/google/uuid"
)

type Admin struct {
    ID           uuid.UUID `db:"id" json:"id"`
    Name         string    `db:"name" json:"name"`
    Email        string    `db:"email" json:"email"`
    PasswordHash string    `db:"password_hash" json:"password_hash"`
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
}