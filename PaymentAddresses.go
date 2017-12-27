package exmodels

import (
    "time"
)

type PaymentAddresses struct {
    Id        uint   `gorm:"primary_key"`

    //belongs_to
    Account   Accounts
    AccountId uint

    //has many
    Transactions []PaymentTransactions `gorm:"ForeignKey:address"`

    Address   string `gorm:"size:255" sql:"default: null"`
    Currency  uint8
    Secret    string `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}