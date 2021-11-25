package main

import (
	"legend/controller"
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

	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}

//路由设置
func registerRouter(router *gin.Engine){
	new(controller.HelloController).Router(router)
	new(controller.ZodiacController).Router(router)
	new(controller.HomeController).Router(router)
	new(controller.HongHuangController).Router(router)
}