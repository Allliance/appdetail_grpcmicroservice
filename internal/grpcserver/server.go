package grpcserver

import (
	"context"
	"fmt"
	client "microservice-grpc/internal/client"
	appdetail "microservice-grpc/pkg/appdetail"
	appcache "microservice-grpc/pkg/cache"

	proto "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type cacheServer struct {
	appdetail.UnimplementedAppDetailServer
	cache          appcache.AppDetailCache
	clientInstance appdetail.AppDetailClient
}

func (server cacheServer) GetAppDetail(ctx context.Context, request *appdetail.GetAppDetailRequest) (*appdetail.GetAppDetailReply, error) {
	if data, err := server.cache.Retrieve(getCacheName(request.GetPackageName())); err != nil {
		result := &appdetail.GetAppDetailReply{}
		err := proto.Unmarshal(data, result)
		if err != nil {
			return result, err
		}
		return result, nil
	}
	appDetail, err := client.GetAppDetail(ctx, request.GetPackageName(), server.clientInstance)
	if err != nil {
		return nil, err
	}
	if data, err := proto.Marshal(appDetail); err != nil {
		return nil, err
	} else {
		server.cache.CacheApp(getCacheName(request.GetPackageName()), data)
	}
	return appDetail, nil
}

func getCacheName(packageName string) string {
	return fmt.Sprintf("#package#%v", packageName)
}

func newAppDetailClient(endpoint string, port int) (appdetail.AppDetailClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", endpoint, port), opts...)
	if err != nil {
		return nil, err
	}
	client := appdetail.NewAppDetailClient(conn)
	return client, nil
}

var endpoint string = "appdetail.test.roo.cloud"
var port int = 8080

func NewCacheServer() (*grpc.Server, error) {
	serverCache, err := appcache.NewCache()
	if err != nil {
		return nil, err
	}
	client, err := newAppDetailClient(endpoint, port)
	if err != nil {
		return nil, err
	}
	var grpcOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(grpcOptions...)
	server := cacheServer{cache: *serverCache, clientInstance: client}
	appdetail.RegisterAppDetailServer(grpcServer, server)
	return grpcServer, nil
}
