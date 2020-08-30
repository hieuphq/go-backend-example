package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.Run(":8080")
}
