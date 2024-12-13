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

func GoGRPCZips(ctx context.Context, projectName, projectId, tempDir, port string) {
	client, err := client.NewClient(projectId, "localhost", port, ZipResponseWithTime)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(projectName)
	if err != nil {
		return
	}

	path, _, _, _, err := grpc.SeparateZipArchive(buffer1, tempDir)
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

func GoGRPCZip(projectName, projectId, tempDir, port string) {

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	buffer1, err := grpc.LoadLocalFile(projectName)
	if err != nil {
		return
	}

	path, _, _, _, err := grpc.SeparateZipArchive(buffer1, tempDir)
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

func GoGRPCDiff(projectId, port string) {
	//fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "b27e150b1ff864c93761512b93e042bf9b03a316de439e9411dde11b7c54597a",
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

	err = client.SendDiff(projectId, files)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("SendDiff1", projectId)
}

func GoGRPCDiff2(projectId, port string) {

	fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "e60a80be40369b1e4de8cb17acd8e267051b9bc29325b7c62acc56735368dde2",
			FileStatus:  "upd",
			FileName:    "main.go",
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

func GoGRPCLibs(libName, projectId, port string) {
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

func GoGRPCLibs2(libName, projectId, port string) {
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
