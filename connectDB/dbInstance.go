package connectdb

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {

	db, err := Connect()
	if err != nil {
		fmt.Println("Error:", err)
	}

	DB = db
}
