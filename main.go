package main

import (
	io "io/ioutil"
	json "encoding/json"
	"fmt"
)

type JsonStruct struct{

}



func NewJsonStruct () *JsonStruct {

	return &JsonStruct{}

}



func (self *JsonStruct) Load (filename string, Cfg interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil{

		return

	}

	datajson := []byte(data)

	//fmt.Println(datajson)

	err = json.Unmarshal(datajson, Cfg)

	if err != nil{

		return

	}

}



type TDB struct {
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

	chWindCode string

	chMarket   string
}

type conf struct {

	TDBConf TDB

	InfluxdbConf Influxdb

}

func main() {

	JsonParse := NewJsonStruct()

	Cfg := conf{}

	JsonParse.Load("conf.json", &Cfg)

	fmt.Println(Cfg)

	fmt.Println(Cfg.InfluxdbConf.chWindCode)
}



