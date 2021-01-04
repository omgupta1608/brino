package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bpData, err := os.Open("boilerplates.json")
	if err != nil {
		fmt.Println(err)
	}

	defer bpData.Close()

	byteValue, _ := ioutil.ReadAll(bpData)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
	// if len(os.Args) > 1 {
	// 	switch os.Args[1] {
	// 	case "hello":
	// 		fmt.Println("Hello")
	// 	case "bye":
	// 		fmt.Println("bye")
	// 	}
	// }
}
