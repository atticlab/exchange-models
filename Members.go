package exmodels

import (
    "time"
    "math/rand"
    "strings"
)

const SN_LENGTH = 8
const SN_PREFIX = "PEA"
const SN_POSTFIX = "TIO"

type Members struct {
    Id                  uint             `gorm:"primary_key"`
    Sn                  string           `gorm:"size:255"`
    DisplayName         string           `gorm:"size:255" sql:"default: null"`
    Email               string           `gorm:"size:255"`
    IdentityId          uint             `sql:"default: null"`
    State               uint             `sql:"default: null"`
    Activated           bool             `sql:"default: null"`
    CountryCode         uint             `sql:"default: null"`
    PhoneNumber         string           `gorm:"size:255" sql:"default: null"`
    Disabled            bool
    ApiDisabled         bool
    Nickname            string           `gorm:"size:255" sql:"default: null"`

    CreatedAt           *time.Time
    UpdatedAt           *time.Time

    //has many
    Accounts            []Accounts
    Deposits            []Deposits
}

func (this *Members) generateSn() {
    var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, SN_LENGTH)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    this.Sn = strings.ToUpper(SN_PREFIX + string(b) + SN_POSTFIX)

    return
}