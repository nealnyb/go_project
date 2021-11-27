package controller

import (
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type HumanoidController struct {
}

func (huc *HumanoidController) Router(engine *gin.Engine) {
	engine.GET("/api/humanoid/query",huc.HumanoidQuery)
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