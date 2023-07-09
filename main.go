package main

import (
	"Packetor/packetor"
	"Packetor/packetor/registries"
	"bufio"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	reg := registries.GetRegistry()
	err = reg.Load("1.20")
	if err != nil {
		logrus.Fatal("failed to load data ", err)
	}
	ctx, cancelFn := context.WithCancel(context.Background())
	fronter := packetor.NewFronter(
		ctx,
		"tcp",
		"127.0.0.1:25566",
		10*time.Second,
		10*time.Second)
	err = fronter.Bind("tcp", "127.0.0.1:25565", 10*time.Second)
	if err != nil {
		logrus.Fatal("failed to bind ", err)
	}
	inReader := bufio.NewReader(os.Stdin)
	for {
		lineB, _, err := inReader.ReadLine()
		if err != nil {
			logrus.Error("failed to read line")
		}
		line := string(lineB)
		if line == "exit" {
			cancelFn()
			break
		} else if line == "mem" {
			stats := runtime.MemStats{}
			runtime.ReadMemStats(&stats)
			logrus.Info("mem: ", float64(stats.HeapAlloc)/1000000)
		}
	}
}
