package tool

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 0 //操作成功
	FAILED int = 1
)

//普通的成功返回
func Success(ctx *gin.Context,v interface{}) {
	ctx.JSON(http.StatusOK,map[string]interface{}{
		"code":SUCCESS,
		"msg":"成功",
		"data":v,
	})
}

//普通失败的返回
func Failed(ctx *gin.Context,v interface{}){
	ctx.JSON(http.StatusOK,map[string]interface{}{
		"code":FAILED,
		"msg":v,
	})
}