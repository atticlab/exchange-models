package exmodels

import (
    "time"
    "github.com/jinzhu/gorm"
)

type Deposits struct {
    Id                   uint    `gorm:"primary_key"`
    AccountId            uint
    MemberId             uint
    Currency             uint8
    Amount               float32 `sql:"type:decimal(10,2);"`
    Fee                  float32 `sql:"type:decimal(10,2);"`
    FundUid              string  `gorm:"size:255" sql:"default: null"`
    FundExtra            string  `gorm:"size:255" sql:"default: null"`
    Txid                 string  `gorm:"size:255" sql:"default: null"`
    State                uint    `sql:"default: null"`
    AasmState            string  `gorm:"size:255" sql:"default: null"`
    Confirmations        uint    `sql:"default: null"`
    Type                 string  `gorm:"size:255" sql:"default: null"`
    PaymentTransactionId uint    `sql:"default: null"`
    Txout                uint    `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    DoneAt    *time.Time `sql:"default: null"`

    BaseModel `sql:"-"`
}

func NewDeposit(conn *gorm.DB, deposit *Deposits) (*Deposits, error) {
    createdAt := time.Now()

    deposit.CreatedAt = &createdAt
    deposit.UpdatedAt = &createdAt
    deposit.BaseModel = BaseModel{MySQLConnection: conn}

    return deposit, nil
}

func (this *Deposits) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *Deposits) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}
