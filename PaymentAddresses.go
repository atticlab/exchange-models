package exmodels

import (
    "time"
    "github.com/jinzhu/gorm"
)

type PaymentAddresses struct {
    Id        uint   `gorm:"primary_key"`
    AccountId uint
    Currency  uint8
    Secret    string `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`

    BaseModel `sql:"-"`
}

func NewPaymentAddress(conn *gorm.DB, accountId uint, currency uint8) (*PaymentAddresses, error) {
    createdAt := time.Now()

    pa := &PaymentAddresses{
        AccountId:        accountId,
        Currency:            currency,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,

        BaseModel: BaseModel{MySQLConnection: conn},
    }

    return pa, nil
}

func (this *PaymentAddresses) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *PaymentAddresses) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}
