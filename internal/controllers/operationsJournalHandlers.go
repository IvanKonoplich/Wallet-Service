package controllers

import (
	"avitoTest/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (co *Controller) getOperationsListByAmount(c *gin.Context) {
	var input entities.User
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	result, err := co.uc.GetOperationsListByAmount(input)
	if err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, Response{result})
}
func (co *Controller) getOperationsListByDate(c *gin.Context) {
	var input entities.User
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	result, err := co.uc.GetOperationsListByDate(input)
	if err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, Response{result})
}
