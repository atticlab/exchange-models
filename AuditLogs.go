package exmodels

import (
    "time"
)

const (
    AUDITLOGS_TYPE_TRANSFER = "Audit::TransferAuditLog"
)

const (
    AUDITABLE_TYPE_DEPOSIT = "Deposit"
)

type AuditLogs struct {
    Id            uint   `gorm:"primary_key"`
    Type          string `gorm:"size:255" sql:"default: null"`
    OperatorId    uint   `sql:"default: null"`
    AuditableId   uint   `sql:"default: null"`
    AuditableType string `gorm:"size:255" sql:"default: null"`
    SourceState   string `gorm:"size:255" sql:"default: null"`
    TargetState   string `gorm:"size:255" sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}

func (this *AuditLogs) BeforeCreate() (err error) {
    time := time.Now()
    this.CreatedAt = &time
    return
}

func (this *AuditLogs) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}