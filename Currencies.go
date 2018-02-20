package exmodels

import (
    "github.com/shopspring/decimal"
    "math"
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
        UnitRatio:     decimal.NewFromFloat(math.Pow10(8)), //1.00 BTC
        IsCoin:        true,
        IsToken:       false,
        Confirmations: 3,
    },
    {
        Id:            3,
        Title:         "Ripple",
        Code:          "XRP",
        Unit:          "ripple",
        UnitRatio:     decimal.NewFromFloat(math.Pow10(0)), //1.00 XRP ???
        IsCoin:        true,
        IsToken:       false,
        Confirmations: 3,
    },
    {
        Id:            4,
        Title:         "Ethereum",
        Code:          "ETH",
        Unit:          "WEI",
        UnitRatio:     decimal.NewFromFloat(math.Pow10(18)), //1.00 ETH
        IsCoin:        true,
        IsToken:       false,
        Confirmations: 6,
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
