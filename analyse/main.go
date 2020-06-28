package analyse

import (
	"github.com/valyala/fastjson"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var collectionJson *fastjson.Value
var exampleJson *fastjson.Value
var Profile="dev"

func GetCollectionJson() *fastjson.Value  {
	if collectionJson==nil{
		log.Println("初始化Collection")
		initResource()
	}else{
		log.Println("无需初始化Collection")
	}
	if collectionJson==nil{
		//如果仍然为空，则异常
		panic("无法初始化 Collection")
	}
	return collectionJson
}
func GetExampleJson() *fastjson.Value  {
	if exampleJson==nil{
		log.Println("初始化Example")
		initResource()
	}else{
		log.Println("无需初始化Example")
	}
	if exampleJson==nil{
		//如果仍然为空，则异常
		panic("无法初始化 Collection")
	}
	return exampleJson
}
func initResource() {
	//读取 environment 源文件
	var environmentSource *fastjson.Value
	envFileName:="postman_environment.json"
	dataFileName :="postman_collection.json"
	if Profile=="prod"{
		envFileName=GetCurrentDirectory()+"/"+envFileName
		dataFileName =GetCurrentDirectory()+"/"+ dataFileName
	}

	if file, err := os.Open(envFileName);err==nil{
		defer file.Close()
		if source,err:= ioutil.ReadAll(file);err==nil{
			var parser fastjson.Parser
			if environmentSource, err= parser.ParseBytes(source);err!=nil{
				panic(err)
			}
		}else{
			panic(err)
		}
	}else{
		panic(err)
	}

	//读取 collection 源文件
	var collectionSource *fastjson.Value
	if file, err := os.Open(dataFileName);err==nil{
		defer file.Close()
		if source,err:= ioutil.ReadAll(file);err==nil{
			sourceText:=string(source)
			//设置环境变量
			sourceText= processEnv(sourceText,environmentSource)
			var parser fastjson.Parser
			if collectionSource, err= parser.Parse(sourceText);err!=nil{
				panic(err)
			}
		}else{
			panic(err)
		}
	}else{
		panic(err)
	}








	//获取Collection
	collection:= getCollection(collectionSource,environmentSource)
	//获取Example
	example:= getExample(collection)

	//输出至文件-Collection
	//collectionFileName:="collection-ruoli.json"
	//var collectionFile *os.File
	//defer collectionFile.Close()
	//if _,err := os.Stat(collectionFileName);err!=nil{
	//	//文件不存在
	//	collectionFile,_ = os.Create(collectionFileName)
	//}else{
	//	collectionFile, _ = os.OpenFile(collectionFileName, os.O_WRONLY|os.O_TRUNC, 0600)
	//}
	//collectionFile.Write([]byte(collection.String()))
	collectionJson=collection

	//输出至文件-Example
	//exampleFileName:="example-ruoli.json"
	//var exampleFile *os.File
	//defer exampleFile.Close()
	//if _,err := os.Stat(exampleFileName);err!=nil{
	//	//文件不存在
	//	exampleFile,_ = os.Create(exampleFileName)
	//}else{
	//	exampleFile, _ = os.OpenFile(exampleFileName, os.O_WRONLY|os.O_TRUNC, 0600)
	//}
	////fmt.Println(example)
	//exampleFile.Write([]byte(example.String()))
	exampleJson=example



}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}