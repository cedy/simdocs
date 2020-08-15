package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Record struct that is the representation of record model
type Record struct {
	gorm.Model
	Lastname  string    `json:"Lastname" binding:"required" gorm:"index:add;size:255"`
	Firstname string    `json:"Firstname" binding:"required" gorm:"size:255"`
	OrderType string    `json:"OrderType" binding:"required" gorm:"size:255"`
	Address   string    `json:"Address"`
	Phone     string    `json:"Phone" binding:"required" `
	Phone2    string    `json:"Phone2"`
	OrderTime time.Time `json:"OrderTime" gorm:"default:current_timestamp"`
	Notes     string    `json:"Notes" gorm:"size:1000"`
	// TODO: handle files after POC
	Files []File
}

//File is representation of file in database
type File struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Data     []byte
	Type     string
	RecordID uint
}
