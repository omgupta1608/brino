package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadData() map[string]interface{} {
	bpData, err := os.Open("boilerplates.json")
	if err != nil {
		fmt.Println(err)
	}
	defer bpData.Close()

	byteValue, _ := ioutil.ReadAll(bpData)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
