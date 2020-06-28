package analyse

import (
	"fmt"
	"github.com/valyala/fastjson"
	"log"
)

func main1() {
	var p fastjson.Parser
	v, err := p.Parse(`{
                "str": "bar",
                "int": 123,
                "float": 1.23,
                "bool": true,
                "arr": [1, "foo", {}]
        }`)
	if err != nil {
		log.Fatal(err)
	}

	//查
	fmt.Printf("foo=%s\n", v.GetStringBytes("str"))
	fmt.Printf("int=%d\n", v.GetInt("int"))
	fmt.Printf("float=%f\n", v.GetFloat64("float"))
	fmt.Printf("bool=%v\n", v.GetBool("bool"))
	fmt.Printf("arr.1=%s\n", v.GetStringBytes("arr", "1"))
	//改、增
	v.Set("hello",fastjson.MustParse(`"world"`))
	v.Set("int",fastjson.MustParse(`"456"`))
	//删
	v.Del("float")

	//增加复合数据
	v.Set("complex",fastjson.MustParse(`{"x":"y"}`))
    v.Get("complex").Set("t",fastjson.MustParse(`["x","y"]`))
	v.Get("complex","t").SetArrayItem(2,fastjson.MustParse(`"z"`))

	fmt.Printf("%s\n", v)
}
