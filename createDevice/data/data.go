// Package data data is package for present input and output model
package data

//Output is model for sending data to api getway
type Output struct {
	Message string `json:"message"`
}

//Input is model for incoming data from api getway
type Input struct {
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}
