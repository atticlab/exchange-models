package exmodels

import (
    "time"
    "github.com/shopspring/decimal"
)

const (
    DEPOSIT_STATE_SUBMITTING = "submitting"
    DEPOSIT_STATE_CANCELLED  = "cancelled"
    DEPOSIT_STATE_SUBMITTED  = "submitted"
    DEPOSIT_STATE_REJECTED   = "rejected"
    DEPOSIT_STATE_ACCEPTED   = "accepted"
    DEPOSIT_STATE_CHECKED    = "checked"
    DEPOSIT_STATE_WARNING    = "warning"
)

const TYPE_DEPOSIT = "Deposit"

type Deposits struct {
    Id                   uint    `gorm:"primary_key"`
    Currency             uint8
    Amount               decimal.NullDecimal `sql:"type:decimal(32,16);"`
    Fee                  decimal.NullDecimal `sql:"type:decimal(32,16);"`
    FundUid              string  `gorm:"size:255" sql:"default: null"`
    FundExtra            string  `gorm:"size:255" sql:"default: null"`
    Txid                 string  `gorm:"size:255" sql:"default: null"`
    State                uint    `sql:"default: 0"`
    AasmState            string  `gorm:"size:255" sql:"default: null"`
    Confirmations        uint    `sql:"default: null"`
    Type                 string  `gorm:"size:255" sql:"default: null"`
    PaymentTransactionId uint    `sql:"default: null"`
    Txout                uint    `sql:"default: 0"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    DoneAt    *time.Time `sql:"default: null"`

    //Account   Accounts //belongs_to
    AccountId uint     //belongs_to
    //Member    Members  //belongs_to
    MemberId  uint     //belongs_to
}

func (this *Deposits) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *Deposits) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}