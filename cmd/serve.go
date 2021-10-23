package cmd

import (
	"microservice-grpc/internal/grpcserver"
	"net"

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

func serve(cmd *cobra.Command, _ []string) {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	server, err := grpcserver.NewCacheServer()
	if err != nil {
		panic(err)
	}
	server.Serve(lis)
}
