package exmodels

import (
    "github.com/shopspring/decimal"
)

type Currency struct {
    Id            uint `gorm:"primary_key"`
    Title         string
    Code          string
    Unit          string
    UnitRatio     decimal.NullDecimal `sql:"type:decimal(32,16);"`
    IsCoin        bool
    IsToken       bool
    Confirmations uint8
}

//var AcceptedCurrencies = []Currency{
//    {
//        Id:            1,
//        Title:         "Bitcoin",
//        Code:          "BTC",
//        Unit:          "Satoshi",
//        UnitRatio:     decimal.NewFromFloat(math.Pow10(8)), //1.00 BTC
//        IsCoin:        true,
//        IsToken:       false,
//        Confirmations: 3,
//    },
//    {
//        Id:            2,
//        Title:         "Ethereum",
//        Code:          "ETH",
//        Unit:          "WEI",
//        UnitRatio:     decimal.NewFromFloat(math.Pow10(18)), //1.00 ETH
//        IsCoin:        true,
//        IsToken:       false,
//        Confirmations: 6,
//    },
//}

//func GetCurrencyById(id uint8) (*Currency) {
//    for _, cur := range AcceptedCurrencies {
//        if cur.Id == id {
//            return &cur
//        }
//    }
//
//    return nil
//}
//
//func GetCurrencyByCode(code string) (*Currency) {
//    for _, cur := range AcceptedCurrencies {
//        if cur.Code == code {
//            return &cur
//        }
//    }
//
//    return nil
//}
