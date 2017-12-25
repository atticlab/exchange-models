package exmodels

import (
    "time"
    "math/rand"
    "strings"
    "github.com/jinzhu/gorm"
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

    BaseModel `sql:"-"`
}

func NewMember(conn *gorm.DB, email string) (member *Members, err error) {

    createdAt := time.Now()

    member = &Members{
        Email:           email,
        CreatedAt:       &createdAt,
        UpdatedAt:       &createdAt,

        BaseModel: BaseModel{MySQLConnection: conn},
    }

    member.GenerateSn()

    return member, nil
}

func (this *Members) Create() error {
    return this.BaseModel.MySQLConnection.Create(&this).Error
}

func (this *Members) Save() error {
    updatedAt := time.Now()
    this.UpdatedAt = &updatedAt

    return this.BaseModel.MySQLConnection.Save(&this).Error
}

func (this *Members) Load(id uint) error {
    return this.BaseModel.MySQLConnection.First(&this, id).Error
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
    return this.BaseModel.MySQLConnection.Where("email = ?", email).First(&this).Error
}