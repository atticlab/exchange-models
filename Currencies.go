package exmodels

type Currency struct {
    Title string
    Code  string
    Id    uint8
    IsCoin bool
    IsToken bool
    Confirmations uint8
}

var AcceptedCurrencies = []Currency{
    {
        Title: "Yuan",
        Code:  "CNY",
        Id:    1,
        IsCoin: false,
        IsToken: false,
    },
    {
        Title: "Bitcoin",
        Code:  "BTC",
        Id:    2,
        IsCoin: true,
        IsToken: false,
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
