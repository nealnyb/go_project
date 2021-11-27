package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type HumanoidController struct {
}

func (huc *HumanoidController) Router(engine *gin.Engine) {
	engine.GET("/api/humanoid/query",huc.HumanoidQuery)
	engine.PUT("/api/humanoid/refresh/update",huc.HumanoidRefreshUpdate)
}

func (huc *HumanoidController) HumanoidQuery(ctx *gin.Context){
	//调用service
	humanoidService := service.HumanoidService{}
	humanoids,err := humanoidService.HumanoidQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
	}
	tool.Success(ctx,humanoids)
}

func (huc *HumanoidController) HumanoidRefreshUpdate(ctx *gin.Context) {
	var data model.Humanoid
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	humanoidService := service.HumanoidService{}
	res,err := humanoidService.HumanoidRefreshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"操作数据库失败")
	}
	tool.Success(ctx,res)
}