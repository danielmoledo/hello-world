package controllers

import (
	"github.com/danielmoledo/hello-world/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//POST /hello-world
func HelloWorld(c *gin.Context) {

	response, error := services.HelloWorld()

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
