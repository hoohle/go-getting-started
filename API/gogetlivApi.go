package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	// ../prisma-client
)

//此函数用来获得天气实况数据
// type a Hybridization{
// 	beijing	Hybridization
// 	chaoyang	Hybridization
// 	shijiazhaung	Hybridization
// 	handan	Hybridization
// }

type Hybridization struct { //给你
	// ID          bson.ObjectId `json:"_id"`
	Windspeed   string `json:"windspeed"`
	Citycode    string `json:"citycode"`
	Airpressure string `json:"airpressure"`
	Phenomena   string `json:"phenomena"`
	Rdesc       string `json:"rdesc"`
	Humidity    string `json:"humidity"`
	Updatetime  string `json:"updatetime"`
	Windpower   string `json:"windpower"`
	Feelst      string `json:"feelst"`
	Winddirect  string `json:"winddirect"`
	Rcode       int32  `json:"rcode"`
	Rain        string `json:"rain"`
	Temperature string `json:"temperature"`

	Time    string `json:"time"`
	PM25    string `json:"PM25"`
	PM10    string `json:"PM10"`
	SO2     string `json:"SO2"`
	O3      string `json:"o3"`
	NO2     string `json:"NO2"`
	Primary string `json:"primary"`
	CO      string `json:"CO"`
	AQI     string `json:"AQI"`
	// Airrcode string `json:"rcode"`
	// Acitycode string `json:"citycode"`
}

// data:=client.RealTimeData(&prisma.RealTimeDataWhereInput{"windspeed":""})

//实时天气&空气数据获取函数
func WeatherData(baseURL, funcCode, secert, cityCode, function string) Hybridization {
	url := baseURL + funcCode + secert //定义url的组成	此处将其拆成四段可根据需求自行组合

	payload := strings.NewReader(cityCode) //json格式的请求body体
	// json.NewEncoder(payload).Encode(cityCode)
	// playload, error := json.Marshal(cityCode)
	// fmt.Println(playload)
	// fmt.Println(string(playload))
	// fmt.Println(error)
	RealtimeResult := Hybridization{}               //把结构体放进RealtimeWeather变量
	req, _ := http.NewRequest("POST", url, payload) //用了http库里的NewRequest函数功能	参数解释（方法，url地址，？）

	res, err := http.DefaultClient.Do(req) //把请求req的返回值赋给res(ponse)，如果的到了一个错误传给err在下面做处理
	if res != nil {                        //如果请求后返回的是空的应对方法
		defer res.Body.Close()
	}
	if err != nil { //如果有错误值的应对方法
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(res.Body) //ioutil库里的readall方法
	/*func ReadAll

	func ReadAll(r io.Reader) ([]byte, error)

	ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。成功的调用返回的err为nil而非EOF。因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误。*/

	// fmt.Println(res) //查看响应的首部
	fmt.Println(string(body)) //查看body	响应字段

	json.Unmarshal([]byte(body), &RealtimeResult) //已将body头里的值存入到了RealtimeWeather变量
	//json库里的Unmarshal方法    //还有一个相对应的方法						对象转换为JSON的方法（函数）为 json.Marshal()
	/*func Unmarshal ¶
	  func Unmarshal(data []byte, v interface{}) error
	  Unmarshal函数解析json编码的数据并将结果存入v指向的值。

	  Unmarshal和Marshal做相反的操作，必要时申请映射、切片或指针，有如下的附加规则：

	  要将json数据解码写入一个指针，Unmarshal函数首先处理json数据是json字面值null的情况。此时，函数将指针设为nil；否则，函数将json数据解码写入指针指向的值；如果指针本身是nil，函数会先申请一个值并使指针指向它。

	  要将json数据解码写入一个结构体，函数会匹配输入对象的键和Marshal使用的键（结构体字段名或者它的标签指定的键名），优先选择精确的匹配，但也接受大小写不敏感的匹配。

	  要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：

	  Bool                   对应JSON布尔类型
	  float64                对应JSON数字类型
	  string                 对应JSON字符串类型
	  []interface{}          对应JSON数组
	  map[string]interface{} 对应JSON对象
	  nil                    对应JSON的null

	  如果一个JSON值不匹配给出的目标类型，或者如果一个json数字写入目标类型时溢出，Unmarshal函数会跳过该字段并尽量完成其余的解码操作。如果没有出现更加严重的错误，本函数会返回一个描述第一个此类错误的详细信息的UnmarshalTypeError。

	  JSON的null值解码为go的接口、指针、切片时会将它们设为nil，因为null在json里一般表示“不存在”。 解码json的null值到其他go类型时，不会造成任何改变，也不会产生错误。

	  当解码字符串时，不合法的utf-8或utf-16代理（字符）对不视为错误，而是将非法字符替换为unicode字符U+FFFD。
	  Example*/
	switch function {
	case "打印":
		fmt.Println(fmt.Sprintf("已经存入到realtimeresult中的值：%+v\n", RealtimeResult))
		fmt.Println(fmt.Sprintf("%+v\n", cityCode))
	case "存取":
		fmt.Println("开始操作prisma存入")
	default:
		fmt.Println("请输入打印（只打印）或者存取（仅存取）")
	}
	// }
	// fmt.Println(fmt.Sprintf("%+v\n", RealtimeWeather))
	// fmt.Sprintf("%+v\n", RealtimeWeather)
	/*func Println

	func Println(a ...interface{}) (n int, err error)

	Println采用默认格式将其参数格式化并写入标准输出。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。返回写入的字节数和遇到的任何错误。
	*/

	/*func Sprintf

			func Sprintf(format string, a ...interface{}) string

			Sprintf根据format参数生成格式化的字符串并返回该字符串。

	通用：
	%v	值的默认格式表示
	%+v	类似%v，但输出结构体时会添加字段名
	%#v	值的Go语法表示
	%T	值的类型的Go语法表示
	%%	百分号

	布尔值：

	%t	单词true或false

	整数：

	%b	表示为二进制
	%c	该值对应的unicode码值
	%d	表示为十进制
	%o	表示为八进制
	%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	%x	表示为十六进制，使用a-f
	%X	表示为十六进制，使用A-F
	%U	表示为Unicode格式：U+1234，等价于"U+%04X"

	浮点数与复数的两个组分：

	%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	%e	科学计数法，如-1234.456e+78
	%E	科学计数法，如-1234.456E+78
	%f	有小数部分但无指数部分，如123.456
	%F	等价于%f
	%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	字符串和[]byte：

	%s	直接输出字符串或者[]byte
	%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	%x	每个字节用两字符十六进制数表示（使用a-f）
	%X	每个字节用两字符十六进制数表示（使用A-F）

	指针：

	%p	表示为十六进制，并加上前导的0x

	没有%u。整数如果是无符号类型自然输出也是无符号的。类似的，也没有必要指定操作数的尺寸（int8，int64）。

	宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：

	%f:    默认宽度，默认精度
	%9f    宽度9，默认精度
	%.2f   默认宽度，精度2
	%9.2f  宽度9，精度2
	%9.f   宽度9，精度0    */

	return RealtimeResult
	// return WeatherLive{}
	// if	funcCode=="cityairlive"
	// var

}
