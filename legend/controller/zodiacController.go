package controller

import (
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ZodiacController struct {
}

func (zc *ZodiacController) Router(engine *gin.Engine) {
	engine.GET("/api/zodiac/add",zc.ZodiacAdd)
	engine.GET("/api/zodiac/query",zc.ZodiacQuery)
}



func (zc *ZodiacController) ZodiacQuery(ctx *gin.Context){
	//调用service
	zodiacService := &service.ZodiacService{}
	zodiacs,err := zodiacService.ZodiacQuery()
	if err != nil {
		tool.Failed(ctx,"数据获取失败")
		return
	}
	tool.Success(ctx,zodiacs)
}

func (zc *ZodiacController) ZodiacAdd(context *gin.Context){
	zs := service.ZodiacService{}
	res := zs.ZodiacAdd()
	if res {
		context.JSON(200,map[string]interface{}{
			"code":1,
			"msg":"新增成功",
		})
		return
	}

	context.JSON(200,map[string]interface{}{
		"code":0,
		"msg":"新增失败",
	})

}