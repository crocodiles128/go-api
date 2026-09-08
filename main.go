package main

import "github.com/crocodiles128/go-api/server"

func main() {
	server := server.NewServer("5000")

	server.Run()
}
