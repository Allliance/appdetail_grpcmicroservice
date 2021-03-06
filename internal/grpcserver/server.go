package grpcserver

import (
	"context"
	"fmt"
	client "microservice-grpc/internal/client"
	appdetail "microservice-grpc/pkg/appdetail"
	appcache "microservice-grpc/pkg/cache"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type cacheServer struct {
	appdetail.UnimplementedAppDetailServer
	cache          appcache.AppDetailCache
	clientInstance appdetail.AppDetailClient
}

type pendingRequest struct {
	packageName string
	mu          *sync.Mutex
}

var pendingRequests map[string]*pendingRequest = make(map[string]*pendingRequest)

func (server cacheServer) GetAppDetail(ctx context.Context, request *appdetail.GetAppDetailRequest) (*appdetail.GetAppDetailReply, error) {
	cacheName := getCacheName(request.PackageName)
	result, err := checkForDataInCache(server.cache, cacheName)
	if err == nil {
		return result, nil
	}
	pending := pendingRequests[request.GetPackageName()]
	if pending == nil {
		pending = &pendingRequest{packageName: cacheName, mu: &sync.Mutex{}}
		pendingRequests[request.GetPackageName()] = pending
		defer delete(pendingRequests, cacheName)
	}
	pending.mu.Lock()
	defer pending.mu.Unlock()
	if result, err := checkForDataInCache(server.cache, cacheName); err == nil {
		return result, nil
	}
	appDetail, err := client.GetAppDetail(ctx, request.GetPackageName(), server.clientInstance)
	if err != nil {
		return nil, err
	}
	appDetailReply, err := proto.Marshal(appDetail)
	if err != nil {
		return nil, err
	}
	server.cache.CacheApp(getCacheName(request.GetPackageName()), appDetailReply)
	return appDetail, nil
}

func checkForDataInCache(cache appcache.AppDetailCache, key string) (*appdetail.GetAppDetailReply, error) {
	data, err := cache.Retrieve(key)
	if err == nil {
		result := &appdetail.GetAppDetailReply{}
		err := proto.Unmarshal(data, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, err
}

func getCacheName(packageName string) string {
	return fmt.Sprintf("#package#%v", packageName)
}

func NewAppDetailClient(endpoint string, port int) (appdetail.AppDetailClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", endpoint, port), opts...)
	if err != nil {
		return nil, err
	}
	client := appdetail.NewAppDetailClient(conn)
	return client, nil
}

func NewCacheServer(client appdetail.AppDetailClient) (*cacheServer, error) {
	serverCache, err := appcache.NewCache()
	if err != nil {
		return nil, err
	}
	server := &cacheServer{cache: *serverCache, clientInstance: client}
	return server, nil
}
