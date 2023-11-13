package main

import (
	"path.finder/ai/server"
)

func main() {
	s := server.NewServer("0.0.0.0", "8080")
	s.Run()
}
