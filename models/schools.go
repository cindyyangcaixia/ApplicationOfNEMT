package models

type School struct {
	Name     string `json:"name"`
	SerialNo string `json:"serial_no"`
	Majors   []Major
	Model
}
