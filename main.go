package main

import (
	"gintest/routes"
	"gintest/config"
	"fmt"
)

func main() {
	config.InitConfig("config/app.ini")
	r := routes.InitRouter()
	r.Run(fmt.Sprintf("%s:%d",config.ServerSetting.HttpAddress,config.ServerSetting.HttpPort))
}
