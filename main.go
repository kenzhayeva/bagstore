package main

import (
	"bagstore/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Run(":8081")
}
