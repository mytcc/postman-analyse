package analyse

import (
	"github.com/valyala/fastjson"
	"strconv"
)



func getExample(source *fastjson.Value) *fastjson.Value  {
	var parser fastjson.Parser
	if example,err := parser.Parse(`{"data":[]}`);err==nil{
	setExampleInfo(source.GetArray("item"),example)
 		return example
	}else{
		panic(err)
	}
}

func setExampleInfo(source []*fastjson.Value, data *fastjson.Value)  {
	for _,item:=range source{
		if item.Exists("response"){
			data.Get("data").SetArrayItem(len(data.GetArray("data")),fastjson.MustParse(`{}`))
			dateItem := data.Get("data",strconv.Itoa(len(data.GetArray("data"))-1))
			dateItem.Set("item_id",item.Get("_postman_id"))
			dateItem.Set("example",fastjson.MustParse(`[]`))
			for i,respItem:=range item.GetArray("response"){
				dateItem.Get("example").SetArrayItem(i,fastjson.MustParse(`{}`))
				exampleItem:= dateItem.Get("example",strconv.Itoa(i))
				exampleItem.Set("name",respItem.Get("name"))
				exampleItem.Set("response",fastjson.MustParse(`{}`))
				exampleItem.Set("request",item.Get("request"))
				response:=exampleItem.Get("response")
				response.Set("_postman_previewlanguage",respItem.Get("_postman_previewlanguage"))
				response.Set("body",respItem.Get("body"))
				response.Set("code",respItem.Get("code"))
				response.Set("header",respItem.Get("header"))
				response.Set("status",respItem.Get("status"))
				//http:=item.Get("request","method").String()
				url:=item.Get("request","url").String()

				exampleItem.Set("snippet",fastjson.MustParse(`{"http":`+url+`}`))
			}

		}
		if item.Exists("item"){
			setExampleInfo(item.GetArray("item"), data)
		}
	}

}
