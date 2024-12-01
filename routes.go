package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	router.GET("/api/verve/accept", handleAccept)
}
