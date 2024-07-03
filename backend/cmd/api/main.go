package main

import (
	"backend/internal/server"
	"fmt"
)

func main() {

	server, err := server.NewServer()
	if err != nil {
		panic(fmt.Sprintf("cannot instantiate server: %s", err))
	}

	err = server.Run()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
