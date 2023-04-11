package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type reportRequest struct {
	Month string `json:"month"`
}

func (co *Controller) getMonthReport(c *gin.Context) {
	var input reportRequest
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	result, err := co.uc.GetMonthReport(input.Month)
	if err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, Response{result})
}
