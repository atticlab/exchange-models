package exmodels

import (
    "time"
    "github.com/shopspring/decimal"
)

const TYPE_ACCOUNT = "Account"

type Accounts struct {
    Id       uint    `gorm:"primary_key"`
    Currency uint
    Balance  decimal.NullDecimal `sql:"type:decimal(32,16);"`
    Locked   decimal.NullDecimal `sql:"type:decimal(32,16);"`
    In       decimal.NullDecimal `sql:"type:decimal(32,16); default: null"`
    Out      decimal.NullDecimal `sql:"type:decimal(32,16); default: null"`

    DefaultWithdrawFundSourceId uint `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`

    //Member           Members            //belongs_to
    MemberId         uint               //belongs_to
    //PaymentAddresses []PaymentAddresses //has many
}

func (this *Accounts) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *Accounts) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}