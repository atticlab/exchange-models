package exmodels

import (
    "time"
)

type WithdrawsComissions struct {
    Id       uint    `gorm:"primary_key"`
    Currency uint8
    Fixed    float64 `sql:"type:decimal(32,16);"`
    Dynamic  float64 `sql:"type:decimal(32,16);"`

    CreatedAt *time.Time `sql:"default: null"`
    UpdatedAt *time.Time `sql:"default: null"`
}

//id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
//currency INT(11),
//fixed DECIMAL(32,16),
//dynamic DECIMAL(32,16),
//created_at DATETIME,
//updated_at DATETIME
