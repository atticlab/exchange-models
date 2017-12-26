package exmodels

import (
    "time"
    "github.com/jinzhu/gorm"
)

type Accounts struct {
    Id                          uint    `gorm:"primary_key"`
    MemberId                    uint
    Currency                    uint8
    Balance                     float32 `sql:"type:decimal(10,2);"`
    Locked                      float32 `sql:"type:decimal(10,2);"`
    In                          float32 `sql:"type:decimal(10,2); default: null"`
    Out                         float32 `sql:"type:decimal(10,2); default: null"`
    DefaultWithdrawFundSourceId uint    `sql:"default: null"`

    CreatedAt *time.Time
    UpdatedAt *time.Time

    BaseModel `sql:"-"`
}

func NewAccount(conn *gorm.DB, memberId uint, currencyId uint8) (account *Accounts, err error) {
    //TODO validate memberId, currencyId
    createdAt := time.Now()

    account = &Accounts{
        MemberId:  memberId,
        Currency:  currencyId,
        CreatedAt: &createdAt,
        UpdatedAt: &createdAt,
        BaseModel: BaseModel{MySQLConnection: conn},
    }

    return account, nil
}

func (this *Accounts) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *Accounts) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}

func (this *Accounts) CreateWithTx(tx *gorm.DB) *gorm.DB {
    return tx.Create(&this)
}

func (this *Accounts) SaveWithTx(tx *gorm.DB) *gorm.DB {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return tx.Save(&this)
}
