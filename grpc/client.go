package grpc

import (
	pb "Libs/grpc/proto/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func GRPC() {
	fmt.Println("Start gRPC client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSastServiceClient(conn)

	file, err := os.Open("C:/Projects Go/dockerwrapper.zip")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := c.Upload(context.Background())
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		err = stream.Send(&pb.UploadRequest{
			Chunk: &pb.FileChunk{
				Content: buffer[:n],
			},
		})
		if err != nil {
			log.Fatalf("failed to send chunk: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close and receive: %v", err)
	}

	log.Printf("Upload response: %v", resp.Message)
	log.Printf("Analysis report: %s", string(resp.Report))
	log.Printf("Analysis errors: %s", string(resp.Errors))
}
