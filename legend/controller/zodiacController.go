package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ZodiacController struct {
}

func (zc *ZodiacController) Router(engine *gin.Engine) {
	engine.GET("/api/zodiac/add",zc.ZodiacAdd)
	engine.GET("/api/zodiac/query",zc.ZodiacQuery)
	engine.PUT("/api/zodiac/refresh/update",zc.ZodiacRefreshUpdate)
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

func (zc *ZodiacController) ZodiacRefreshUpdate(ctx *gin.Context) {
	var data model.Zodiac
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	zodiacService := &service.ZodiacService{}
	res,err := zodiacService.ZodiacRefreshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"更新数据失败")
	}
	tool.Success(ctx,res)
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