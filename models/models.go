package models

import (
	"fmt"
	"log"
	"time"

	"github.com/cindyyangcaixia/gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type BaseModel struct {
	ID        int       `gorm:"primary_key" json:"id" "auto_increment"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"default:null"`
}

type Model struct {
	BaseModel
	CreatedBy string `json:"created_by" gorm:"default:null"`
	UpdatedBy string `json:"updated_by" gorm:"default:null"`
	DeleteBy  string `json:"delete_by" gorm:"default:null"`
}

func Setup() {
	var err error

	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// todo migration
	db.AutoMigrate(&Major{})
	db.AutoMigrate(&School{})
}

func CloseDB() {
	defer db.Close()
}

// func IsExist(obj interface{}) bool {
// 	return obj != nil && obj.ID > 0
// }
