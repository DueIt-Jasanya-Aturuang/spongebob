package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RedisImpl struct {
	Client *redis.Client
}

func NewRedisConn() *RedisImpl {
	host := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)
	rDB := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: config.RedisPass,
		DB:       config.RedisDb,
	})

	ctx := context.TODO()
	ping, err := rDB.Ping(ctx).Result()
	if err != nil {
		log.Err(err).Msg("ping redis error")
		os.Exit(1)
	}

	log.Info().Msgf("connection to redis : %s", ping)
	return &RedisImpl{
		Client: rDB,
	}
}
