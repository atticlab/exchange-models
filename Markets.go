package exmodels

import (
    "github.com/shopspring/decimal"
)

type Market struct {
    Id        uint `gorm:"primary_key"`
    Code      string
    BaseUnit  uint
    QuoteUnit uint //ask_currency
    SortOrder uint //bid_currency

    BidFee decimal.NullDecimal `sql:"type:decimal(32,16);"`
    AskFee decimal.NullDecimal `sql:"type:decimal(32,16);"`

    BidFixed uint
    AskFixed uint
}
