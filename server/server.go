package server

import (
	"github.com/ajm896/adlibai/db"
)

// StartServer starts the server
func StartServer() {
	db := db.ConnectDB()
	println("Server is running! %s", db)
}
