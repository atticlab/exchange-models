package exmodels

import (
    "time"
    "errors"
    "fmt"
    "github.com/thanhpk/randstr"
    "github.com/jinzhu/gorm"
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

    BaseModel `sql:"-"`
}

func NewToken(conn *gorm.DB, memberId uint, AuthTokenTtl int64) (token *Tokens, err error) {
    createdAt := time.Now()
    expiresAt := time.Unix(time.Now().Unix() + AuthTokenTtl, 0)

    token = &Tokens{
        MemberId:        memberId,
        Type:            "Token::Activation",
        Token:           randstr.Hex(16),
        ExpiresAt:       &expiresAt,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,

        BaseModel: BaseModel{MySQLConnection: conn},
    }

    return token, nil
}

func (this *Tokens) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *Tokens) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}

func GetTokenByMemberId(conn *gorm.DB, memberId uint) (*Tokens, error) {
    if memberId == 0 {
        return nil, errors.New("Empty memberId")
    }

    var tokenObj Tokens

    err := conn.Where("member_id = ?", memberId).First(&tokenObj).Error
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Cannot load token from mysql: %s", err.Error()))
    }

    tokenObj.BaseModel.MySQLConnection = conn

    return &tokenObj, nil
}

func GetTokenByActivationCode(conn *gorm.DB, token string) (*Tokens, error) {
    if token == "" {
        return nil, errors.New("empty token")
    }

    var tokenObj Tokens

    err := conn.Where("token = ?", token).First(&tokenObj).Error
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Cannot load token from mysql: %s", err.Error()))
    }

    tokenObj.BaseModel.MySQLConnection = conn

    return &tokenObj, nil
}