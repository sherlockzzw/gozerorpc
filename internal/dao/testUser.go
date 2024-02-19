package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/tool"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm/clause"
)

type TestUserDao struct {
	DB    mysqlx.DaoHandler[entity.User]
	Redis *redis.Client
}

func NewTestUserDao(svcCtx *svc.ServiceContext) *TestUserDao {
	return &TestUserDao{
		DB:    mysqlx.NewDao[entity.User](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

func (t *TestUserDao) Find(ctx context.Context, where *entity.User) (*entity.User, error) {

	return t.DB.FindOne(ctx, where)
}
func (t *TestUserDao) Add(ctx context.Context, add *entity.User) (id *int64, err error) {
	k, err := t.DB.Insert(ctx, add)
	return k.Id, err
}

func (b *TestUserDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.User) ([]*entity.User, int64, error) {
	order := []clause.OrderByColumn{
		{
			Column: clause.Column{
				Name: "id",
			},
			Desc:    true,
			Reorder: false,
		},
	}
	count, err := b.DB.Count(ctx, where)
	if err != nil {
		return nil, 0, err
	}
	items, err := b.DB.PaginateWithModel(ctx, where, order, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (b *TestUserDao) Update(ctx context.Context, where, date *entity.User) (row int64, err error) {
	tx := b.DB.GetClient().Where(where).Updates(date)
	return tx.RowsAffected, tx.Error
}
func (b *TestUserDao) Delete(ctx context.Context, id int64) (row int64, err error) {
	tx := b.DB.GetClient().Where("id = ?", id).Updates(&entity.User{CommonDoc: &mysqlx.CommonDoc{DeletedAt: tool.GetTimeStamp()}})
	return tx.RowsAffected, tx.Error
}
