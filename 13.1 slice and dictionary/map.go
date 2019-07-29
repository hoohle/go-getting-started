package main

import (
	"fmt"
)

func main() {
	var x = map[string]string{
		"{\"citycode\"": `"101010100"}`,
		`{"PM25"`:       `"31"}`,
		`{"time"`:       ` "2019072318"}`,
		`{"PM10"`:       `"25"}`,
	}
	fmt.Println(x)
	for key, val := range x {
		// fmt.Println("Key", key, "Value:", val)
		fmt.Printf("%v:%v\n", key, val)
	}
	for val := range x {
		fmt.Println("Value:", val)
	}

	for _, val := range x {
		fmt.Println("Value:", val)
	}

	for key, _ := range x {
		fmt.Println("Key:", key)
	}

	//制作一个字典
	y := make(map[string]int)
	y["A"] = 0
	y["B"] = 20
	fmt.Println(y["C"])               //此处值本不应存在
	if val, ok := x["citycode"]; ok { //可以看到x["C"]的值有两个，一个是值，另一个是是否存在此值的bool型变量
		fmt.Println(val)
	} else {
		fmt.Println("查询的值不存在")
	}

}
