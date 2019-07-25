package main

import (
	"context"
	"fmt"
	"log"
	"time"

	api "./API"
	"./prisma-client"
	"github.com/spf13/viper"
)

var ( //这种因式分解关键字的写法一般用于声明全局变量
	prismaserver, secret string
	a                    = "http://service.envicloud.cn:8082" //baseUrl,
	b                    = "/v2/"
	c                    = "" //secert,
	postbody             = `{"accesskey":"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=","citycodes":[ "101010100", "101010300" ,"101090101","101091001"]}`
	apostbody            = `{"accesskey":"Y2RPBZEXOTE1NJI2NDMXMJA5MTY=","citycodes":[ "101010100", "101071201" ,"101090101","101091001"]}`
)

type payload struct {
	accesskey string `json:accesskey`
	citycodes string `json:citycodes`
}

const configName = "config"
const configPath = "./com.weather.get"

func SaveData2Prisma(recode api.Hybridization) { //构造一个函数 参数record是api包中定义的结构体
	opts := prisma.Options{
		Endpoint: prismaserver,
		Secret:   secret,
	}

	client := prisma.New(&opts)
	// client :=  prisma.New(nil)

	ctx := context.Background()

	cityCode, err := client.UpdateCityCodeInfo(
		prisma.CityCodeInfoUpdateParams{
			Where: prisma.CityCodeInfoWhereUniqueInput{
				// var myCitycode string
				// myCitycode="101010100"
				WeatherCityCode: &recode.Citycode,
				// AirCitycode: &recode.Citycode,
			},
			Data: prisma.CityCodeInfoUpdateInput{ //我要更新的数据
				WeatherLive: &prisma.WeatherLiveUpdateOneWithoutCityInfoInput{ //WeatherLive存取开始
					Update: &prisma.WeatherLiveUpdateWithoutCityInfoDataInput{

						// Citycode: &recode.Citycode,
						Rdesc: &recode.Rdesc,
						// Rcode: &recode.Rcode,

						Updatetime:  &recode.Updatetime,
						Windspeed:   &recode.Windspeed,
						Airpressure: &recode.Airpressure,
						Phenomena:   &recode.Phenomena,
						Humidity:    &recode.Humidity,
						Windpower:   &recode.Windpower,
						Feelst:      &recode.Feelst,
						Winddirect:  &recode.Winddirect,
						Rain:        &recode.Rain,
						Temperature: &recode.Temperature,
					},
				}, //WeatherLive存取结束

				// AirLiveData: &prisma.AirLiveDataUpdateOneWithoutCityInfoInput{ //存空气实况存取开始	测试
				// 	Update: &prisma.AirLiveDataUpdateWithoutCityInfoDataInput{
				// 		// Citycode: &recode.Citycode,
				// 		// Rdesc:    &recode.Rdesc,
				// 		// Rcode:    &recode.Airrcode,

				// 		Time:    &recode.Time,
				// 		Pm25:    &recode.PM25,
				// 		Pm10:    &recode.PM10,
				// 		So2:     &recode.SO2,
				// 		O3:      &recode.O3,
				// 		No2:     &recode.NO2,
				// 		Primary: &recode.Primary,
				// 		Co:      &recode.CO,
				// 		Aqi:     &recode.AQI,
				// 	},
				// }, //存空气实况存取结束	测试

				// 			//-----------------------------------存历史测试------------------------------------------------
				// 			// WeatherLiveHistory:&prisma.WeatherLiveHistoryUpdateManyWithoutCityInfoInput{
				// 			// Connect:&prisma.WeatherLiveHistoryWhereUniqueInput{},
				// 			// Create:&prisma.WeatherLiveHistoryWhereUniqueInput{},
				// 			// UpdateMany:prisma.WeatherLiveHistoryUpdateManyWithWhereNestedInput{
				// 			// Data:prisma.WeatherLiveHistoryScalarWhereInput {

				// 			// 	}，
				// 			// },//对应Connect Create功能体的括号

				// 			// 			},

				// 			//-----------------------------------测玩--------------------------------------失败

			},
		},
	).Exec(ctx)

	if err != nil {
		fmt.Printf("err: %+v", err)
	}
	fmt.Printf("cityCode:%v\r\n", cityCode)

	//存一堆历史数据（创建数据全部保留）
	wlhistoryData, err := client.CreateWeatherHistoryData(
		prisma.WeatherHistoryDataCreateInput{
			CityInfo: &prisma.CityCodeInfoCreateOneWithoutWeatherHistoryDatasInput{ //								CityInfo		CityCodeInfoCreateOneWithoutWeatherLiveHistoryInput				CityCodeInfoUpdateOneWithoutWeatherLiveHistoryInput
				Connect: &prisma.CityCodeInfoWhereUniqueInput{
					WeatherCityCode: &recode.Citycode,
				},
			},
			Windspeed:   &recode.Windspeed,
			Citycode:    recode.Citycode,
			Airpressure: &recode.Airpressure,
			Phenomena:   &recode.Phenomena,
			Rdesc:       &recode.Rdesc,
			Humidity:    &recode.Humidity,
			Updatetime:  &recode.Updatetime,
			Windpower:   &recode.Windpower,
			Feelst:      &recode.Feelst,
			Winddirect:  &recode.Winddirect,
			// Rcode:       &recode.Rcode,
			Rain:        &recode.Rain,
			Temperature: &recode.Temperature,
		},
	).Exec(ctx)
	//-----------------------------------------------实时历史数据保存完毕----------------------------------

	if err != nil {
		fmt.Printf("err: %+v", err)
	}
	fmt.Printf("cityCode: %v\r\n", wlhistoryData)
	fmt.Printf("historyData: %v\r\n", wlhistoryData)

} //savedate2prisma函数体结束

