package main

import (
	"gintest/routes"
	"gintest/config"
)

func main() {
	r := routes.InitRouter()
	r.Run(fmt.Sprintf("%s:%d",config.Server.Address,config.Server.Port))
}
