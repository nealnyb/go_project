package controller

import (
	"legend/service"

	"github.com/gin-gonic/gin"
)

type ZodiacController struct {
}

func (zc *ZodiacController) Router(engine *gin.Engine) {
	engine.GET("/api/zodiac",zc.ZodiacAdd)
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