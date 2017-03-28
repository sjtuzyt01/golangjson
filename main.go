package main

import (
	io "io/ioutil"
	"encoding/json"
	"fmt"
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

type conf struct {

	TDBConf Thandle `json:"Thandle"`

	Influxconf Influxdb  `json:"Influxdb"`

}

func main() {

	JsonParse := NewJsonStruct()

	cfg := conf{}

	JsonParse.Load("conf.json", &cfg)

	fmt.Println(cfg)

	fmt.Println(cfg.Influxconf.LocalAddr)
}



