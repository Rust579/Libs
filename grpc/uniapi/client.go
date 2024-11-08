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

func ZipResponse(resp *rpc.Response) error {
	log.Printf("Upload response: %v,%v", resp.Success, resp.ProjectID)
	return nil
}

func GRPCZips(projectName, projectId, tempDir, port string) {

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

	for i := 0; i < 10; i++ {

		err = client.SendProjectZip(strconv.Itoa(i), buffer2)
		if err != nil {
			fmt.Println("err", err)
		}
		log.Println("SendProjectZip", strconv.Itoa(i))

		time.Sleep(100 * time.Millisecond)

		files := []*rpc.DiffFile{
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

		err = client.SendDiff(strconv.Itoa(i), files)
		if err != nil {
			fmt.Println("err", err)
		}
		log.Println("SendDiff", strconv.Itoa(i))

	}
}

func GRPCZip(projectName, projectId, tempDir, port string) {

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

func GRPCLibs(libName, projectId, port string) {
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
			Path: "$GOPATH/pkg/mod/github.com/unione-pro/core@v0.1.68",
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

func GRPCLibs2(libName, projectId, port string) {
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
			Path: "$GOPATH/pkg/mod/github.com/unione-pro/core@v0.1.69",
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

func GRPCDiff(projectId, port string) {
	//fmt.Println("Start gRPC client for diff handler")

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
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

	err = client.SendDiff(projectId, files)
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("SendDiff1", projectId)
}

func GRPCDiff2(projectId, port string) {

	client, err := client.NewClient(projectId, "localhost", port, ZipResponse)
	if err != nil {
		fmt.Println("Run error", err)
		return
	}

	files := []*rpc.DiffFile{
		{
			FileHashSum: "08e2036ea0f0a6f66219d5e32b24639c90c2652d784847e7f513a82b045594cb",
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
