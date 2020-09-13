package models

import (
	"github.com/jinzhu/gorm"
	// Using sqllite fir dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//DB is a global connection variable
var DB *gorm.DB

//ConnectDataBase is a method that opens connection to the database
func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(err.Error() + "Failed to connect to database!")
	}

	database.AutoMigrate(&Record{})
	database.AutoMigrate(&File{})

	DB = database
}
