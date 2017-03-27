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



func (self *JsonStruct) Load (filename string, v interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil{

		return

	}

	datajson := []byte(data)

	fmt.Println(datajson)

	err = json.Unmarshal(datajson, v)

	if err != nil{

		return

	}

}



type ValueTestAtmp1 struct{

	StringValue string

	NumericalValue int

	BoolValue bool

}

type ValueTestAtmp2 struct{

	FloatValue float64

}

type testdata struct {

	ValueTestA ValueTestAtmp1

	ValueTestB ValueTestAtmp2

}



func main() {

	JsonParse := NewJsonStruct()

	v := testdata{}

	JsonParse.Load("conf.json", &v)

	fmt.Println(v)

	fmt.Println(v.ValueTestA.StringValue)

}



