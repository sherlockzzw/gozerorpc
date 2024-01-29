package config

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

func WatchEtcd(config Config) {
	wg := new(sync.WaitGroup)
	watchDB(config, wg)
	watchRedis(config, wg)
	watchRabbitMq(config, wg)
	watchMongo(config, wg)
}

func watchMongo(config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	watchMongoChan := etcdClient.Watch(context.Background(), config.MongoX)

	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch mongo func panic: %s", p)
			}
		}()

		for {
			consume := <-watchMongoChan

			mongoX := &MongoX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, mongoX); err != nil {
				panic(err)
			}
			initMongo(mongoX)
			return
		}
	}()
}

func watchDB(config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	watchDBChan := etcdClient.Watch(context.Background(), config.DBX)

	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch mysql func panic: %s", p)
			}
		}()

		for {
			consume := <-watchDBChan

			dbX := &DBX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, dbX); err != nil {
				panic(err)
			}
			initDB(config.Mode, dbX)
			return
		}
	}()
}

func watchRedis(config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	watchDBChan := etcdClient.Watch(context.Background(), config.RedisX)

	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch redis func panic: %s", p)
			}
		}()

		for {
			consume := <-watchDBChan

			redisX := &RedisX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, redisX); err != nil {
				panic(err)
			}
			initRedis(redisX)
			return
		}
	}()
}

func watchRabbitMq(config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	watchRabbitMqChan := etcdClient.Watch(context.Background(), config.RabbitMqX)

	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch rabbitMQ func panic: %s", p)
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
			return
		}
	}()
}
