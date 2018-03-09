package router

import (
	"github.com/fatjiong/wxdati/controllers/user"
	"github.com/gin-gonic/gin"
)

//初始化路由
func Init(){
	gin := gin.Default()
	//用户
	gin.GET("/login",user.Login)
	gin.GET("/userInfo",user.Info)
	gin.Run(":8888")
}