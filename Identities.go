package models

import (
    "golang.org/x/crypto/bcrypt"
    "time"
    "../di"
)

type Identities struct {
    Id                  uint             `gorm:"primary_key"`
    Email               string           `gorm:"size:255"`
    PasswordDigest      string           `gorm:"size:255"`
    IsActive            bool             `sql:"default: null"`
    RetryCount          int              `sql:"default: null"`
    IsLocked            bool             `sql:"default: null"`
    LockedAt            *time.Time
    LastVerifyAt        *time.Time
    CreatedAt           *time.Time
    UpdatedAt           *time.Time
}

func (this *Identities) HashPassword(password string) error {
    hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    this.PasswordDigest = string(hashByte)
    return err
}

func (this *Identities) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(this.PasswordDigest), []byte(password))
}

func (this *Identities) FindFirstByEmail(email string) error {
    mysql := di.Get().Mysql

    return mysql.Where("email = ?", email).First(&this).Error
}

func CreateIdentity(email string, password string) (identity *Identities, err error) {

    createdAt := time.Now()

    identity = &Identities{
        Email:           email,
        PasswordDigest:  "",
        RetryCount:      0,
        IsActive:        false,
        IsLocked:        false,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,
    }

    err = identity.HashPassword(password)

    if err != nil {
        return nil, err
    }

    return identity, nil
}

func (this *Identities) Create() error {
    mysql := di.Get().Mysql

    return mysql.Create(&this).Error
}

func (this *Identities) Save() error {
    mysql := di.Get().Mysql

    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return mysql.Save(&this).Error
}