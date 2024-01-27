package svc

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rabbitmqx"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"rpc-template/internal/config"
)

type (
	ServiceContext struct {
		Config       config.Config
		DBCli        *gorm.DB
		RedisCli     *redis.Client
		MongoCli     *mongo.Client
		ConsumerPool *rabbitmqx.RabbitPool
		ProducerPool *rabbitmqx.RabbitPool
	}
)

func NewServiceContext(c config.Config) *ServiceContext {

	config.WatchEtcd(c)

	return &ServiceContext{
		Config:       c,
		DBCli:        config.RegisterDBCli(c),
		RedisCli:     config.RegisterRedisCli(c),
		MongoCli:     config.RegisterMongoCli(c),
		ConsumerPool: config.RegisterRabbitMq(c, rabbitmqx.RABBITMQ_TYPE_CONSUME),
		ProducerPool: config.RegisterRabbitMq(c, rabbitmqx.RABBITMQ_TYPE_PUBLISH),
	}
}
