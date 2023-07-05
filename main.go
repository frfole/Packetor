package main

import (
	"Packetor/packetor"
	"Packetor/packetor/data"
	"context"
	"log"
	"time"
)

func main() {
	registry := data.GetBlockRegistry()
	err := registry.Load("1.20")
	if err != nil {
		log.Fatal("failed to load data", err)
	}
	fronter := packetor.NewFronter(
		context.Background(),
		"tcp",
		"127.0.0.1:25566",
		10*time.Second,
		10*time.Second)
	err = fronter.Bind("tcp", "127.0.0.1:25565", 10*time.Second)
	if err != nil {
		log.Println("failed to bind", err)
	}
	select {}
}
