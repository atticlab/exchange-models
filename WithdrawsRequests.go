package exmodels

import (
    "time"
)

type WithdrawsRequests struct {
    Id           uint    `gorm:"primary_key"`
    AccountId    uint
    MemberId     uint
    Currency     uint8
    Amount       float64 `sql:"type:decimal(32,16);"`
    Fee          float64 `sql:"type:decimal(32,16);"`
    WithdrawTxId uint
    State        string  `gorm:"size:255" sql:"default: null"`
    Type         string  `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}
