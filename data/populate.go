package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	models "github.com/omgupta1608/brino/models"
)

//LoadData - loading biolerplates data from json
func LoadData() []models.BoilerPlate {
	var bp []models.BoilerPlate
	bpData, err := os.Open("boilerplates.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer bpData.Close()

	byteValue, _ := ioutil.ReadAll(bpData)

	json.Unmarshal([]byte(byteValue), &bp)

	return bp
}
