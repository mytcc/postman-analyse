
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"postman-analyse/analyse"
	"strconv"
	"strings"
	"time"
)


func main() {


	r:=gin.New()
	r.Static("/view", "./resource")
	r.Static("/images", "./resource/images")
	r.Static("/js", "./resource/js")
	r.Static("/settings", "./resource/settings")
	r.Static("/styles", "./resource/styles")
    //设置环境
    //analyse.Profile="prod"
	//处理Json请求
	r.POST("/api/examples/2742622/SzzkdxWs.json", func(context *gin.Context) {
		//判断是否已经解析过
		result:=analyse.GetExampleJson()
		jsonData := []byte(result.String())
		var v interface{}
		json.Unmarshal(jsonData, &v)
		data := v.(map[string]interface{})
		context.JSON(http.StatusOK,data)
	})
	r.GET("/api/collections/2742622/SzzkdxWs.json", func(context *gin.Context) {
		//判断是否已经解析过
		result:=analyse.GetCollectionJson();
		jsonData := []byte(result.String())
		var v interface{}
		json.Unmarshal(jsonData, &v)
		data := v.(map[string]interface{})
		context.JSON(http.StatusOK,data)
	})
	r.GET("/scss/custom.css", func(context *gin.Context) {
		context.Request.Response.Header.Add("Content-Type","text/css; charset=UTF-8")
		r.HandleContext(context)
	})

	//r.POST("/login",controller.Login)//登陆
	//r.POST("/register",controller.Register)//注册
	//r.GET("logout",controller.Logout)
	//auth:=r.Use(middleware.TokenAuth)
	ports:=80
	for{
		if err:=r.Run(":"+strconv.Itoa(ports));err!=nil{
			error:=string(err.Error())
			if strings.Contains(error,"bind: Only one usage of each socket address"){
				ports++
			}else{
				fmt.Println("发生错误："+err.Error())
				break
			}
		}
	}
	time.Sleep(time.Duration(5)*time.Minute)

}