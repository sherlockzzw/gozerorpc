package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mongox"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/model/entity"
	"github.com/go-redis/redis/v8"
)

type DetailDao struct {
	DB    mysqlx.DaoHandler[entity.Detail]
	Mongo mongox.DaoHandle[entity.DetailMongo]
	Redis *redis.Client
}

func NewDetailDao() *DetailDao {
	return &DetailDao{
		DB:    mysqlx.NewDao[entity.Detail](config.DBClient),
		Redis: config.RedisClient,
		Mongo: mongox.NewDao[entity.DetailMongo]().SetDatabase("test").SetCollection("detail").SetClient(config.MongoClient),
	}
}
