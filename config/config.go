package config

import (
	"gopkg.in/ini.v1"
	"log"
	"fmt"
	"os"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	//server
	HttpAddress string
	HttpPort int
	ReadTimeout  time.Duration
	WriteTimeout  time.Duration
	//app
	PageSize int

	//database-mysql
	MysqlUser string
	MysqlPassword string
	MysqlHost string
	MysqlName string
	MysqlPrefix string
	
)

func init() {
	var err error
	Cfg, err = ini.Load("./config/app.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	LoadBasice()
	LoadApp()
	LoadServer()
	LoadDatabase()
}

//加载基础配置
func LoadBasice() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//加载app配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
    }

    //JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
    PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

//加载http服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("Fail to get section 'server': %v", err)
    }

	HttpAddress = sec.Key("HTTP_ADDRESS").MustString("0.0.0.0")
    HttpPort = sec.Key("HTTP_PORT").MustInt(8080)
    ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
    WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second 
}

//加载数据库配置
func LoadDatabase() {
	sec, err := Cfg.GetSection("database-mysql")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
    }

    MysqlUser = sec.Key("USER").MustString("!@)*#)!@U#@*!@!)")
	MysqlPassword = sec.Key("PASSWORD").MustString("123456")
	MysqlHost = sec.Key("HOST").MustString("127.0.0.1:3306")
	MysqlName = sec.Key("NAME").MustString("test")
	MysqlPrefix = sec.Key("TABLE_PREFIX").MustString("test_")
}


