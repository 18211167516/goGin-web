package config

import (
	"gopkg.in/ini.v1"
	"log"
	"fmt"
	"os"
)

var (
	Cfg *ini.File
	RunMode string
	HttpAddress string
	HttpPort string
	
)

func init() {
	fig, err := ini.Load("./config/app.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
	}
	Server = &Http{}
	Server.Address = fig.Section("http").Key("address").MustString("0.0.0.0")
	Server.Port = fig.Section("http").Key("port").MustInt(8080)
}


