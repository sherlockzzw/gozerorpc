package config

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

func WatchEtcd(config Config) {
	watchDB(config)
	watchRedis(config)
	watchRabbitMq(config)
	watchMongo(config)
}

func watchMongo(config Config) {
	watchMongoChan := EtcdClient.Watch(context.Background(), config.MongoX)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch DB func panic: %s", p)
			}
		}()

		for {
			consume := <-watchMongoChan

			mongoX := &MongoX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, mongoX); err != nil {
				panic(err)
			}
			initMongo(mongoX)
		}
	}()
}

func watchDB(config Config) {
	watchDBChan := EtcdClient.Watch(context.Background(), config.DBX)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch DB func panic: %s", p)
			}
		}()

		for {
			consume := <-watchDBChan

			dbX := &DBX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, dbX); err != nil {
				panic(err)
			}
			initDB(config.Mode, dbX)
		}
	}()
}

func watchRedis(config Config) {
	watchDBChan := EtcdClient.Watch(context.Background(), config.RedisX)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch DB func panic: %s", p)
			}
		}()

		for {
			consume := <-watchDBChan

			redisX := &RedisX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, redisX); err != nil {
				panic(err)
			}
			initRedis(redisX)
		}
	}()
}

func watchRabbitMq(config Config) {
	watchRabbitMqChan := EtcdClient.Watch(context.Background(), config.RedisX)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch DB func panic: %s", p)
			}
		}()

		for {
			consume := <-watchRabbitMqChan

			rabbitMqX := &RabbitMqX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, rabbitMqX); err != nil {
				panic(err)
			}

			initProducer(rabbitMqX)
			initConsumer(rabbitMqX)
		}
	}()
}
