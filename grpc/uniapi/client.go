package uniapi

import (
	"fmt"
	"log"
	"sync"
	"tea.gitpark.ru/sast/shpack/uniapi_v2/proto/rpc"
	"time"
)

type RequestTime struct {
	StartTime time.Time
	EndTime   time.Time
}

var (
	RequestTimes = make(map[string]*RequestTime)
	Mu           sync.Mutex
)

func GoRunTest() {

	fmt.Println("GoWr test run")

	projectZip := "grpc/GoTest.zip"
	projectId := "aaa"
	unzipDir := "unzipped"
	port := "50051"

	GoGRPCZip(projectZip, projectId, unzipDir, port)

	time.Sleep(5 * time.Second)

	GoGRPCDiff(projectId, port)

	//libsZip1 := "C:/Projects_Go/testProjects/core68.zip"
	//libsZip2 := "C:/Projects_Go/testProjects/core69.zip"

	//uniapi.GRPCZips(ctx, projectZip, projectId, unzipDir, port)

	/*go func() {
		uniapi.GoGRPCZip(projectZip, projectId, unzipDir, port)
	}()

	time.Sleep(20 * time.Second)

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
}

func JSRunTest() {

	fmt.Println("JSWr test run")

	projectZip := "grpc/JsTest.zip"
	projectId := "aaa"
	unzipDir := "unzipped"
	port := "50054"

	JSGRPCZip(projectZip, projectId, unzipDir, port)

	time.Sleep(5 * time.Second)

	JSGRPCDiff(projectId, port)

	//libsZip1 := "C:/Projects_Go/testProjects/core68.zip"
	//libsZip2 := "C:/Projects_Go/testProjects/core69.zip"
}

func PyRunTest() {

	fmt.Println("PyWr test run")

	projectZip := "grpc/PyTest.zip"
	projectId := "bbb"
	unzipDir := "unzipped"
	port := "50055"

	PyGRPCZip(projectZip, projectId, unzipDir, port)

	time.Sleep(2 * time.Second)

	PyGRPCDiff(projectId, port)

	//libsZip1 := "C:/Projects_Go/testProjects/core68.zip"
	//libsZip2 := "C:/Projects_Go/testProjects/core69.zip"
}

func DockerRunTest() {

	fmt.Println("DockerWr test run")

	projectZip := "grpc/Dockerfile.zip"
	projectId := "aaa"
	unzipDir := "unzipped"
	port := "50052"

	DockerGRPCZip(projectZip, projectId, unzipDir, port)

	time.Sleep(5 * time.Second)

	DockerGRPCDiff(projectId, port)
}

func ZipResponseWithTime(resp *rpc.Response) error {
	Mu.Lock()
	defer Mu.Unlock()

	projectID := resp.ProjectID
	log.Printf("Upload response: %v,%v", resp.Success, projectID)

	if reqTime, ok := RequestTimes[projectID]; ok {
		currentEndTime := time.Now()
		if reqTime.EndTime.IsZero() || currentEndTime.After(reqTime.EndTime) {
			reqTime.EndTime = currentEndTime
		}
		RequestTimes[projectID] = reqTime

	} else {
		log.Printf("No request time found for projectID %s", projectID)
	}

	return nil
}

func ZipResponse(resp *rpc.Response) error {

	log.Printf("Upload response: %v", resp)
	return nil
}
