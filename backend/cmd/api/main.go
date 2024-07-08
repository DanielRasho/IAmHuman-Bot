package main

import (
	"backend/internal/server"
	"fmt"
	"time"
)

func main() {

	time.Sleep(3 * time.Second)
	server, err := server.NewServer()
	if err != nil {
		panic(fmt.Sprintf("cannot instantiate server: %s", err))
	}

	err = server.Run()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
