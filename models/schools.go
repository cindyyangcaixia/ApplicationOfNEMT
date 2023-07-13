package models

import (
	"log"
	"net/http"

	"github.com/cindyyangcaixia/ApplicationOfNEMT/pkg/app"
	"github.com/cindyyangcaixia/ApplicationOfNEMT/pkg/e"
	"github.com/jinzhu/gorm"
)

type School struct {
	Name     string `json:"name"`
	SerialNo string `json:"serialNo"`
	Majors   []Major
	Model
}

func CreateSchool(name string, serialNo string) (*School, *app.ResponseMessage) {
	school := School{
		Name:     name,
		SerialNo: serialNo,
	}
	var duplicateSchool *School

	duplicateSchool, err := GetSchool(school.Name, school.SerialNo)
	if err != nil {
		return nil, &app.ResponseMessage{Status: http.StatusInternalServerError, Code: e.ERROR, Message: err.Error()}
	}

	// if IsExist(duplicateSchool) {
	// 	return nil, e.SCHOOL_EXIST
	// }

	if duplicateSchool.ID > 0 {
		return nil, &app.ResponseMessage{Status: http.StatusBadRequest, Code: e.SCHOOL_EXIST}
	}

	if err := db.Create(&school).Error; err != nil {
		return nil, &app.ResponseMessage{Status: http.StatusInternalServerError, Code: e.ERROR, Message: err.Error()}
	}

	// why deleteAt is not null ? todo
	return &school, nil
}

func GetSchool(name string, serialNo string) (*School, error) {
	var school School
	err := db.Where("name = ? OR serial_no = ?", name, serialNo).Select("id").First(&school).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	log.Println(school)
	return &school, nil
}
