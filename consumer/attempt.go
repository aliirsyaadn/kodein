package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/nsq"
	"github.com/aliirsyaadn/kodein/model"
)

const attempt = "attempt"
const attemptConsumerTag = "AttempConsumerTag"

func main(){
	cfgConsumer := config.LoadConfigNSQConsumer()

	consumer := nsq.NewConsumer(cfgConsumer, attempt, "test", handleAttempt)

	consumer.Consume()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	// Gracefully stop the consumer.
	consumer.Stop()
}

func handleAttempt(body []byte) error {
	var data model.Attempt
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.ErrorDetail(attemptConsumerTag, "error unmarshal message", err)
		return err
	}
	log.InfoDetail(attemptConsumerTag, "consume data %+v", data)
	return nil
}