package exmodels

import (
    "time"
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
