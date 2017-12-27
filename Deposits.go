package exmodels

import (
    "time"
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

type Deposits struct {
    Id                   uint    `gorm:"primary_key"`
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

    Account   Accounts //belongs_to
    AccountId uint     //belongs_to
    Member    Members  //belongs_to
    MemberId  uint     //belongs_to
}
