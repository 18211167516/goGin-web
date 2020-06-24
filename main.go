package main

import (
	"gintest/routes"
	"gintest/config"
	"fmt"
)

func main() {
	r := routes.InitRouter()
	r.Run(fmt.Sprintf("%s:%d",config.ServerSetting.HttpAddress,config.ServerSetting.HttpPort))
}
