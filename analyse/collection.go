package analyse

import (
	uuid "github.com/satori/go.uuid"
	"github.com/valyala/fastjson"
	"strings"
)

func getCollection(collectionSource *fastjson.Value,environmentSource *fastjson.Value) *fastjson.Value  {
    collection :=collectionSource
	info:=collection.Get("info")
	info.Set("toc",fastjson.MustParse(`[]`))
	info.Set("owner",fastjson.MustParse(`"2742622"`))
	info.Set("collectionId",collection.Get("info","_postman_id"))
	info.Set("publishedId",fastjson.MustParse(`"SzzkdxWs"`))
	info.Set("public",fastjson.MustParse(`true`))
	info.Set("customColor",fastjson.MustParse(`{"highlight":"EF5B25","right-sidebar":"303030","top-bar":"FFFFFF"}`))
    //格式化 info-description
	info.Set("description",fastjson.MustParse(`"<html><head></head><body><p>`+strings.Trim(info.Get("description").String(),"\"")+`</p>\n</body></html>"`))
	//设置 schema
	info.Set("schema",fastjson.MustParse(`"https://schema.getpostman.com/json/collection/v2.0.0/collection.json"`))
	//设置 customColor
	topbar:=info.Get("customColor","top-bar")
	sidebar:=info.Get("customColor","right-sidebar")
	highlight:=info.Get("customColor","highlight")
	info.Del("customColor")
	info.Set("customColor",fastjson.MustParse(`{}`))
	info.Get("customColor").Set("top-bar",topbar)
	info.Get("customColor").Set("right-sidebar",sidebar)
	info.Get("customColor").Set("highlight",highlight)

	if collection.Exists("item"){
		//设置postmanId
		setCollectionItemInfo(collection.GetArray("item"))
		//处理 request-url
		setRequestInfo(collection.GetArray("item"))
		//处理 response-url
		setResponseInfo(collection.GetArray("item"))
	}

	return collection
}

func setCollectionItemInfo(items []*fastjson.Value)   {
	for _,item:=range items{
		uid:=uuid.NewV4()
		item.Set("_postman_id",fastjson.MustParse(`"`+uid.String()+`"`))
		if !item.Exists("description"){
			item.Set("description",fastjson.MustParse(`""`))
		}
		if item.Exists("item"){
			setCollectionItemInfo(item.GetArray("item"))
		}
	}
}
func setRequestInfo(items []*fastjson.Value)   {
	for _,item:=range items{
		if item.Exists("request"){
			requestItem:=item.Get("request")
			url:=requestItem.Get("url")
			host:=url.Get("host")
			path:=url.Get("path")
			raw:=url.Get("raw")
			requestItem.Set("urlObject",fastjson.MustParse(`{}`))
			urlObject:=item.Get("request","urlObject")
			urlObject.Set("path",path)
			urlObject.Set("host",host)
			if url.Exists("query"){
				urlObject.Set("query",url.Get("query"))
				for _,queryItem:=range urlObject.GetArray("query"){
					description:=fastjson.MustParse(`{"content":"<p>`+strings.Trim(queryItem.Get("description").String(),"\"")+`</p>\n","type":"text/plain"}`)
					//调整顺序
					key:=queryItem.Get("key")
					value:=queryItem.Get("value")
					queryItem.Del("description")
					queryItem.Del("key")
					queryItem.Del("value")
					queryItem.Set("description",description)
					queryItem.Set("key",key)
					queryItem.Set("value",value)
				}

			}else{
				urlObject.Set("query",fastjson.MustParse(`[]`))
			}
			if url.Exists("variable"){
				urlObject.Set("variable",url.Get("variable"))
			}else{
				urlObject.Set("variable",fastjson.MustParse(`[]`))
			}
			requestItem.Set("url",raw)
			//设置 description
			if(requestItem.Exists("description")){
				requestItem.Set("description",fastjson.MustParse(`"<p>`+strings.Trim(requestItem.Get("description").String(),"\"")+`</p>\n"`))
			}else{
				requestItem.Set("description",fastjson.MustParse(`"<p>暂无描述</p>\n"`))
			}

		}
		if item.Exists("item"){
			setRequestInfo(item.GetArray("item"))
		}
	}
}

func setResponseInfo(items []*fastjson.Value)   {
	for _,item:=range items{
		if item.Exists("response"){
			for _,response:=range item.GetArray("response"){
				url:=response.Get("originalRequest","url")
				raw:=url.Get("raw")
				response.Get("originalRequest").Set("url",raw)
				body:=response.Get("body")
				response.Del("body")
				response.Set("responseTime",fastjson.MustParse(`null`))
				response.Set("body",body)
			}
		}
		if item.Exists("item"){
			setResponseInfo(item.GetArray("item"))
		}
	}
}


func processEnv(textSource string,environmentSource *fastjson.Value) string  {
	//解析Env
	if environmentSource.Exists("values"){
		for _,env:=range environmentSource.GetArray("values"){
			if strings.Compare(env.Get("enabled").String(),"true")==0{
				key:="{{"+strings.Trim(env.Get("key").String(),"\"")+"}}"
				value:=strings.Trim(env.Get("value").String(),"\"")
				textSource=strings.Replace(textSource,key,value,-1)
			}
		}
	}
	return textSource;
}
