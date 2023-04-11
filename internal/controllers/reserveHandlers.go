package controllers

import (
	"avitoTest/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (co *Controller) reserveFunds(c *gin.Context) {
	var input entities.Order
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	if err := co.uc.ReserveFunds(input); err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusCreated)
}

func (co *Controller) revenueApproval(c *gin.Context) {
	var input entities.Order
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	if err := co.uc.RevenueApproval(input); err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusOK)
}

func (co *Controller) revenueDeny(c *gin.Context) {
	var input entities.Order
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	if err := co.uc.RevenueDeny(input); err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusOK)
}
