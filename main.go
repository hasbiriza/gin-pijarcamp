package main

import (
	"latihan_git/src/config"
	"latihan_git/src/helpers"
	"latihan_git/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helpers.Migration()
	defer config.DB.Close()
	routes.Router()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// routes.Router()
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
