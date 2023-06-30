package models

import "github.com/jinzhu/gorm"

type Major struct {
	Name     string `json:"name"`
	SerialNo string `json:"serial_no"`
	SchoolID int    `json:"school_id"`
	Model
}

func GetMajor(schoolId int, name string) (*Major, error) {
	var major Major
	err := db.Where("school_id = ? AND name = ?", schoolId, name).First(&major).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &major, nil
}

func CreateMajor(data map[string]interface{}) error {
	major := Major{
		Name:     data["name"].(string),
		SerialNo: data["serialNo"].(string),
		SchoolID: data["schoolID"].(int),
	}
	var duplicateMajor *Major

	duplicateMajor, _ = GetMajor(major.SchoolID, major.Name)

	if duplicateMajor == nil {

		if err := db.Create(&major).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetMajors(schoolId int) ([]*Major, error) {
	var majors []*Major
	err := db.Where("school_id = ?", schoolId).Find(&majors).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return majors, nil
}
