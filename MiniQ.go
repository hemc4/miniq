package miniq

import (
	"gopkg.in/redis.v3"
	"time"
)


type MiniQ struct {
	redisClient    *redis.Client
	Name           string
	lastStatsWrite int64
}

type Consumer struct {
	Name  string
	Queue *MiniQ
}
type MsgBundle struct {
	Id string
	Payload string
	Queue *MiniQ
	Consumer *Consumer
	Created time.Time
	Ack bool

}


func CreateQueue(redisHost, redisPort, redisPassword string, redisDB int64, name string) *MiniQ {
	// each queue will be a list in redis and all queue names will be stored in set
	// create the redis client
	// sadd to ad queue to a set
	return nil
}


func (queue *MiniQ) Put(payload string) error {
	//lpush to redis queue
	//payload will have some meta info as MsgBundle
	return nil
}




func (queue *MiniQ) AddConsumer(name string) (c *Consumer, err error) {
	// each consumer will a "inprocess queue"
	// sadd to consumer set, a set to keep track of all the consumers
	return c, nil
}

func (queue *MiniQ) Delete() error {

	//delete the queue and related comsumers

	return nil
}

func (consumer *Consumer) Get() error {
	//brpoplpush , atomic operation pop from the list and push to processing list, in case the consumer fails can
	// read again from the processing list
	return nil
}


func (consumer *Consumer) Ack(msg *MsgBundle) error {
	//notify that message has processed

	// this removes message from the processing queue
	// in case on ack the message will go back to main queue after some time
	return nil
}


