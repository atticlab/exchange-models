package exmodels

import (
	"time"
)

type PaymentAddresses struct {
	Id          uint   `gorm:"primary_key"`
	Address     string `gorm:"size:255" sql:"default: null"`
	Currency    uint
	Secret      string `gorm:"size:255" sql:"default: null"`
	ExpectedTag string `gorm:"size:255" sql:"default: null"`

	CreatedAt *time.Time `sql:"default: null"`
	UpdatedAt *time.Time `sql:"default: null"`

	//Account      Accounts                                          //belongs_to
	AccountId uint //belongs_to
	//Transactions []PaymentTransactions `gorm:"ForeignKey:address"` //has many
}

func (this *PaymentAddresses) BeforeCreate() (err error) {
	time := time.Now()
	this.CreatedAt = &time
	return
}

func (this *PaymentAddresses) BeforeUpdate() (err error) {
	time := time.Now()
	this.UpdatedAt = &time
	return
}