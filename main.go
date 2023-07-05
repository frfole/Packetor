package main

import (
	"Packetor/packetor"
	"Packetor/packetor/registries"
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	reg := registries.GetRegistry()
	err := reg.Load("1.20")
	if err != nil {
		logrus.Fatal("failed to load data ", err)
	}
	fronter := packetor.NewFronter(
		context.Background(),
		"tcp",
		"127.0.0.1:25566",
		10*time.Second,
		10*time.Second)
	err = fronter.Bind("tcp", "127.0.0.1:25565", 10*time.Second)
	if err != nil {
		logrus.Fatal("failed to bind ", err)
	}
	select {}
}
