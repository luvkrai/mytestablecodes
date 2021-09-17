package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/luvkrai/mytestablecodes/services"
	
)

var (
	PingController = pingController{}
)

type pingController struct{}

func (pc *pingController) Ping(c *gin.Context) {
	pService := services.PingService
	result, err := pService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, result)
}