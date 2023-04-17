package controllers

import "github.com/gin-gonic/gin"

func (c *Controller) InitRouter() *gin.Engine {
	router := gin.Default()
	balance := router.Group("/balance")
	{
		balance.GET("/:id", c.getBalance)
		balance.POST("/", c.balanceIncrease)
		balance.POST("/move-funds", c.transferOfFunds)
	}

	reserve := router.Group("/reserve")
	{
		reserve.POST("/", c.reserveFunds)
		reserve.POST("/revenue-confirm", c.revenueApproval)
		reserve.POST("/revenue-deny", c.revenueDeny)
	}

	report := router.Group("/report")
	{
		report.POST("/", c.getMonthReport)
	}

	operationsJournalMux := router.Group("/operationsJournal")
	{
		operationsJournalMux.POST("/amount", c.getOperationsListByAmount)
		operationsJournalMux.POST("/date", c.getOperationsListByDate)
	}

	return router
}
