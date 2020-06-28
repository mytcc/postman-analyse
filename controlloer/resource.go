package controlloer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	user, response := baseLogin(c)
	if response!=nil{
		c.JSON(http.StatusOK,response)
	}else{
		if component.Env=="Dev" {
			//存储用户缓存，获取 token
			//token:=agent.SaveUserToRedis(&user)
			token:=agent.GetJwtToken(user.Id)
			//登录成功
			c.JSON(http.StatusOK,Success(gin.H{
				"token":token,
			}))
		}else {
			c.JSON(http.StatusOK,Success(user))
		}

	}
}
