package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Record struct that is the representation of record model
type Record struct {
	gorm.Model
	Lastname  string    `form:"Lastname" json:"Lastname" binding:"required" gorm:"index:add;size:255"`
	Firstname string    `form:"Firstname" json:"Firstname" binding:"required" gorm:"size:255"`
	OrderType string    `form:"OrderType" json:"OrderType" binding:"required" gorm:"size:255"`
	Address   string    `form:"Address" json:"Address"`
	Phone     string    `form:"Phone" json:"Phone" binding:"required" `
	Phone2    string    `form:"Phone2" json:"Phone2"`
	OrderTime time.Time `form:"OrderTime" json:"OrderTime" time_format:"2006-01-02" gorm:"default:current_timestamp"`
	Notes     string    `form:"Notes" json:"Notes" gorm:"size:1000"`
}

var OrderTypes = []string{"DBepi Bxidni", "DBepi kimnatni", "Bikna metaloplastuk", "Poletu/Zaluzi", "Pidvikonnya", "Mebli", "Inshe"}

//File is representation of file in database
type File struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Path     string
	RecordID uint
}

//FormatedDate return string representating of OrderTime in YYYY-MM-DD format
func (r Record) FormatedDate() string {
	return r.OrderTime.Format("2006-01-02")
}
