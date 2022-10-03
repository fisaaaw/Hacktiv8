package main

import (
	"swag-tutor/app/interface/container"
	"swag-tutor/app/interface/server"
)

func main() {
	server.StartServer(*container.SetUpContainer())
}