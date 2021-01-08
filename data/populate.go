package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	models "github.com/omgupta1608/brino/models"
)

//LoadData - loading biolerplates data from json
func LoadData() map[string]models.BoilerPlate {
	var bp []models.BoilerPlate
	bpData, err := os.Open("boilerplates.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer bpData.Close()

	byteValue, _ := ioutil.ReadAll(bpData)

	json.Unmarshal([]byte(byteValue), &bp)

	var bpHash = make(map[string]models.BoilerPlate)
	for i := 0; i < len(bp); i++ {
		bpHash[bp[i].ID] = bp[i]
	}

	return bpHash
}
