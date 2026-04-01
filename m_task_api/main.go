package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/calculate", calculate)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}