package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	PrefixUrl       string
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

var ServerSetting = &Server{}

type Databse struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Databse{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)
	mapTo("redis", RedisSetting)
	mapTo("app", AppSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %v", err)
	}
}
