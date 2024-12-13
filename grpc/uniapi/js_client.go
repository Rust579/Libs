package uniapi

import (
	"Libs/grpc"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	client "tea.gitpark.ru/sast/shpack/uniapi_v2/goclient"
	"tea.gitpark.ru/sast/shpack/uniapi_v2/proto/rpc"
	"time"
)

func JSGRPCZips(ctx context.Context, projectName, projectId, tempDir, port string) {
	client, err := client.NewClient(projectId, "localhost", port, ZipResponseWithTime)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(projectName)
	if err != nil {
		return
	}

	_, _, path, _, err := grpc.SeparateZipArchive(buffer1, tempDir)
	if err != nil {
		fmt.Println("Error separate zip archive:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	buffer2, err := grpc.LoadLocalFile(path)
	if err != nil {
		return
	}

	var totalDuration time.Duration
	var requestCount int

	for i := 40; i < 50; i++ {
		//time.Sleep(100 * time.Millisecond)
		go func() {

			requestID := strconv.Itoa(i)

			Mu.Lock()
			RequestTimes[requestID] = &RequestTime{StartTime: time.Now()}
			Mu.Unlock()

			err = client.SendProjectZip(requestID, buffer2)
			if err != nil {
				fmt.Println("err", err)
			}
			log.Println("SendProjectZip", requestID)

			/*files := []*rpc.DiffFile{
				{
					FileHashSum: "9f4d2724c2178cb51901681285662fbb235f05e3e9f69da0e7f5c59e0ec1299a",
					FileStatus:  "upd",
					FileName:    "main.go",
					Data: []*rpc.Hunk{
						{
							OldStart: 1,
							OldLines: 0,
							NewStart: 1,
							NewLines: 1,
							Lines: []string{
								"+// fff",
							},
						},
					},
				},
			}

			err = client.SendDiff(requestID, files)
			if err != nil {
				fmt.Println("err", err)
			}
			log.Println("SendDiff", requestID)*/

		}()
	}

	<-ctx.Done()

	for i, reqTime := range RequestTimes {
		if !reqTime.EndTime.IsZero() {
			duration := reqTime.EndTime.Sub(reqTime.StartTime) //- 100*time.Millisecond
			totalDuration += duration
			requestCount++
			fmt.Println("requestTime", i, duration)
		}
	}

	if requestCount > 0 {
		averageDuration := totalDuration / time.Duration(requestCount)
		log.Printf("Average response time: %v", averageDuration)
	} else {
		log.Println("No responses received")
	}
}

func JSGRPCZip(projectName, projectId, tempDir, port string) {

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(projectName)
	if err != nil {
		return
	}

	_, _, path, _, err := grpc.SeparateZipArchive(buffer1, tempDir)
	if err != nil {
		fmt.Println("Error separate zip archive:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	buffer2, err := grpc.LoadLocalFile(path)
	if err != nil {
		return
	}

	err = client.SendProjectZip(projectId, buffer2)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("SendProjectZip", projectId)
}

func JSGRPCDiff(projectId, port string) {
	//fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "19dbb9fc6e0bbd6166b5932ec730fee964f11f999f7e1fe7eb35d398cae0a655",
			FileStatus:  "upd",
			FileName:    "jsexample.js",
			Data: []*rpc.Hunk{
				{
					OldStart: 1,
					OldLines: 0,
					NewStart: 1,
					NewLines: 1,
					Lines: []string{
						"+// fff",
					},
				},
			},
		},
	}

	err = client.SendDiff(projectId, files)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("SendDiff1", projectId)
}

func JSGRPCDiff2(projectId, port string) {

	fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "9fbb9e8744296e671292fa8e6f5432d928067cc84055f1ad8daf62b198a8b514",
			FileStatus:  "upd",
			FileName:    "jsexample.js",
			Data: []*rpc.Hunk{
				{
					OldStart: 1,
					OldLines: 0,
					NewStart: 1,
					NewLines: 1,
					Lines: []string{
						"+// bbb",
					},
				},
			},
		},
	}

	err = client.SendDiff(projectId, files)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("SendDiff2", projectId)
}

func JSGRPCLibs(libName, projectId, port string) {
	fmt.Println("Start gRPC client for libs handler")
	t1 := time.Now()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(libName)
	if err != nil {
		return
	}

	libs := []*rpc.Lib{
		{
			Path: "node_modules/my-private-lib", // Заменить на нужное название для Go либы
			Zip:  buffer1,
		},
	}

	err = client.SendLibsZip(projectId, libs)
	if err != nil {
		fmt.Println("err", err)
	}

	<-ctx.Done()

	log.Println("time cost:", time.Since(t1))

}

func JSGRPCLibs2(libName, projectId, port string) {
	fmt.Println("Start gRPC client for libs handler")
	t1 := time.Now()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer, err := grpc.LoadLocalFile(libName)
	if err != nil {
		return
	}

	libs := []*rpc.Lib{
		{
			Path: "node_modules/my-private-lib2", // Заменить на нужное название для Go либы
			Zip:  buffer,
		},
	}

	err = client.SendLibsZip(projectId, libs)
	if err != nil {
		fmt.Println("err", err)
	}

	<-ctx.Done()

	log.Println("time cost:", time.Since(t1))

}
