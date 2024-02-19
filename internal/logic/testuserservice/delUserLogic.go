package testuserservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	TestUserDao *dao.TestUserDao
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		TestUserDao: dao.NewTestUserDao(svcCtx),
	}
}

func (l *DelUserLogic) DelUser(in *rpcTemplate.DelUserReq) (*rpcTemplate.DelUserResp, error) {

	if in.Id == 0 {
		return &rpcTemplate.DelUserResp{Status: true}, nil
	}
	_, err := l.TestUserDao.Delete(l.ctx, in.Id)

	if err != nil {
		l.Errorf("删除音乐失败,err:%+v", err)
		return &rpcTemplate.DelUserResp{Status: false}, err
	}
	return &rpcTemplate.DelUserResp{Status: true}, nil
}
