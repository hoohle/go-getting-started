package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//一下是废弃的url改为爬虫爬
// var weatherhisurlArr = []string{
// 	"http://service.envicloud.cn:8082/v2/weatherhistory/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/”+ strconv.Itoa()+“/20160909/01",
// 	"http://service.envicloud.cn:8082/v2/weatherhistory/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101071201/20160909/01",
// 	"http://service.envicloud.cn:8082/v2/weatherhistory/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090101/20160909/01",
// 	"http://service.envicloud.cn:8082/v2/weatherhistory/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010800/20160909/01",
// 	"http://service.envicloud.cn:8082/v2/weatherhistory/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090314/20160909/01",
// 	"http://service.envicloud.cn:8082/v2/weatherhistory/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101091001/20160909/01",
// }

// var airhisurlArr = []string{
// 	"http://service.envicloud.cn:8082/v2/cityhourlyair/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010100/2016090901",
// 	"http://service.envicloud.cn:8082/v2/cityhourlyair/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101071201/2016090901",
// 	"http://service.envicloud.cn:8082/v2/cityhourlyair/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090101/2016090901",
// 	"http://service.envicloud.cn:8082/v2/cityhourlyair/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010800/2016090901",
// 	"http://service.envicloud.cn:8082/v2/cityhourlyair/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090314/2016090901",
// 	"http://service.envicloud.cn:8082/v2/cityhourlyair/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101091001/2016090901",
// }

//以下是测试码
// var start ,end	int
// start:
// for
// type varible struct {
// 	addyear   string
// 	addtime   string
// 	exampleid string
// }

// var A = varible{"xxx", "xxx", "101010100"}
var weatherurlArr = []string{ //天气实况的REST API 接口
	"http://service.envicloud.cn:8082/v2/weatherlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010100", //北京
	"http://service.envicloud.cn:8082/v2/weatherlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010300", //朝阳
	"http://service.envicloud.cn:8082/v2/weatherlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010200", //海淀
	"http://service.envicloud.cn:8082/v2/weatherlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010800", //延庆
	"http://service.envicloud.cn:8082/v2/weatherlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090314", //崇礼
	"http://service.envicloud.cn:8082/v2/weatherlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101091001", //邯郸
}
var airurlArr = []string{ //空气实况的REST API 接口
	"http://service.envicloud.cn:8082/v2/cityairlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010100", //北京
	"http://service.envicloud.cn:8082/v2/cityairlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101071201", //朝阳
	"http://service.envicloud.cn:8082/v2/cityairlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090101", //石家庄
	"http://service.envicloud.cn:8082/v2/cityairlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101010800",
	"http://service.envicloud.cn:8082/v2/cityairlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101090314",
	"http://service.envicloud.cn:8082/v2/cityairlive/Y2RPBZEXOTE1NJI2NDMXMJA5MTY=/101091001", //邯郸
}

// type WeatherInfoJson struct {
// 	Weatherinfo WeatherInfoObject
// }
// type AirInfoJson struct {
// 	AirInfo AirInfoObject
// }

type weatherInfoObject struct {
	Citycode    string `json:"citycode"`
	Rcode       int    `json:"rcode"`
	Rdesc       string `json:"rdesc"`
	Updatetime  string `json:"updatetime"`
	Phenomena   string `json:"phenomena"`
	Temperature string `json:"temperature"`
	Feelst      string `json:"feelst"`
	Airpressure string `json:"airpressure"`
	Humidity    string `json:"humidity"`
	Rain        string `json:"rain"`
	Winddirect  string `json:"winddirect"`
	Windpower   string `json:"windpower"`
	Windspeed   string `json:"windspeed"`
}
type airInfoObject struct {
	Citycode string `json:"citycode"`
	PM25     string `json:"PM25"`
	Time     string `json:"time"`
	Rdesc    string `json:"rdesc"`
	PM10     string `json:"PM10"`
	SO2      string `json:"SO2"`
	O3       string `json:"o3"`
	NO2      string `json:"NO2"`
	Primary  string `json:"primary"`
	Rcode    int    `json:"rcode"`
	CO       string `json:"CO"`
	AQI      string `json:"AQI"`
}

// type WeatherLive struct {
// 	ID          bson.ObjectId `json:"_id"`
// 	Windspeed   string        `json:"windspeed"`
// 	// Citycode    string        `json:"citycode"`
// 	airpressure string        `json:"airpressure"`
// 	phenomena   string        `json:"phenomena"`
// 	rdesc       string        `json:"rdesc"`
// 	humidity    string        `json:"humidity"`
// 	updatetime  string        `json:"updatatime"`
// 	windpower   string        `json:"windpower"`
// 	feelst      string        `json:"feelst"`
// 	winddirect  string        `json:"winddirect"`
// 	rcode       int           `json:"rcode"`
// 	rain        string        `json:"rain"`
// 	temperature string        `json:"temperature"`
// }

func main() {
	for {
		for _, data := range weatherurlArr {
			url := data
			s := weatherInfoObject{}
			req, _ := http.NewRequest("GET", url, nil)

			res, err := http.DefaultClient.Do(req)
			if res != nil {
				defer res.Body.Close()
			}
			if err != nil {
				log.Fatal(err)
			}

			body, _ := ioutil.ReadAll(res.Body)

			json.Unmarshal([]byte(body), &s)

			fmt.Println(fmt.Sprintf("%+v", s))

		}
		// time.Sleep(1 * time.Second)
		//以上为天气的json数据获取

		for _, data := range airurlArr {
			url := data
			s := airInfoObject{}
			req, _ := http.NewRequest("GET", url, nil)

			res, err := http.DefaultClient.Do(req)
			if res != nil {
				defer res.Body.Close()
			}
			if err != nil {
				log.Fatal(err)
			}

			// input,err:=ioutil.ReadAll(resp.Body)
			body, _ := ioutil.ReadAll(res.Body)

			json.Unmarshal([]byte(body), &s)
			// var jsonWeather WeatherInfoJson
			// json.Unmarshal(body, &jsonWeather)
			// log.Printf("Results:%v\n", jsonWeather)

			// log.Println(jsonWeather.Weatherinfo.citycode)
			// log.Println(jsonWeather.Weatherinfo.PM25)

			// fmt.Println(fmt.Sprintf("%+v", jsonWeather))
			// fmt.Println(res)
			fmt.Println(fmt.Sprintf("%+v", s))
			// fmt.Println(string(body))
		}
		//以上与获取天气json数据相同的操作，获取空气json数据

		time.Sleep(20 * time.Minute) //定时器20分钟一次
	}
}
