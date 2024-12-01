package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleAccept(c *gin.Context) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "invalid id"})
		return
	}

	endpoint := c.Query("endpoint")

	if !addUniqueID(id) {
		c.JSON(http.StatusOK, gin.H{"status": "failure"})
		return
	}

	if endpoint != "" {
		count := getUniqueCount()
		go sendHTTPRequest(endpoint, count)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
