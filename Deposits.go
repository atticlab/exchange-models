package exmodels

import (
    "time"
)

type Deposits struct {
    Id uint `gorm:"primary_key"`

    //belongs_to
    Account              Accounts
    AccountId            uint
    Member               Members
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
}