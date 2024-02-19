package testuserservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	TestUserDao *dao.TestUserDao
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		TestUserDao: dao.NewTestUserDao(svcCtx),
	}
}

func (l *GetUserListLogic) GetUserList(in *rpcTemplate.GetUserListReq) (*rpcTemplate.GetUserListResp, error) {

	where := &entity.User{
		CommonDoc: &mysqlx.CommonDoc{DeletedAt: mysqlx.DeletedAtTimeUnix(0)},
		State:     1,
	}

	list, count, err := l.TestUserDao.GetPageList(l.ctx, in.GetPage(), in.GetPageSize(), where)
	res := &rpcTemplate.GetUserListResp{
		Total: count,
		Page:  in.GetPage(),
	}
	if err != nil {
		l.Errorf("获取列表失败,err:%+v", in)
		return &rpcTemplate.GetUserListResp{}, errors.New(fmt.Sprintf("获取列表失败,err:%+v", in))
	}
	for _, info := range list {
		res.List = append(res.List, &rpcTemplate.UserInfo{
			Id:        *info.Id,
			Uuid:      info.Uuid,
			Mobile:    info.Mobile,
			Email:     info.Email,
			State:     info.State,
			CreatedAt: *info.CreatedAt,
			UpdatedAt: *info.UpdatedAt,
		})
	}
	return res, nil
}
