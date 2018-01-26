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
    Id               uint   `gorm:"primary_key"`
    Sn               string `gorm:"size:255"`
    DisplayName      string `gorm:"size:255" sql:"default: null"`
    Email            string `gorm:"size:255"`
    IdentityId       uint   `sql:"default: null"`
    State            uint   `sql:"default: null"`
    Activated        bool   `sql:"default: null"`
    CountryCode      uint   `sql:"default: null"`
    PhoneNumber      string `gorm:"size:255" sql:"default: null"`
    IsPhoneConfirmed bool
    Disabled         bool
    ApiDisabled      bool
    Nickname         string `gorm:"size:255" sql:"default: null"`
    RefId            uint   `sql:"default: null"`

    CreatedAt *time.Time
    UpdatedAt *time.Time

    //Accounts []Accounts //has many
    //Deposits []Deposits //has many
}

func (this *Members) generateSn() {
    rand.Seed(time.Now().UnixNano())
    var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, SN_LENGTH)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    this.Sn = strings.ToUpper(SN_PREFIX + string(b) + SN_POSTFIX)

    return
}

func (this *Members) BeforeCreate() (err error) {
    this.generateSn()

    time := time.Now()
    this.CreatedAt = &time

    return nil
}

func (this *Members) BeforeUpdate() (err error) {
    time := time.Now()
    this.UpdatedAt = &time
    return
}