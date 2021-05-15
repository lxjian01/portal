package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context){
	c.String(http.StatusOK, "Hello %s", "name")
}
