package main

import (
	"Libs/configs"
	"Libs/grpc/uniapi"
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

var gfsd = make(chan bool)

func main() {

	err := configs.Init()
	if err != nil {
		logrus.Fatalf("error open config file: %s", err.Error())
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	//uniapi.GoRunTest()
	//uniapi.DockerRunTest()
	//uniapi.JSRunTest()
	uniapi.PyRunTest()

	<-ctx.Done()
}
