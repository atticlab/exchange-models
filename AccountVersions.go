package exmodels

import (
    "time"
)

const (
    REASON_CODE_UNKNOWN          = 0
    REASON_CODE_FIX              = 1
    REASON_CODE_STRIKE_FEE       = 100
    REASON_CODE_STRIKE_ADD       = 110
    REASON_CODE_STRIKE_SUB       = 120
    REASON_CODE_STRIKE_UNLOCK    = 130
    REASON_CODE_ORDER_SUBMIT     = 600
    REASON_CODE_ORDER_CANCEL     = 610
    REASON_CODE_ORDER_FULLFILLED = 620
    REASON_CODE_WITHDRAW_LOCK    = 800
    REASON_CODE_WITHDRAW_UNLOCK  = 810
    REASON_CODE_DEPOSIT          = 1000
    REASON_CODE_WITHDRAW         = 2000
)

const (
    _                         = iota
    FUNS_UNLOCK_FUNDS          //=1
    FUNS_LOCK_FUNDS
    FUNS_PLUS_FUNDS
    FUNS_SUB_FUNDS
    FUNS_UNLOCK_AND_SUB_FUNDS  //=5
)

type AccountsVersions struct {
    Id             uint    `gorm:"primary_key"`
    AccountId      uint
    MemberId       uint
    Reason         uint
    Balance        float64 `sql:"type:decimal(32,16);"`
    Locked         float64 `sql:"type:decimal(32,16);"`
    Fee            float64 `sql:"type:decimal(32,16); default: null"`
    Amount         float64 `sql:"type:decimal(32,16); default: null"`
    ModifiableId   uint
    ModifiableType string

    Currency uint8
    Fun      uint8

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}

func (this *AccountsVersions) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *AccountsVersions) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}
