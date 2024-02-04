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

type UserDetailCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DetailDao *dao.DetailDao
}

func NewUserDetailCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailCountLogic {
	return &UserDetailCountLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		DetailDao: dao.NewDetailDao(svcCtx),
	}
}

func (l *UserDetailCountLogic) UserDetailCount(in *rpcTemplate.UserDetailCountReq) (*rpcTemplate.UserDetailCountResp, error) {
	doc := &entity.Detail{
		CommonDoc: &mysqlx.CommonDoc{
			DeletedAt: mysqlx.DeletedAtTimeUnix(0),
		},
	}

	count, err := l.DetailDao.DB.Count(l.ctx, doc)
	if err != nil {
		return nil, err
	}

	return &rpcTemplate.UserDetailCountResp{
		Count: count,
	}, nil
}
