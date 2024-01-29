package job

import (
	server "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/server/userdetailservice"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
