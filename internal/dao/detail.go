package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mongox"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"github.com/go-redis/redis/v8"
	"rpc-template/internal/config"
	"rpc-template/internal/model/entity"
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
