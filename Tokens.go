package exmodels

import (
    "time"
)

const TOKEN_TYPE_ACTIVATION = "Token::Activation"
const TOKEN_TYPE_REFRESH = "Token::Refresh"
const TOKEN_TYPE_PASSWORD_UPDATE = "Token::PasswordUpdate"
const TOKEN_TYPE_PASSWORD_RESET = "Token::PasswordReset"
const TOKEN_TYPE_PHONE_CONFIRMATION = "Token::PhoneConfirmation"
const TOKEN_TYPE_PHONE_2FA_CODE = "Token::Phone2FACode"
const TOKEN_TYPE_WITHDRAW = "Token::Withdraw"

type Tokens struct {
    Id       uint   `gorm:"primary_key"`
    Token    string `gorm:"size:255"`
    MemberId uint
    IsUsed   bool
    Type     string `gorm:"size:255"`
    ReqIp    string `gorm:"size:255"`
    ReqData  string `sql:"type: text"`

    ExpiresAt *time.Time `sql:"default: null"`
    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}

func (this *Tokens) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *Tokens) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}