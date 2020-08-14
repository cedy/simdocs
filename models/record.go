package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Record struct {
	gorm.Model
	Lastname  string    `json:"lastname" gorm:"index:add;size:255"`
	Firstname string    `json:"firstname gorm:"size:255"`
	OrderType string    `json:"orderType gorm:"size:255"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Phone2    string    `json:"phone2"`
	OrderTime time.Time `json:"orderTime"`
	Notes     string    `json:"notes" gorm:"size:1000"`
	Files     []File
}

type File struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Data     []byte
	Type     string
	RecordID uint
}
