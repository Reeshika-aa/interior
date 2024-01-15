package main

import (
	"main/db"
	"main/router"
	"main/tables"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.DB.AutoMigrate(&tables.User{})
	db.DB.AutoMigrate(&tables.Projects{})
	db.DB.AutoMigrate(&tables.Designer{})
	r := router.Router()
	r.Run()
}
