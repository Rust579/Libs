package uniapi

import (
	"Libs/grpc"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	client "tea.gitpark.ru/sast/shpack/uniapi_v2/goclient"
	"tea.gitpark.ru/sast/shpack/uniapi_v2/proto/rpc"
	"time"
)

func DockerGRPCZips(ctx context.Context, projectName, projectId, tempDir, port string) {
	client, err := client.NewClient(projectId, "localhost", port, ZipResponseWithTime)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(projectName)
	if err != nil {
		return
	}

	_, path, _, _, err := grpc.SeparateZipArchive(buffer1, tempDir)
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

func DockerGRPCZip(projectName, projectId, tempDir, port string) {

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(projectName)
	if err != nil {
		return
	}

	_, path, _, _, err := grpc.SeparateZipArchive(buffer1, tempDir)
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

func DockerGRPCDiff(projectId, port string) {
	//fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "5954e382342773cff67b1f61fd22d5567957d8ffccc4b193ed8ab72c8c5316c7",
			FileStatus:  "upd",
			FileName:    "Dockerfile",
			Data: []*rpc.Hunk{
				{
					OldStart: 1,
					OldLines: 0,
					NewStart: 1,
					NewLines: 1,
					Lines: []string{
						"+# fff",
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

func DockerGRPCDiff2(projectId, port string) {

	fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "bc0bef226f255844418bf9377f5c02582fd813a2b8e511fb306948e534ec08f8",
			FileStatus:  "upd",
			FileName:    "Dockerfile",
			Data: []*rpc.Hunk{
				{
					OldStart: 1,
					OldLines: 0,
					NewStart: 1,
					NewLines: 1,
					Lines: []string{
						"+# bbb",
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
