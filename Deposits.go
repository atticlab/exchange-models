package exmodels

import (
    "time"
    "github.com/jinzhu/gorm"
    "errors"
)

type Deposits struct {
    Id                   uint    `gorm:"primary_key"`
    AccountId            uint
    MemberId             uint
    Currency             uint8
    Amount               float32 `sql:"type:decimal(10,2);"`
    Fee                  float32 `sql:"type:decimal(10,2);"`
    FundUid              string  `gorm:"size:255" sql:"default: null"`
    FundExtra            string  `gorm:"size:255" sql:"default: null"`
    Txid                 string  `gorm:"size:255" sql:"default: null"`
    State                uint    `sql:"default: null"`
    AasmState            string  `gorm:"size:255" sql:"default: null"`
    Confirmations        uint    `sql:"default: null"`
    Type                 string  `gorm:"size:255" sql:"default: null"`
    PaymentTransactionId uint    `sql:"default: null"`
    Txout                uint    `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    DoneAt    *time.Time `sql:"default: null"`

    BaseModel `sql:"-"`
}

func NewDeposit(conn *gorm.DB, deposit *Deposits) (*Deposits, error) {
    createdAt := time.Now()

    deposit.CreatedAt = &createdAt
    deposit.UpdatedAt = &createdAt
    deposit.BaseModel = BaseModel{MySQLConnection: conn}

    return deposit, nil
}

func (this *Deposits) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *Deposits) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}

func (this *Deposits) CreateInDBTransaction(tx *gorm.DB) *gorm.DB {
    return tx.Create(&this)
}

func (this *Deposits) SaveInDBTransaction(tx *gorm.DB) *gorm.DB {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return tx.Save(&this)
}

func GetDepositByTxid(conn *gorm.DB, txid string) (*Deposits, error) {
    if txid == "" {
        return nil, errors.New("empty txid")
    }

    var depositObj Deposits

    result := conn.Where(&Deposits{Txid: txid}).First(&depositObj)
    if result.Error != nil {
        if result.RecordNotFound() {
            return nil, nil
        }

        return nil, result.Error
    }

    depositObj.BaseModel.MySQLConnection = conn

    return &depositObj, nil
}

func GetDepositByPaymentTransactionId(conn *gorm.DB, ptid uint) (*Deposits, error) {
    if ptid == 0 {
        return nil, errors.New("empty ptid (Payment Transaction Id)")
    }

    var depositObj Deposits

    result := conn.Where(&Deposits{PaymentTransactionId: ptid}).First(&depositObj)
    if result.Error != nil {
        if result.RecordNotFound() {
            return nil, nil
        }

        return nil, result.Error
    }

    depositObj.BaseModel.MySQLConnection = conn

    return &depositObj, nil
}
