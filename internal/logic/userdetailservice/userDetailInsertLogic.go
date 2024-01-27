package userdetailservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"context"
	"rpc-template/internal/dao"
	"rpc-template/internal/model/entity"

	"rpc-template/internal/svc"
	"rpc-template/pb/rpcTemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailInsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DetailDao *dao.DetailDao
}

func NewUserDetailInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailInsertLogic {
	return &UserDetailInsertLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		DetailDao: dao.NewDetailDao(),
	}
}

func (l *UserDetailInsertLogic) UserDetailInsert(in *rpcTemplate.UserDetailInsertReq) (*rpcTemplate.UserDetailInsertResp, error) {
	doc := &entity.Detail{
		Name:      in.GetName(),
		CommonDoc: mysqlx.NewCommonDoc(),
	}

	if insert, err := l.DetailDao.DB.Insert(l.ctx, doc); err != nil {
		return nil, err
	} else {
		return &rpcTemplate.UserDetailInsertResp{
			Id: *insert.Id,
		}, nil
	}
}
