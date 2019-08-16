package api

import "github.com/gin-gonic/gin"

func GetAccountDetail(c *gin.Context){
	c.JSON(200, gin.H{
		"number": 5732110628,
		"balance":10000.00,
	})
}

func PostTransfer(c *gin.Context){
	c.JSON(200, gin.H{
		"transaction_id": 1,
		"fee":0,
		"status":false,
	})
}

func PostConfirmTransfer(c *gin.Context){
	c.JSON(200, gin.H{
		"transaction_id": 1,
		"source_account_number":5732110628,
		"destination_account_number":4351223508,
		"amount":8500.00,
		"fee":0,
		"status":true,
	})
}