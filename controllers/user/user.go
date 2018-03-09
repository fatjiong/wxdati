package user

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/fatjiong/wxdati/controllers/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
)

func Login(c *gin.Context){
	error := common.Error
	code := c.DefaultQuery("code","")
	if code == ""{
		error("code不得为空",c)
		return
	}

	appID := ""
	secret := ""
	codeAskUrl :="https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&js_code={code}&grant_type=authorization_code"
	codeAskUrl  = strings.Replace(codeAskUrl, "{appid}",  appID,  -1)
	codeAskUrl  = strings.Replace(codeAskUrl, "{secret}", secret, -1)
	codeAskUrl  = strings.Replace(codeAskUrl, "{code}",   code,   -1)

	resp,err := http.Get(codeAskUrl)
	if err != nil {
		fmt.Println(err.Error())
		error("error",c)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		error("error",c)
		return
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {error("error",c)
		fmt.Println(err.Error())
		error("error",c)
		return
	}

	if _,ok :=data["session_key"];!ok{
		fmt.Println("session_key 不存在")
		fmt.Println(data)
		error("error",c)
		return
	}

	//var openid string
	//var token string
	//openid  = data["openid"].(string)
	//token = data["session_key"].(string)


	fmt.Println(data)

	c.JSON(200,gin.H{
		"code":http.StatusOK,
		"message":"登录成功",
		"openid":data["openid"].(string),
	})

}



func Info(c *gin.Context){

}