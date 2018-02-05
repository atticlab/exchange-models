package exmodels

import (
    "github.com/shopspring/decimal"
)

type Currency struct {
    Id            uint8
    Title         string
    Code          string
    Unit          string
    UnitRatio     decimal.Decimal //big used reserved for WEI
    IsCoin        bool
    IsToken       bool
    Confirmations uint8
}

var AcceptedCurrencies = []Currency{
    {
        Id:        1,
        Title:     "Yuan",
        Code:      "CNY",
        Unit:      "Yuan",
        UnitRatio: decimal.NewFromFloat(1),
        IsCoin:    false,
        IsToken:   false,
    },
    {
        Id:            2,
        Title:         "Bitcoin",
        Code:          "BTC",
        Unit:          "Satoshi",
        UnitRatio:     decimal.NewFromFloat(100000000), //1.00 BTC
        IsCoin:        true,
        IsToken:       false,
        Confirmations: 3,
    },
}

func GetCurrencyById(id uint8) (*Currency) {
    for _, cur := range AcceptedCurrencies {
        if cur.Id == id {
            return &cur
        }
    }

    return nil
}

func GetCurrencyByCode(code string) (*Currency) {
    for _, cur := range AcceptedCurrencies {
        if cur.Code == code {
            return &cur
        }
    }

    return nil
}
