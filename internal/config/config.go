package config

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rabbitmqx"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"rpc-template/internal/model/entity"
	"sync"
	"time"
)

type (
	Config struct {
		zrpc.RpcServerConf
		DBX       string
		RedisX    string
		RabbitMqX string
		MongoX    string
	}

	DBX struct {
		Uri string `json:"uri"`
	}

	RedisX struct {
		Addr     string `json:"addr"`
		UserName string `json:"username"`
		PassWord string `json:"password"`
		Db       int    `json:"db"`
	}

	RabbitMqX struct {
		Host        string `json:"host"`
		PassWord    string `json:"password"`
		Port        int    `json:"port"`
		User        string `json:"user"`
		VirtualHost string `json:"virtualHost"`
	}

	MongoX struct {
		Uri            string        `json:"uri"`
		ConnectTimeout time.Duration `json:"connectTimeout"`
		SocketTimeout  time.Duration `json:"socketTimeout"`
		MaxPoolSize    uint64        `json:"maxPoolSize"`
		MinPoolSize    uint64        `json:"minPoolSize"`
	}
)

var (
	configOnce = sync.Once{}

	DBClient     *gorm.DB
	MongoClient  *mongo.Client
	RedisClient  *redis.Client
	EtcdClient   *clientv3.Client
	ProducerPool *rabbitmqx.RabbitPool
	ConsumerPool *rabbitmqx.RabbitPool
)

func Init() (config *Config) {
	config = new(Config)

	configOnce.Do(func() {
		var (
			filePath string
			err      error
		)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				panic("read config yaml file failed")
			case <-time.After(time.Nanosecond):
				env := os.Getenv("environ")
				switch env {
				case service.ProMode:
					filePath, err = filepath.Abs("etc/rpcTemplate.yaml")
					if err != nil {
						panic(err)
					}
				case service.TestMode:
					filePath, err = filepath.Abs("etc/test.rpcTemplate.yaml")
					if err != nil {
						panic(err)
					}
				default:
					filePath, err = filepath.Abs("etc/dev.rpcTemplate.yaml")
					if err != nil {
						panic(err)
					}
				}

				conf.MustLoad(filePath, config)

				logx.WithContext(ctx).Infof("successfully set config, path: %s, env: %s", filePath, env)
				registerEtcd(ctx, config)

				return
			}
		}
	})

	return
}

func registerEtcd(ctx context.Context, config *Config) {
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   config.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	logx.WithContext(ctx).Info("successfully connect etcd client")
}

func RegisterDBCli(config Config) *gorm.DB {
	return initDB(config.Mode, getDB(config))
}

func RegisterRedisCli(config Config) *redis.Client {
	return initRedis(getRedis(config))
}

func RegisterRabbitMq(config Config, classes ...uint) *rabbitmqx.RabbitPool {
	for _, class := range classes {
		switch class {
		case rabbitmqx.RABBITMQ_TYPE_CONSUME:
			return initConsumer(getRabbitMq(config))
		case rabbitmqx.RABBITMQ_TYPE_PUBLISH:
			return initProducer(getRabbitMq(config))
		}
	}

	return nil
}

func RegisterMongoCli(config Config) *mongo.Client {
	return initMongo(getMongo(config))
}

func getMongo(config Config) (mongoX *MongoX) {
	dbResp, err := EtcdClient.Get(context.Background(), config.MongoX)
	if err != nil {
		panic(err)
	}

	mongoX = &MongoX{}
	if err = json.Unmarshal(dbResp.Kvs[0].Value, mongoX); err != nil {
		panic(err)
	}

	return
}

func getDB(config Config) (dbx *DBX) {
	dbResp, err := EtcdClient.Get(context.Background(), config.DBX)
	if err != nil {
		panic(err)
	}

	dbx = &DBX{}
	if err = json.Unmarshal(dbResp.Kvs[0].Value, dbx); err != nil {
		panic(err)
	}

	return
}

func getRedis(config Config) (redisX *RedisX) {
	redisResp, err := EtcdClient.Get(context.Background(), config.RedisX)
	if err != nil {
		panic(err)
	}

	redisX = &RedisX{}
	if err = json.Unmarshal(redisResp.Kvs[0].Value, redisX); err != nil {
		panic(err)
	}

	return
}

func getRabbitMq(config Config) (rabbitMqX *RabbitMqX) {
	rabbitResp, err := EtcdClient.Get(context.Background(), config.RabbitMqX)
	if err != nil {
		panic(err)
	}

	rabbitMqX = &RabbitMqX{}
	if err = json.Unmarshal(rabbitResp.Kvs[0].Value, rabbitMqX); err != nil {
		panic(err)
	}

	return
}

func initProducer(rabbitMqX *RabbitMqX) *rabbitmqx.RabbitPool {
	ProducerPool = rabbitmqx.NewProductPool()
	if err := ProducerPool.ConnectVirtualHost(rabbitMqX.Host, rabbitMqX.Port, rabbitMqX.User, rabbitMqX.PassWord, rabbitMqX.VirtualHost); err != nil {
		panic(err)
	}
	logx.Info("rabbitMQ producer pool init successfully")

	return ProducerPool
}

func initConsumer(rabbitMqX *RabbitMqX) *rabbitmqx.RabbitPool {
	ConsumerPool = rabbitmqx.NewConsumePool()
	if err := ConsumerPool.ConnectVirtualHost(rabbitMqX.Host, rabbitMqX.Port, rabbitMqX.User, rabbitMqX.PassWord, rabbitMqX.VirtualHost); err != nil {
		panic(err)
	}
	logx.Info("rabbitMQ consumer pool init successfully")

	return ConsumerPool
}

func initRedis(redisX *RedisX) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Username: redisX.UserName,
		Addr:     redisX.Addr,
		Password: redisX.PassWord, // 密码
		DB:       redisX.Db,       // 默认数据库
	})
	// 设置连接超时时间

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	logx.Infof("successfully connect redis client")

	return RedisClient
}

func initDB(env string, db *DBX) *gorm.DB {
	var err error

	DBClient, err = gorm.Open(mysql.Open(db.Uri), mysqlx.DefaultOpts())
	if err != nil {
		panic(err)
	}
	logx.Info("successfully connect mysql client")

	if env != service.ProMode {
		_ = DBClient.Set("gorm:table_options", "COMMENT='用户详情表'").AutoMigrate(&entity.Detail{})
	}

	return DBClient
}

func initMongo(mongoX *MongoX) *mongo.Client {
	var err error
	// 通过传进来的uri连接相关的配置
	opts := options.Client().ApplyURI(mongoX.Uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	opts.SetMaxPoolSize(mongoX.MaxPoolSize)
	opts.SetMinPoolSize(mongoX.MinPoolSize)
	opts.SetSocketTimeout(mongoX.SocketTimeout * time.Second)
	opts.SetConnectTimeout(mongoX.ConnectTimeout * time.Second)
	MongoClient, err = mongo.Connect(context.Background(), opts)

	err = MongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	logx.Info("mongo client init successfully")

	return MongoClient
}
