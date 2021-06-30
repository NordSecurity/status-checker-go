package checker

import (
	"github.com/NordSec/status-checker-go"
	"github.com/go-redis/redis/v7"
)

func NewRedisChecker(client redis.UniversalClient) status.Checker {
	return &redisChecker{client}
}

type redisChecker struct {
	client redis.UniversalClient
}

func (rc *redisChecker) Name() string {
	return "redis"
}

func (rc *redisChecker) Status() status.Status {
	result, err := rc.client.Ping().Result()
	if err != nil || result != "PONG" {
		return status.DOWN
	}

	return status.OK
}
