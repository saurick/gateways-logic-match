package main

import (
	"flag"
	"gateways-logic-match/server"
)

var (
	port = flag.Int("port", 30052, "The gateway server port")
)

func main() {
	flag.Parse()
	//	启动网关匹配服务
	server.LaunchGatewaysLogicMatch(port)
}
