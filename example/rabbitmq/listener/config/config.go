package config

import "github.com/lidaboo/go-queue/rabbitmq"

type Config struct {
	ListenerConf rabbitmq.RabbitListenerConf
}
