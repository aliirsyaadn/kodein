package nsq

import (
	"encoding/json"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/nsqio/go-nsq"
)

const intNSQTag = "InternalNSQTag"

type Producer interface {
	Publish(topic string, data interface{}) error
}

type producerImpl struct {
	client *nsq.Producer
}

func NewProducer(cfg config.NSQProducerConfig) Producer{
	client, err := nsq.NewProducer(ParseDSN(cfg.ServerConfig), nsq.NewConfig())
	if err != nil {
		log.FatalDetail(intNSQTag, "error create producer: %v", err)
	}
	return &producerImpl{
		client: client,
	}
}

func (p *producerImpl) Publish(topic string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		log.ErrorDetail(intNSQTag, "error marshall publish message: %v", err)
		return err
	}
	return p.client.Publish(topic, body)
}
