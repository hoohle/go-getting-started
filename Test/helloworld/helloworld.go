package main

/*
文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
文件夹名与包名没有直接关系，并非需要一致。
同一个文件夹下的文件只能有一个包名，否则编译报错。
*/
import (
	"encoding/json"
	"fmt"

	mathClass "../myMath"
)

var postbody string

type Postbody struct {
	accesskey string
	citycodes string
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(mathClass.Add(1, 1))
	fmt.Println(mathClass.Sub(1, 1))
	fmt.Println(mathClass.Sub2(6, 6))
	postbody = "{\"accesskey\":\"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=\",\"citycodes\":[ \"101010100\", \"101010300\" ,\"101090101\",\"101091001\"]}"
	fmt.Printf("json格式体是%+v\n", postbody)
	postbody2 := `{"accesskey":"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=","citycodes":[ "101010100", "101010300" ,"101090101","101091001"]}`
	fmt.Printf("第二种格式是%+v\n", postbody2)

	c := Postbody{"secert", `"101010300","101010100"`}
	v, error := json.Marshal(c)
	fmt.Println(string(v))
	fmt.Println(error)
	fmt.Println("-------------------------------")
	//fmt.Println("{ "accesskey":"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=", "citycodes":[   "101190101",   "101010100"  ]}")
	/*{
	  "accesskey":"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=",
	  "citycodes":[
	    "101010100"
	  ]
	}*/

	//{"accesskey":"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=","citycodes":["101190101","101010100","101010300"]}

}
