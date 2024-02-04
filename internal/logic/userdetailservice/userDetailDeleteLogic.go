package userdetailservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DetailDao *dao.DetailDao
}

func NewUserDetailDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailDeleteLogic {
	return &UserDetailDeleteLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		DetailDao: dao.NewDetailDao(svcCtx),
	}
}

func (l *UserDetailDeleteLogic) UserDetailDelete(in *rpcTemplate.UserDetailDeleteReq) (*rpcTemplate.UserDetailDeleteResp, error) {
	id := in.GetId()
	filter := &entity.Detail{
		CommonDoc: &mysqlx.CommonDoc{
			Id: &id,
		},
	}

	update := &entity.Detail{
		CommonDoc: &mysqlx.CommonDoc{
			DeletedAt: mysqlx.NewTimeUnix(),
		},
	}

	if err := l.DetailDao.DB.UpdateWithModel(l.ctx, filter, update); err != nil {
		return nil, err
	} else {
		return &rpcTemplate.UserDetailDeleteResp{
			Id: id,
		}, nil
	}
}
