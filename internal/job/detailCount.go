package job

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	server "rpc-template/internal/server/userdetailservice"
	"rpc-template/internal/svc"
	"rpc-template/pb/rpcTemplate"
)

type DetailCount struct {
	userService *server.UserDetailServiceServer
}

func registerDetailCountJob(svcCtx *svc.ServiceContext) *DetailCount {
	return &DetailCount{
		userService: server.NewUserDetailServiceServer(svcCtx),
	}
}

func (d *DetailCount) run(ctx context.Context) error {
	logx.WithContext(ctx).Infof("detail count cron job is running")
	resp, err := d.userService.UserDetailCount(ctx, &rpcTemplate.UserDetailCountReq{})
	if err != nil {
		return err
	}

	logx.WithContext(ctx).Infof("detail count: %d", resp.GetCount())

	return nil
}
