package main

import (
	"Packetor/packetor"
	"context"
	"log"
	"time"
)

func main() {
	fronter := packetor.NewFronter(
		context.Background(),
		"tcp",
		"127.0.0.1:25566",
		10*time.Second,
		10*time.Second)
	err := fronter.Bind("tcp", "127.0.0.1:25565", 10*time.Second)
	if err != nil {
		log.Println("failed to bind", err)
	}
	select {}
}
