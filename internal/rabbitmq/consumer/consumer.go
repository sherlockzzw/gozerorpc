package consumer

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rabbitmqx"
	server "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/server/userdetailservice"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type Consumer struct {
	userService *server.UserDetailServiceServer
}

func registerConsumerService(svcCtx *svc.ServiceContext) *Consumer {
	return &Consumer{
		userService: server.NewUserDetailServiceServer(svcCtx),
	}
}

func RegisterConsumer(svcCtx *svc.ServiceContext, wg *sync.WaitGroup) {
	registerConsumer(svcCtx, wg)
}

func registerConsumer(svcCtx *svc.ServiceContext, wg *sync.WaitGroup) {
	svcCtx.ConsumerPool.RegisterConsumeReceive(consumeDetailListCount(svcCtx, wg))

	if err := svcCtx.ConsumerPool.RunConsume(); err != nil {
		logx.Errorf("error:%s", err)
	}
}

func consumeDetailListCount(svcCtx *svc.ServiceContext, wg *sync.WaitGroup) *rabbitmqx.ConsumeReceive {
	wg.Add(1)

	defer wg.Done()

	_ = registerConsumerService(svcCtx)

	receive := &rabbitmqx.ConsumeReceive{
		ExchangeName: "testChange",
		ExchangeType: rabbitmqx.EXCHANGE_TYPE_FANOUT,
		Route:        "",
		QueueName:    "testQueue",
		IsTry:        true,
		MaxReTry:     5,
		IsAutoAck:    false,
		EventSuccess: func(data []byte, header map[string]interface{}, retryClient rabbitmqx.RetryClientInterface) bool {
			_ = retryClient.Ack()
			logx.Infof("rabbitmq receive msg: %s", data)
			return true
		},
		EventFail: func(code int, err error, bytes []byte) {
			logx.Errorf("rabbitmq receive code: %d, err: %s, body: %s", code, err, bytes)
		},
	}

	return receive

}
