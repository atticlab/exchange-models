package exmodels

import (
    "time"
)

type PaymentAddresses struct {
    Id       uint   `gorm:"primary_key"`
    Address  string `gorm:"size:255" sql:"default: null"`
    Currency uint8
    Secret   string `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`

    //Account      Accounts                                          //belongs_to
    //AccountId    uint                                              //belongs_to
    //Transactions []PaymentTransactions `gorm:"ForeignKey:address"` //has many
}
