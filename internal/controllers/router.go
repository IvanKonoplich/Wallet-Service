package controllers

import "github.com/gin-gonic/gin"

func (c *Controller) InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/test", c.test)
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
		report.GET("/", c.getMonthReport)
	}

	operationsJournalMux := router.Group("/operationsJournal")
	{
		operationsJournalMux.GET("/amount", c.getOperationsListByAmount)
		operationsJournalMux.GET("/date", c.getOperationsListByDate)
	}

	return router
}
