package models

//BoilerPlate - JSON Array Element
type BoilerPlate struct {
	ID       string
	Commands []string
	PreReq   []string
	Details  string
}
