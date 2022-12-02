package config

import (
	"log"
	"mymodule/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string //サーバーのポート番号
	SQLDriver string //SQLドライバーのツリー設定
	DbName    string //データーベース名
	LogFile   string //Logを残す
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
