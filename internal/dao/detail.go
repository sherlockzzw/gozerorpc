package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"github.com/go-redis/redis/v8"
	"rpc-template/internal/config"
	"rpc-template/internal/model/entity"
)

type DetailDao struct {
	DB    mysqlx.DaoHandler[entity.Detail]
	Redis *redis.Client
}

func NewDetailDao() *DetailDao {
	return &DetailDao{
		DB:    mysqlx.NewDao[entity.Detail](config.DBClient),
		Redis: config.RedisClient,
	}
}
