package main

import (
	db "mwdowns/rest-api/DB"
	"mwdowns/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
