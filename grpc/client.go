package grpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"tea.gitpark.ru/sast/shpack/uniapi/goclient"
	"tea.gitpark.ru/sast/shpack/uniapi/proto/rpc"
	"time"
)

func GRPCZip() {
	fmt.Println("Start gRPC client for zip handler")

	client := goclient.NewClient("localhost", "50051")

	filePath := "C:/Projects Go/dockerwrapper.zip"
	buffer1, err := loadLocalFile(filePath)
	if err != nil {
		return
	}

	tempDir := "unzipped"

	zip, _, err := SeparateZipArchive(buffer1, tempDir)
	if err != nil {
		fmt.Println("Error separate zip archive:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	buffer2, err := loadLocalFile(zip)
	if err != nil {
		return
	}

	t1 := time.Now()

	resp, err := client.SendStream(context.Background(), "1", buffer2)
	if err != nil {
		fmt.Println("err", err)
	}

	log.Println("time cost:", time.Since(t1))
	log.Printf("Upload response: %v", resp)
}

func GRPCDiff() {
	fmt.Println("Start gRPC client for diff handler")

	client := goclient.NewClient("localhost", "50051")

	req := rpc.DiffRequest{
		ProjectID:   "3",
		FileName:    "internal\\service\\logic\\logic_http.go",
		FileHashSum: "41c48e1bec341b318aace5fb5073c85c6471e63ef6021a19802dbee3b8299cc1",
		Data: &rpc.Diff{
			Line: 45,
			Text: "\"github.com/unione-pro/auth/internal/pkg/notifer\"",
		},
	}

	t1 := time.Now()

	resp, err := client.SendDiff(context.Background(), &req)
	if err != nil {
		fmt.Println("err", err)
	}

	log.Println("time cost:", time.Since(t1))
	log.Printf("Upload response: %v", resp)

}

func loadLocalFile(filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("err file info", err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("err buffer", err)
	}

	return buffer, nil
}
