package client

import (
	"context"
	"microservice-grpc/pkg/appdetail"
)

func GetAppDetail(ctx context.Context, packageName string, appdetailClient appdetail.AppDetailClient) (*appdetail.GetAppDetailReply, error) {
	request := appdetail.GetAppDetailRequest{PackageName: packageName}
	return appdetailClient.GetAppDetail(ctx, &request)
}
