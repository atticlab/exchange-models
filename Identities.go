package exmodels

import (
    "time"
)

type Identities struct {
    Id             uint       `gorm:"primary_key"`
    Email          string     `gorm:"size:255"`
    PasswordDigest string     `gorm:"size:255"`
    IsActive       bool       `sql:"default: null"`
    RetryCount     int        `sql:"default: null"`
    IsLocked       bool       `sql:"default: null"`
    LockedAt       *time.Time `sql:"default: null"`
    LastVerifyAt   *time.Time `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}