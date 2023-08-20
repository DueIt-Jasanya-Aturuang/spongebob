package infrastructures

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/db"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/logs"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/redis"
)

type Init struct {
	db.SQL
	*redis.RedisImpl
}

func NewInit() *Init {
	config.EnvInit()
	logs.LogInit()
	db := db.NewSQLImpl()
	return &Init{
		db,
		redis.NewRedisConn(),
	}
}
