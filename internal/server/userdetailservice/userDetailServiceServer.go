// Code generated by goctl. DO NOT EDIT.
// Source: rpcTemplate.proto

package server

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/logic/userdetailservice"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"
)

type UserDetailServiceServer struct {
	svcCtx *svc.ServiceContext
	rpcTemplate.UnimplementedUserDetailServiceServer
}

func NewUserDetailServiceServer(svcCtx *svc.ServiceContext) *UserDetailServiceServer {
	return &UserDetailServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserDetailServiceServer) UserDetailList(ctx context.Context, in *rpcTemplate.UserDetailListRequest) (*rpcTemplate.UserDetailListResponse, error) {
	l := userdetailservicelogic.NewUserDetailListLogic(ctx, s.svcCtx)
	return l.UserDetailList(in)
}

func (s *UserDetailServiceServer) UserDetailCount(ctx context.Context, in *rpcTemplate.UserDetailCountReq) (*rpcTemplate.UserDetailCountResp, error) {
	l := userdetailservicelogic.NewUserDetailCountLogic(ctx, s.svcCtx)
	return l.UserDetailCount(in)
}

func (s *UserDetailServiceServer) UserDetailUpdate(ctx context.Context, in *rpcTemplate.UserDetailUpdateRequest) (*rpcTemplate.UserDetailUpdateResponse, error) {
	l := userdetailservicelogic.NewUserDetailUpdateLogic(ctx, s.svcCtx)
	return l.UserDetailUpdate(in)
}

func (s *UserDetailServiceServer) UserDetailInsert(ctx context.Context, in *rpcTemplate.UserDetailInsertReq) (*rpcTemplate.UserDetailInsertResp, error) {
	l := userdetailservicelogic.NewUserDetailInsertLogic(ctx, s.svcCtx)
	return l.UserDetailInsert(in)
}

func (s *UserDetailServiceServer) UserDetailDelete(ctx context.Context, in *rpcTemplate.UserDetailDeleteReq) (*rpcTemplate.UserDetailDeleteResp, error) {
	l := userdetailservicelogic.NewUserDetailDeleteLogic(ctx, s.svcCtx)
	return l.UserDetailDelete(in)
}
