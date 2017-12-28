package exmodels

import (
    "time"
)

const TokenTypeActivation = "Token::Activation"

type Tokens struct {
    Id       uint   `gorm:"primary_key"`
    Token    string `gorm:"size:255"`
    MemberId uint
    IsUsed   bool
    Type     string `gorm:"size:255"`

    ExpiresAt *time.Time `sql:"default: null"`
    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}
