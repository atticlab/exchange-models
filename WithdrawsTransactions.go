package exmodels

import (
    "time"
)

type WithdrawsTransactions struct {
    Id         uint    `gorm:"primary_key"`
    Txid       string  `gorm:"size:255" sql:"default: null"`
    Fee        float64 `sql:"type:decimal(32,16);"`
    Txbody     string  `sql:"type: text"`
    TxbodyHash string  `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}
