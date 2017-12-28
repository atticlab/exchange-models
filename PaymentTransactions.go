package exmodels

import (
    "time"
)

const (
    PAYMENT_TX_STATE_UNCONFIRM  = "unconfirm"
    PAYMENT_TX_STATE_CONFIRMING = "confirming"
    PAYMENT_TX_STATE_CONFIRMED  = "confirmed"
)

const PAYMENT_TX_TYPE_NORMAL = "PaymentTransaction::Normal"

type PaymentTransactions struct {
    Id            uint    `gorm:"primary_key"`
    Txid          string  `gorm:"size:255" sql:"default: null"`
    Amount        float32 `sql:"type:decimal(10,2);"`
    Confirmations uint    `sql:"default: null"`
    Address       string  `gorm:"size:255" sql:"default: null"`
    State         uint    `sql:"default: null"`
    AasmState     string  `gorm:"size:255" sql:"default: null"`

    Currency uint8
    Type     string `gorm:"size:255" sql:"default: null"`
    Txout    uint   `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    ReceiveAt *time.Time `sql:"default: null"`
    DontAt    *time.Time `sql:"default: null"`

    Deposit        Deposits                                     //has one
    PaymentAddress PaymentAddresses `gorm:"ForeignKey:address"` //belongs to
}
