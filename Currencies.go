package exmodels

type Currency struct {
    Title string
    Code  string
    Id    uint8
}

var AcceptedCurrencies = []Currency{
    {
        Title: "Yuan",
        Code:  "CNY",
        Id:    1,
    },
    {
        Title: "Bitcoin",
        Code:  "BTC",
        Id:    2,
    },
}
