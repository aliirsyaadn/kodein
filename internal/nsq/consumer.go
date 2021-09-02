package nsq

import (
	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/nsqio/go-nsq"
)

type consumerImpl struct {
	client  *nsq.Consumer
	host    string
	handler entity.ConsumerHandlerFunc
}

func NewConsumer(cfg config.NSQConsumerConfig, topic, channel string, handler entity.ConsumerHandlerFunc) *consumerImpl {
	conf := nsq.NewConfig()
	conf.MaxAttempts = cfg.MaxAttempts
	conf.MaxInFlight = cfg.MaxInFlight
	client, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		log.FatalDetail(intNSQTag, "error create consumer: %v", err)
	}
	return &consumerImpl{
		client:  client,
		host:    ParseDSN(cfg.ServerConfig),
		handler: handler,
	}
}

func (c *consumerImpl) Consume() {
	c.client.AddHandler(c)
	err := c.client.ConnectToNSQLookupd(c.host)
	if err != nil {
		log.FatalDetail(intNSQTag, "error consume handler: %v", err)
	}
}

func (c *consumerImpl) HandleMessage(msg *nsq.Message) error {
	return c.handler(msg.Body)
}

func (c *consumerImpl) Stop(){
	c.client.Stop()
}
