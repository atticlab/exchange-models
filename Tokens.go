package models

import (
    "time"
    "../di"
    "../conf"
    "errors"
    "fmt"
    "github.com/thanhpk/randstr"
)

type Tokens struct {
    Id                  uint             `gorm:"primary_key"`
    Token               string           `gorm:"size:255"`
    MemberId            uint
    IsUsed              bool
    Type                string           `gorm:"size:255"`

    ExpiresAt           *time.Time
    CreatedAt           *time.Time
    UpdatedAt           *time.Time
}

func GetTokenByActivationCode(token string) (*Tokens, error) {
    if token == "" {
        return nil, errors.New("Empty token")
    }

    mysql := di.Get().Mysql
    var tokenObj Tokens

    err := mysql.Where("token = ?", token).First(&tokenObj).Error
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Cannot load token from mysql: %s", err.Error()))
    }

    return &tokenObj, nil
}

func CreateToken(memberId uint) (token *Tokens, err error) {

    createdAt := time.Now()
    expiresAt := time.Unix(time.Now().Unix() + conf.AuthTokenTtl, 0)

    token = &Tokens{
        MemberId:        memberId,
        Type:            "Token::Activation",
        Token:           randstr.Hex(16),
        ExpiresAt:       &expiresAt,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,
    }

    return token, nil
}

func (this *Tokens) Create() error {
    mysql := di.Get().Mysql

    return mysql.Create(&this).Error
}

func (this *Tokens) Save() error {
    mysql := di.Get().Mysql

    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return mysql.Save(&this).Error
}