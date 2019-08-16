package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"transfer/api"
)

func main() {
	host := flag.String("host", "3000", "your start host port")
	flag.Parse()

	r := gin.Default()
	r.GET("/account/:account_number", api.GetAccountDetail)
	r.POST("/transfer", api.PostTransfer)
	r.POST("/confirm-transfer", api.PostConfirmTransfer)
	r.Run(":"+*host)
}
