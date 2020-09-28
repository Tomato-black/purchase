package contorller

import (
	"github.com/gin-gonic/gin"
	"purchase/src/bean/vo"
)

func SupplierRouter(router *gin.Engine) {
	supplier := router.Group("/api/supplier")
	{
		supplier.GET("/save", startPage)
	}
}

func startPage(c *gin.Context) {
	var supplierVo vo.SupplierVo
	// If `GET`, only `Form` binding engine (`query`) used.
	// 如果是Get，那么接收不到请求中的Post的数据？？
	// 如果是Post, 首先判断 `content-type` 的类型 `JSON` or `XML`, 然后使用对应的绑定器获取数据.
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&supplierVo) == nil {
		supplier := supplierVo.VoToModel()
		supplier.Save()
	}

	c.String(200, "Success")
}
