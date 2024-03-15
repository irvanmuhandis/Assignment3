package utill

import (
	"assignment3/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FetchData() []byte {
	data, err := ioutil.ReadFile("status.json")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return data
}

func ConvertMap(jsonData []byte) map[string]interface{} {
	var MAP = make(map[string]interface{})
	err := json.Unmarshal(jsonData, &MAP)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return MAP["status"].(map[string]interface{})
}

func ConvertStruct(jsonData []byte) model.Status {

	var STRUCT = model.JsonFile{}
	err := json.Unmarshal(jsonData, &STRUCT)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return STRUCT.Status
}
