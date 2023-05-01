package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "SELAM!!")
	})
	route.Run(":5555")
}
