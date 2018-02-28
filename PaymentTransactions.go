package exmodels

import (
    "time"
    "github.com/shopspring/decimal"
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
    Amount        decimal.NullDecimal `sql:"type:decimal(32,16);"`
    Confirmations uint    `sql:"default: null"`
    Address       string  `gorm:"size:255" sql:"default: null"`
    State         uint    `sql:"default: 0"`
    AasmState     string  `gorm:"size:255" sql:"default: null"`

    Currency uint
    Type     string `gorm:"size:255" sql:"default: null"`
    Txout    uint   `sql:"default: 0"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    ReceiveAt *time.Time `sql:"default: null"`
    DontAt    *time.Time `sql:"default: null"`

    //Deposit        Deposits                                     //has one
    //PaymentAddress PaymentAddresses `gorm:"ForeignKey:address"` //belongs to
}

func (this *PaymentTransactions) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *PaymentTransactions) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}
