package testuserservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	TestUserDao *dao.TestUserDao
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		TestUserDao: dao.NewTestUserDao(svcCtx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(res *rpcTemplate.GetUserRes) (*rpcTemplate.GetUserResp, error) {
	info, err := l.TestUserDao.Find(l.ctx, &entity.User{CommonDoc: &mysqlx.CommonDoc{
		Id: &res.Id, DeletedAt: mysqlx.DeletedAtTimeUnix(0),
	}})

	if err != nil {
		l.Errorf("获取失败,err:%+v", err)
		return &rpcTemplate.GetUserResp{}, err
	}

	return &rpcTemplate.GetUserResp{
		Id:     *info.Id,
		Mobile: info.Mobile,
		Uuid:   info.Uuid,
		Email:  info.Email,
		State:  info.State,
	}, nil
}
