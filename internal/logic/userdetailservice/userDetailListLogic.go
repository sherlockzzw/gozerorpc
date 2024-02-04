package userdetailservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rabbitmqx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"context"
	"encoding/json"
	"gorm.io/gorm/clause"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/pb/rpcTemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DetailDao *dao.DetailDao
}

func NewUserDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailListLogic {
	return &UserDetailListLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		DetailDao: dao.NewDetailDao(svcCtx),
	}
}

func (l *UserDetailListLogic) UserDetailList(in *rpcTemplate.UserDetailListRequest) (*rpcTemplate.UserDetailListResponse, error) {
	order := []clause.OrderByColumn{
		{
			Column: clause.Column{
				Name: "id",
			},
			Desc:    true,
			Reorder: false,
		},
	}

	filter := &entity.Detail{
		CommonDoc: &mysqlx.CommonDoc{
			DeletedAt: mysqlx.DeletedAtTimeUnix(0),
		},
	}

	items, err := l.DetailDao.DB.PaginateWithModel(l.ctx, filter, order, in.GetPage(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	resp := &rpcTemplate.UserDetailListResponse{
		Items: make([]*rpcTemplate.UserDetailItem, 0),
	}

	for _, item := range items {
		doc := &rpcTemplate.UserDetailItem{
			Id:        *item.Id,
			Name:      item.Name,
			CreatedAt: *item.CreatedAt,
			UpdatedAt: *item.UpdatedAt,
		}

		resp.Items = append(resp.Items, doc)
	}

	data := map[string]interface{}{
		"count": len(resp.Items),
	}

	if err = l.DetailDao.Redis.HMSet(l.ctx, "user_detail", data).Err(); err != nil {
		l.Logger.Errorf("redis err: %+v", err)
		return nil, err
	}

	if err_ := l.PushMQ(data); err_ != nil {
		l.Logger.Errorf("rabbitMQ err: %+v", err_)
		return nil, err
	}

	return resp, nil
}

func (l *UserDetailListLogic) PushMQ(data interface{}) *rabbitmqx.RabbitMqError {
	marshal, err := json.Marshal(data)
	if err != nil {
		return rabbitmqx.NewRabbitMqError(500, err.Error(), "")
	}
	dataFormat := rabbitmqx.GetDataFormat("testChange", rabbitmqx.EXCHANGE_TYPE_FANOUT, "testQueue", "", marshal, uint8(0))
	if err_ := l.svcCtx.ProducerPool.Push(dataFormat); err_ != nil {
		return err_
	} else {
		return nil
	}
}
