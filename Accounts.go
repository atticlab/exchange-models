package models

import (
    "time"
    "../di"
)

type Accounts struct {
    Id                          uint             `gorm:"primary_key"`
    MemberId                    uint
    Currency                    uint8
    Balance                     float32          `sql:"type:decimal(10,2);"`
    Locked                      float32          `sql:"type:decimal(10,2);"`
    In                          float32          `sql:"type:decimal(10,2); default: null"`
    Out                         float32          `sql:"type:decimal(10,2); default: null"`
    DefaultWithdrawFundSourceId uint             `sql:"default: null"`

    CreatedAt                   *time.Time
    UpdatedAt                   *time.Time
}

func (this *Accounts) Create() error {
    mysql := di.Get().Mysql

    return mysql.Create(&this).Error
}

func (this *Accounts) Save() error {
    mysql := di.Get().Mysql

    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return mysql.Save(&this).Error
}

func CreateAccount(memberId uint, currencyId uint8) (account *Accounts, err error) {

    //TODO validate memberId, currencyId
    createdAt := time.Now()

    account = &Accounts{
        MemberId:        memberId,
        Currency:        currencyId,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,
    }

    return account, nil
}