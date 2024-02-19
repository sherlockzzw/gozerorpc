package testuserservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/tool"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	TestUserDao *dao.TestUserDao
}

func NewUpdUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdUserLogic {
	return &UpdUserLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		TestUserDao: dao.NewTestUserDao(svcCtx),
	}
}

func (l *UpdUserLogic) UpdUser(in *rpcTemplate.UpdUserReq) (*rpcTemplate.UpdUserResp, error) {

	data := &entity.User{CommonDoc: &mysqlx.CommonDoc{UpdatedAt: tool.GetTimeStamp()}}
	if len(in.Uuid) > 0 {
		data.Uuid = in.GetUuid()
	}
	if len(in.Mobile) > 0 {
		data.Mobile = in.GetMobile()
	}
	if len(in.Email) > 0 {
		data.Email = in.GetEmail()
	}
	if in.State != 0 {
		data.State = in.GetState()
	}
	if len(in.Password) > 0 {
		data.Password = in.GetPassword()
	}

	row, err := l.TestUserDao.Update(l.ctx, &entity.User{CommonDoc: &mysqlx.CommonDoc{Id: &in.Id}}, data)
	if err != nil {
		l.Errorf("更新失败,err:%+v", err)
		return &rpcTemplate.UpdUserResp{}, err
	}
	return &rpcTemplate.UpdUserResp{Row: row}, nil
}
