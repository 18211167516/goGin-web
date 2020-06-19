package main

import (
	"gintest/routes"
	"gintest/config"
	"fmt"
	//"net/http"
)

func main() {
	r := routes.InitRouter()
	r.Run(fmt.Sprintf("%s:%d",config.HttpAddress,config.HttpPort))
}
