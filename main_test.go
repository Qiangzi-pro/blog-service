package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestMainFunc(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run()
}
