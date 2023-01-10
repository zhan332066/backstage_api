//main.go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mooce_api/Config"
	"mooce_api/Models"
	"mooce_api/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Customer{}, &Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
