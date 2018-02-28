package exmodels

import (
    "time"
    "github.com/shopspring/decimal"
)

type WithdrawsComissions struct {
    Id       uint    `gorm:"primary_key"`
    Currency uint
    Fixed    decimal.NullDecimal `sql:"type:decimal(32,16);"`
    Dynamic  decimal.NullDecimal `sql:"type:decimal(5,4);"`

    MinWithdrawAmount decimal.NullDecimal `sql:"type:decimal(32,16);"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}

func (this *WithdrawsComissions) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *WithdrawsComissions) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}
