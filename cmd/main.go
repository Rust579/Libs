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

	projectZip := "C:/Projects_Go/testProjects/GoTestWithOutCore.zip"
	//libsZip1 := "C:/Projects_Go/testProjects/core68.zip"
	//libsZip2 := "C:/Projects_Go/testProjects/core69.zip"
	projectId := "aaa"
	unzipDir := "unzipped"
	port := "50051"

	uniapi.GRPCZips(projectZip, projectId, unzipDir, port)

	/*for i := 0; i < 1; i++ {
		go func() {
			go func() {
				uniapi.GRPCZip(projectZip, strconv.Itoa(i), strconv.Itoa(i), port)
			}()
			time.Sleep(300 * time.Millisecond)
			go func() {
				uniapi.GRPCDiff(strconv.Itoa(i), port)
			}()
			go func() {
				uniapi.GRPCDiff2(strconv.Itoa(i), port)
			}()
		}()
	}*/

	/*go func() {
		uniapi.GRPCZip(projectZip, projectId, unzipDir, port)
	}()

	time.Sleep(10 * time.Second)

	go func() {
		uniapi.GRPCLibs(libsZip1, projectId, port)
	}()

	time.Sleep(15 * time.Second)

	go func() {
		uniapi.GRPCDiff2(projectId, port)
	}()

	time.Sleep(15 * time.Second)

	go func() {
		uniapi.GRPCLibs2(libsZip2, projectId, port)
	}()*/

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	<-ctx.Done()
}
