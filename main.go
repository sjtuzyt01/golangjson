package main

import (
	io "io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
)

type JsonStruct struct{

}



func NewJsonStruct () *JsonStruct {

	return &JsonStruct{}

}



func (self *JsonStruct) Load (filename string, v interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil{

		return

	}

	datajson := []byte(data)

	err = json.Unmarshal(datajson, v)

	if err != nil{

		return

	}

}



type Thandle struct {
	SzIP       string

	SzPort     string

	SzUser     string

	SzPassword string
}

type Influxdb struct {
	Addr       string

	LocalAddr  string

	Username   string

	Password   string

	StartTime  string

	EndTime    string

	ChWindCode string

	ChMarket   string
}

type Function struct {
	KLine       bool

	Tick        bool

	Transaction bool

	Order       bool

	OrderQueue  bool
}

type conf struct {

	TDBConf Thandle `json:"Thandle"`

	Influxconf Influxdb  `json:"Influxdb"`

	Functionconf Function `json:"Function"`
}

func timeSplit(time string)(int, int, int){
	year,_ := strconv.Atoi(time[0:4])
	month,_ := strconv.Atoi(time[4:6])
	day,_ := strconv.Atoi(time[6:8])
	return year, month, day
}

func main() {

	JsonParse := NewJsonStruct()

	cfg := conf{}

	JsonParse.Load("conf.json", &cfg)

	fmt.Println(cfg)

	startYear, startMonth, startDay := timeSplit(cfg.Influxconf.StartTime)

	endYear, endMonth, endDay := timeSplit(cfg.Influxconf.EndTime)

	println(startYear, startMonth, startDay)

	println(endYear, endMonth, endDay)
}



