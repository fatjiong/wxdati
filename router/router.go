package router

import (
	"github.com/fatjiong/wxdati/controllers/category"
	"github.com/fatjiong/wxdati/controllers/chapters"
	"github.com/fatjiong/wxdati/controllers/feedback"
	"github.com/fatjiong/wxdati/controllers/novels"
	"github.com/fatjiong/wxdati/controllers/user"
	"github.com/gin-gonic/gin"
)

//初始化路由
func Init(){
	gin := gin.Default()
	//用户登录
	gin.GET("/user/login",user.Login)
	//用户信息
	gin.GET("/userInfo",user.Info)

	//分类
	gin.GET("novels/getCategory",category.List)

	//小说
	gin.GET("novels",novels.List)
	gin.GET("/novels/get",novels.Detail)

	//章节
	gin.GET("/novels/getChapters",chapters.List)
	gin.GET("/novels/getChaptersContent",chapters.Detail)

	//搜索
	gin.GET("novels/search",category.List)
	//热搜关键字
	gin.GET("novels/getSearchKeywords",category.List)

	//反馈
	gin.POST("feedback/send",feedback.Send)

	gin.Run(":8888")
}