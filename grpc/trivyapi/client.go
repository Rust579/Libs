package trivyclient

import (
	grpc2 "Libs/grpc"
	"Libs/grpc/proto/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func GRPCZip() {
	fmt.Println("Start gRPC client for zip handler")

	client := NewClient("localhost", "50053")

	filePath := "C:/Projects Go/gog.zip"
	buffer, err := grpc2.LoadLocalFile(filePath)
	if err != nil {
		return
	}

	t1 := time.Now()

	resp, err := client.SendStream(context.Background(), buffer)
	if err != nil {
		fmt.Println("err", err)
	}

	log.Println("time cost:", time.Since(t1))
	log.Printf("Upload response: %v", resp)
}

func GRPCImage() {
	fmt.Println("Start gRPC client for image handler")

	client := NewClient("localhost", "50053")

	req := rpc.ImageRequest{
		Data: []string{
			"golang:1.22",
			"golang:1.23",
		},
	}

	t1 := time.Now()

	resp, err := client.SendImage(context.Background(), &req)
	if err != nil {
		fmt.Println("err", err)
	}

	log.Println("time cost:", time.Since(t1))
	log.Printf("Upload response: %v", resp)

}

const chunkSize = 1024

type Client struct {
	host string
	port string
}

func NewClient(host, port string) Client {
	return Client{host: host, port: port}
}

func (c Client) SendStream(ctx context.Context, data []byte) (*rpc.Response, error) {
	conn, err := grpc.NewClient(c.host+":"+c.port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	grpcCl := rpc.NewTrivyServiceClient(conn)

	stream, err := grpcCl.UploadStream(ctx)
	if err != nil {
		return nil, err
	}

	// send full chunks
	for i := 0; i < len(data)/chunkSize; i++ {
		if err = stream.Send(&rpc.ZipRequest{
			Chunk: data[chunkSize*i : chunkSize*(i+1)],
		}); err != nil {
			return nil, err
		}
	}
	// send tale chunk
	if err = stream.Send(&rpc.ZipRequest{
		Chunk: data[chunkSize*(len(data)/chunkSize):],
	}); err != nil {
		return nil, err
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) SendImage(ctx context.Context, req *rpc.ImageRequest) (*rpc.Response, error) {
	conn, err := grpc.NewClient(c.host+":"+c.port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	grpcCl := rpc.NewTrivyServiceClient(conn)

	resp, err := grpcCl.UploadImage(ctx, req)

	return resp, nil
}
