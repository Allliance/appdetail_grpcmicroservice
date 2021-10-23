package main

import (
	"microservice-grpc/internal/grpcserver"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		panic(err)
	}
	grpcServer, err := grpcserver.NewCacheServer()
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}
