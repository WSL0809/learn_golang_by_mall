package main

import (
	"init_golang/conf"
	"init_golang/routes"
)

func main() {

    conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}