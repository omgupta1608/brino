package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//LoadData - loading biolerplates data from json
func LoadData() map[string]interface{} {
	bpData, err := os.Open("boilerplates.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer bpData.Close()

	byteValue, _ := ioutil.ReadAll(bpData)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

/*
JSON to Go Map

map[boilerplates:
        [map[test:
            map[
                commands:[go version]
                id:1
                pre_req:[go]
            ]
        ]
        map[git:
            map[
                commands:[git --version]
                id:2
                pre_req:[git]
            ]
        ]
    ]
] */
