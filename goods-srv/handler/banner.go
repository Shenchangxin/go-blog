package handler

import (
	"context"
	"go-blog/goods-srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*GoodsServer) BannerList(context.Context, *emptypb.Empty) (*proto.BannerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BannerList not implemented")
}
func (*GoodsServer) CreateBanner(context.Context, *proto.BannerRequest) (*proto.BannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanner not implemented")
}
func (*GoodsServer) DeleteBanner(context.Context, *proto.BannerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (*GoodsServer) UpdateBanner(context.Context, *proto.BannerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
