package exmodels

import (
    "time"
)

type Accounts struct {
    Id                          uint    `gorm:"primary_key"`

    //belongs_to
    Member                      Members
    MemberId                    uint

    //has many
    PaymentAddresses            []PaymentAddresses

    Currency                    uint8
    Balance                     float32 `sql:"type:decimal(10,2);"`
    Locked                      float32 `sql:"type:decimal(10,2);"`
    In                          float32 `sql:"type:decimal(10,2); default: null"`
    Out                         float32 `sql:"type:decimal(10,2); default: null"`
    DefaultWithdrawFundSourceId uint    `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}
