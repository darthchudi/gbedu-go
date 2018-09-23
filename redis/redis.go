package redis

import (
	"encoding/json"
	"fmt"
	Redis "github.com/go-redis/redis"
)

type channelMessage struct {
	message string `json:"message"`
}

func channelListener(channel <-chan *Redis.Message, closeChannel chan<- bool) {
	for {
		message, ok := <-channel
		res := channelMessage{}

		if !ok {
			break
		}

		if err := json.Unmarshal([]byte(message.Payload), &res); err != nil {
			fmt.Println("Could not parse response")
		}
		fmt.Println(res.message)
	}
	closeChannel <- true
}

//CreateRedisInstance creates a new Redis instance
func CreateRedisInstance() *Redis.Client {
	redis := Redis.NewClient(&Redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return redis
}

//InitPubSub initializes a pub sub listener
func InitPubSub(redis *Redis.Client) {
	pubsub := redis.Subscribe("go-node")
	defer pubsub.Close()

	_, err := pubsub.Receive()
	if err != nil {
		fmt.Println("Could not subscribe to channel")
	}

	channel := pubsub.Channel()
	closeChannel := make(chan bool)

	go channelListener(channel, closeChannel)
}
