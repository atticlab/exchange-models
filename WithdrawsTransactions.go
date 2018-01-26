package exmodels

import (
    "time"
)

const (
    WITH_TX_STATE_CREATED = "Created"
    WITH_TX_STATE_PROCESSING = "Processing"
    WITH_TX_STATE_DONE = "Done"
)

type WithdrawsTransactions struct {
    Id         uint    `gorm:"primary_key"`
    Txid       string  `gorm:"size:255" sql:"default: null"`
    State      string  `gorm:"size:255" sql:"default: null"`
    Fee        float64 `sql:"type:decimal(32,16);"`
    Txbody     string  `sql:"type: text"`
    TxbodyHash string  `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}

func (this *WithdrawsTransactions) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *WithdrawsTransactions) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}
