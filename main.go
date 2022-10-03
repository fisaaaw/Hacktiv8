package main

import (
	"ass-david/app/interface/container"
	"ass-david/app/interface/server"
)

func main() {
	container := container.SetupContainer()
	server.SetupServer(container)
}
