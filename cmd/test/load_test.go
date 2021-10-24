package main

import (
	"context"
	"log"
	"microservice-grpc/internal/grpcserver"
	"microservice-grpc/pkg/appdetail"
	"net"
	"sync"
	"testing"

	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc"
)

type mockClient struct {
}

func (client *mockClient) GetAppDetail(ctx context.Context, in *appdetail.GetAppDetailRequest, opts ...grpc.CallOption) (*appdetail.GetAppDetailReply, error) {
	apiCalls++
	return &appdetail.GetAppDetailReply{
		Detail: &appdetail.App{
			Name: "Test App Name",
		},
		StatusCode: 200,
	}, nil
}

const bufSize = 1024 * 1024

var (
	lis      *bufconn.Listener
	apiCalls int
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cacheServer, err := grpcserver.NewCacheServer(&mockClient{})
	if err != nil {
		panic(err)
	}
	appdetail.RegisterAppDetailServer(s, cacheServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestLoad(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(100)
	client := appdetail.NewAppDetailClient(conn)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := client.GetAppDetail(ctx, &appdetail.GetAppDetailRequest{PackageName: "Some Package"})
			if err != nil {
				t.Errorf("Error recieving package from server")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if apiCalls > 1 {
		t.Errorf("Server called api multiple times for some package : %v", apiCalls)
	}
}
