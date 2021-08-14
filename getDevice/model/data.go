// Package model is package for present input and output model
package model
//Input is structure of input data in lambda func
type Input struct {
	Id string `json:"id"`
}
//Output is structure of output data in lambda func
type Output struct {
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}
