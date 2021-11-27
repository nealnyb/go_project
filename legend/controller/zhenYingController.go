package controller

import (
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ZhenYingController struct {
}

func (zyc *ZhenYingController) Rotuer(engine *gin.Engine){
	engine.GET("/api/zhenying/query",zyc.ZhenYingQuery)
}

func (zyc *ZhenYingController) ZhenYingQuery(ctx *gin.Context) {
	//调用service
	zhenYingService := service.ZhenYingService{}
	zhenYings,err := zhenYingService.ZhenYingQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
	}
	tool.Success(ctx,zhenYings)
}