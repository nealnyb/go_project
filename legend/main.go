package main

import (
	"legend/controller"
	"legend/middleWare"
	"legend/tool"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg,err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_,err = tool.OrmEngine(cfg)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	
	app := gin.Default()

	registerRouter(app)

	//设置全局跨域访问
	app.Use(middleWare.Cors())

	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}

//路由设置
func registerRouter(router *gin.Engine){
	new(controller.HelloController).Router(router)
	new(controller.ZodiacController).Router(router)
	new(controller.HomeController).Router(router)
	new(controller.HongHuangController).Router(router)
	new(controller.ShengYuController).Router(router)
	new(controller.ShenLongController).Router(router)
	new(controller.ZhenYingController).Rotuer(router)
	new(controller.HumanoidController).Router(router)
}

