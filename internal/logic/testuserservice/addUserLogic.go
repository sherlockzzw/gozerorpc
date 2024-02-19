package testuserservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/tool"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	TestUserDao *dao.TestUserDao
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		TestUserDao: dao.NewTestUserDao(svcCtx),
	}
}

func (l *AddUserLogic) AddUser(in *rpcTemplate.AddUserRes) (*rpcTemplate.AddUserResp, error) {
	id, err := l.TestUserDao.Add(l.ctx, &entity.User{
		CommonDoc: &mysqlx.CommonDoc{
			CreatedAt: tool.GetTimeStamp(),
			UpdatedAt: tool.GetTimeStamp(),
		},
		Uuid:   in.Uuid,
		Mobile: in.Mobile,
		State:  1,
		Email:  in.Email,
	})
	if err != nil {
		l.Errorf("添加失败,err:%+v", err)
		return &rpcTemplate.AddUserResp{}, err
	}
	return &rpcTemplate.AddUserResp{Id: id}, nil
}
