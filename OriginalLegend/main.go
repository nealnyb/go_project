package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Handle("GET","/hello",func(c *gin.Context) {
		name := c.DefaultQuery("name","hello")
		fmt.Println(name)
		c.Writer.Write([]byte("hello," + name))
	})
	err := engine.Run(":16950")
	if err != nil {
		log.Fatal(err.Error())
	}
}