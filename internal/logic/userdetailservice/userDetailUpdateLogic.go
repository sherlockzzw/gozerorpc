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

type UserDetailUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DetailDao *dao.DetailDao
}

func NewUserDetailUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailUpdateLogic {
	return &UserDetailUpdateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		DetailDao: dao.NewDetailDao(),
	}
}

func (l *UserDetailUpdateLogic) UserDetailUpdate(in *rpcTemplate.UserDetailUpdateRequest) (*rpcTemplate.UserDetailUpdateResponse, error) {
	id := in.GetQuery().GetId()
	filter := &entity.Detail{
		CommonDoc: &mysqlx.CommonDoc{
			Id: &id,
		},
	}

	update := &entity.Detail{
		Name: in.GetDoc().GetName(),
		CommonDoc: &mysqlx.CommonDoc{
			UpdatedAt: mysqlx.NewTimeUnix(),
		},
	}

	if err := l.DetailDao.DB.UpdateWithModel(l.ctx, filter, update); err != nil {
		return nil, err
	} else {
		return &rpcTemplate.UserDetailUpdateResponse{
			Id: id,
		}, nil
	}
}
