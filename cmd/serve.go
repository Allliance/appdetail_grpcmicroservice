package cmd

import (
	"microservice-grpc/internal/grpcserver"
	"microservice-grpc/pkg/appdetail"
	"net"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts cache server",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

var (
	endpoint string = "appdetail.test.roo.cloud"
	port     int    = 8080
)

func serve(cmd *cobra.Command, _ []string) {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	cacheClient, err := grpcserver.NewAppDetailClient(endpoint, port)
	if err != nil {
		panic(err)
	}
	cacheServer, err := grpcserver.NewCacheServer(cacheClient)
	if err != nil {
		panic(err)
	}
	var grpcOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(grpcOptions...)
	appdetail.RegisterAppDetailServer(grpcServer, cacheServer)

	grpcServer.Serve(lis)
}
