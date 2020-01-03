package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func urlMappings() {
	router.GET("/hey", hey)
}


func hey(c *gin.Context) {
	c.JSON(http.StatusOK, "hey")
}