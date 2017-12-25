package models

import (
    "time"
    "math/rand"
    "../di"
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
}

func CreateMember(email string) (member *Members, err error) {

    createdAt := time.Now()

    member = &Members{
        Email:           email,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,
    }

    member.GenerateSn()

    return member, nil
}

func (this *Members) Load(id uint) error {
    mysql := di.Get().Mysql

    return mysql.First(&this, id).Error
}

func (this *Members) GenerateSn() {

    var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, SN_LENGTH)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }

    this.Sn = strings.ToUpper(SN_PREFIX + string(b) + SN_POSTFIX)

    return
}

func (this *Members) FindFirstByEmail(email string) error {
    mysql := di.Get().Mysql

    return mysql.Where("email = ?", email).First(&this).Error
}

func (this *Members) Create() error {
    mysql := di.Get().Mysql

    return mysql.Create(&this).Error
}

func (this *Members) Save() error {
    mysql := di.Get().Mysql

    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return mysql.Save(&this).Error
}