func main() {
	// historyDatas := prisma.HistoryDatas(nil).Exec()

	// 设置viper参数配置文件
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	// 使用viper读取参数配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf(fmt.Sprintf("Fatal error when reading %s config file:%s", configName, err))
	}
	//使用viper读取参数配置
	prismaserver = viper.GetString("PrismaServer.address") //address = "202.206.168.131:4466"
	secret = viper.GetString("PrismaServer.secret")        //secret = ""

	// fmt.Println(api.WeatherData(a, "weatherlive", c, "101010100"))
	for {
		// SaveData2Prisma(api.WeatherData(a, "weatherlive", c, "101010100", "打印")) //api.Weather是返回了json值
		// SaveData2Prisma(api.WeatherData(a, "weatherlive", c, "101010300", "打印")) //api.Weather是返回了json值
		// SaveData2Prisma(api.WeatherData(a, "weatherlive", c, "101090101", "打印")) //api.Weather是返回了json值
		// SaveData2Prisma(api.WeatherData(a, "weatherlive", c, "101091001", "打印")) //api.Weather是返回了json值
		api.WeatherData(a, b, "weatherlive", postbody, "打印")
		api.WeatherData(a, b, "weatherlive", apostbody, "打印")
		// api.WeatherData(a,b,"cityairlive","")
		// SaveData2Prisma(api.WeatherData(a, "cityairlive", c, "101010100", "打印")) //api.Weather是返回了json值
		// SaveData2Prisma(api.WeatherData(a, "cityairlive", c, "101071201", "打印")) //api.Weather是返回了json值
		// SaveData2Prisma(api.WeatherData(a, "cityairlive", c, "101090101", "打印")) //api.Weather是返回了json值
		// SaveData2Prisma(api.WeatherData(a, "cityairlive", c, "101091001", "打印")) //api.Weather是返回了json值
		time.Sleep(20 * time.Minute) //定时器20分钟一次
	}
	// fmt.Println(api.WeatherData(a, "weatherlive", c, "101010100", "调试"))
	// fmt.Println(variable)

}
