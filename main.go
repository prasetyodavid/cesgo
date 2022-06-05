package main

import (
	"cesgo/server"
)

func main() {
	s := server.NewServer()
	s.Start(":8000")
}
