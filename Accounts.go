package exmodels

import (
    "time"
)

const TYPE_ACCOUNT = "Account"

type Accounts struct {
    Id       uint    `gorm:"primary_key"`
    Currency uint8
    Balance  float64 `sql:"type:decimal(32,16);"`
    Locked   float64 `sql:"type:decimal(32,16);"`
    In       float64 `sql:"type:decimal(32,16); default: null"`
    Out      float64 `sql:"type:decimal(32,16); default: null"`

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