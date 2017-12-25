package exmodels

import "github.com/jinzhu/gorm"

type BaseModel struct {
    MySQLConnection *gorm.DB
}