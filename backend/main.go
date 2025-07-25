package main

import (
	"alexandrie/server"
	"fmt"
)

func main() {
	server, application := server.SetupServer()
	defer application.DB.Close()

	server.Run("192.168.0.25:" + fmt.Sprintf("%d", application.Config.Port))
}
