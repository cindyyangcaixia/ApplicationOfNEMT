package models

import (
	"log"

	"github.com/cindyyangcaixia/gin-example/pkg/e"
	"github.com/jinzhu/gorm"
)

type School struct {
	Name     string `json:"name"`
	SerialNo string `json:"serial_no"`
	Majors   []Major
	Model
}

func CreateSchool(name string, serialNo string) (*School, int) {
	school := School{
		Name:     name,
		SerialNo: serialNo,
	}
	var duplicateSchool *School

	duplicateSchool, err := GetSchool(school.Name, school.SerialNo)
	if err != nil {
		return nil, e.ERROR
	}

	log.Println(duplicateSchool)
	log.Println(duplicateSchool.ID)

	// if IsExist(duplicateSchool) {
	// 	return nil, e.SCHOOL_EXIST
	// }

	if duplicateSchool.ID > 0 {
		return nil, e.SCHOOL_EXIST
	}

	if err := db.Create(&school).Error; err != nil {
		return nil, e.ERROR
	}

	return &school, e.SUCCESS
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
