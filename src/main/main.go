package main

import (
	"github.com/gin-gonic/gin"
	"purchase/src/contorller"
	)

func main()  {
  router := gin.Default()
  contorller.SupplierRouter(router)
  router.Run(":8080")
}
