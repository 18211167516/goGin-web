package config

import (
	"gopkg.in/ini.v1"
	"log"
	"fmt"
	"os"
	"time"
)

type App struct {
	Template string
	PageSize int
	JwtSecret string
	JwtExpiresAt time.Duration
	SigningMethod string
}

var AppSetting = &App{}

type Server struct {
	HttpAddress string
	HttpPort int
	ReadTimeout  time.Duration
	WriteTimeout  time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	MysqlUser string
	MysqlPassword string
	MysqlHost string
	MysqlName string
	MysqlPrefix string
}

var DatabaseSetting = &Database{}

var (
	Cfg *ini.File
	RunMode string
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
	
	err = sec.MapTo(AppSetting)
	if err != nil {
        log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
    }

	AppSetting.JwtExpiresAt = time.Duration(sec.Key("JWT_EXPIRE_TIME").MustInt(10))*time.Minute
}

//加载http服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("Fail to get section 'server': %v", err)
	}
	
	err = sec.MapTo(ServerSetting)
	if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }

    ServerSetting.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
    ServerSetting.WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second 
}

//加载数据库配置
func LoadDatabase() {
	sec, err := Cfg.GetSection("database-mysql")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
    }

	err = sec.MapTo(DatabaseSetting)
	if err != nil {
        log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}


