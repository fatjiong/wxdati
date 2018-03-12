package common

import (
	"github.com/fatjiong/wxdati/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendErrJSON 有错误发生时，发送错误JSON
func Error(msg string,c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code" : model.ErrorCode.ERROR,
		"message"   : msg,
		"data"  : gin.H{},
	})
}

//成功方法
func Success(msg string,data *gin.H,c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":model.ErrorCode.SUCCESS,
		"message":msg,
		"data":data,
	})
}