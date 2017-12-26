package exmodels

import (
    "golang.org/x/crypto/bcrypt"
    "time"
    "errors"
    "github.com/jinzhu/gorm"
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

    BaseModel `sql:"-"`
}

func NewIdentity(conn *gorm.DB, email string, password string) (*Identities, error) {
    createdAt := time.Now()

    identity := &Identities{
        Email:          email,
        PasswordDigest: "",
        RetryCount:     0,
        IsActive:       false,
        IsLocked:       false,
        CreatedAt:      &createdAt,
        UpdatedAt:      &createdAt,

        BaseModel: BaseModel{MySQLConnection: conn},
    }

    err := identity.HashPassword(password)

    if err != nil {
        return nil, err
    }

    return identity, nil
}

func (this *Identities) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *Identities) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}

func (this *Identities) CreateWithTx(tx *gorm.DB) *gorm.DB {
    return tx.Create(&this)
}

func (this *Identities) SaveWithTx(tx *gorm.DB) *gorm.DB {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return tx.Save(&this)
}

func (this *Identities) HashPassword(password string) error {
    hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    this.PasswordDigest = string(hashByte)
    return err
}

func (this *Identities) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(this.PasswordDigest), []byte(password))
}

func GetIdentityByEmail(conn *gorm.DB, email string) (*Identities, error) {
    if email == "" {
        return nil, errors.New("Empty identity email")
    }

    var identityObj Identities

    result := conn.Where("email = ?", email).First(&identityObj)
    if result.Error != nil {
        if result.RecordNotFound() {
            return nil, nil
        }

        return nil, result.Error
    }

    identityObj.BaseModel.MySQLConnection = conn

    return &identityObj, nil
}
