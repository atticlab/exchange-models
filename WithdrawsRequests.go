package exmodels

import (
    "time"
)

const (
    WITH_REQ_STATE_CREATED = "Created"
    WITH_REQ_STATE_SUSPECT = "Suspect"
    WITH_REQ_STATE_PROCESSING = "Processing"
    WITH_REQ_STATE_DONE    = "Done"
)

const TYPE_WITHDRAW = "Withdraw"

type WithdrawsRequests struct {
    Id           uint    `gorm:"primary_key"`
    AccountId    uint
    MemberId     uint
    Currency     uint8
    Address      string  `gorm:"size:255" sql:"default: null"`
    Amount       float64 `sql:"type:decimal(32,16);"`
    Fee          float64 `sql:"type:decimal(32,16);"`
    WithdrawTxId uint
    State        string  `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    DoneAt    *time.Time `sql:"default: null"`
}

func (this *WithdrawsRequests) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *WithdrawsRequests) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}
