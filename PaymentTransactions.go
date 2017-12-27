package exmodels

import (
    "time"
    "github.com/jinzhu/gorm"
    "errors"
)

const (
    PAYMENT_TX_STATE_UNCONFIRM  = "unconfirm"
    PAYMENT_TX_STATE_CONFIRMING = "confirming"
    PAYMENT_TX_STATE_CONFIRMED  = "confirmed"
)

const PAYMENT_TX_TYPE_NORMAL = "PaymentTransaction::Normal"

type PaymentTransactions struct {
    Id            uint    `gorm:"primary_key"`
    Txid          string  `gorm:"size:255" sql:"default: null"`
    Amount        float32 `sql:"type:decimal(10,2);"`
    Confirmations uint    `sql:"default: null"`
    Address       string  `gorm:"size:255" sql:"default: null"`
    State         uint    `sql:"default: null"`
    AasmState     string  `gorm:"size:255" sql:"default: null"`

    Currency uint8
    Type     string `gorm:"size:255" sql:"default: null"`
    Txout    uint   `sql:"default: null"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
    ReceiveAt *time.Time `sql:"default: null"`
    DontAt    *time.Time `sql:"default: null"`

    BaseModel `sql:"-"`
}

func NewPaymentTransaction(conn *gorm.DB, pt *PaymentTransactions) (*PaymentTransactions, error) {
    createdAt := time.Now()

    pt.CreatedAt = &createdAt
    pt.UpdatedAt = &createdAt
    pt.BaseModel = BaseModel{MySQLConnection: conn}

    return pt, nil
}

func (this *PaymentTransactions) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *PaymentTransactions) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}

func (this *PaymentTransactions) CreateInDBTransaction(tx *gorm.DB) *gorm.DB {
    return tx.Create(&this)
}

func (this *PaymentTransactions) SaveInDBTransaction(tx *gorm.DB) *gorm.DB {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return tx.Save(&this)
}

func GetPaymentTransactionByAddress(conn *gorm.DB, address string) (*PaymentTransactions, error) {
    if address == "" {
        return nil, errors.New("empty address")
    }

    var paymentTransactionObj PaymentTransactions

    result := conn.Where(&PaymentTransactions{Address: address}).First(&paymentTransactionObj)
    if result.Error != nil {
        if result.RecordNotFound() {
            return nil, nil
        }

        return nil, result.Error
    }

    paymentTransactionObj.BaseModel.MySQLConnection = conn

    return &paymentTransactionObj, nil
}
