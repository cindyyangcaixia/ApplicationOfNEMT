package models

type Major struct {
	Name     string `json:"name"`
	SerialNo string `json:"serial_no"`
	SchoolID int    `json:"school_id"`
	Model
}
