package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Starts cache server",
	Run:   version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	Version string = "1.0.0"
	Title   string = "caching server with grpc"
	Author  string = "alliance"
)

func version(cmd *cobra.Command, _ []string) {
	fmt.Printf("%v\nBy: %v\nv%v", Title, Author, Version)
}
